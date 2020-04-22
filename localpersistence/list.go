package localpersistence

import (
	"bytes"
	"encoding/binary"

	"github.com/syndtr/goleveldb/leveldb"
)

type List struct {
	db         *DB
	opts       *Options
	leftIndex  int64
	rightIndex int64
}

// 打开数据库
func OpenList(file string, opts *Options) (*List, error) {
	opts = fillDefaultOptions(opts)
	db, err := openDb(file, opts.DB)
	if err != nil {
		return nil, err
	}
	return &List{db: db, opts: opts}, nil
}

// 更新索引
func (l *List) updateIndex() (err error) {
	iter := l.db.NewIterator(nil, nil)
	defer iter.Release()
	if !iter.First() {
		return nil
	}
	first := iter.Key()
	l.leftIndex, err = l.decodeKey(first)
	if err != nil {
		return err
	}
	if !iter.Last() {
		l.rightIndex = l.leftIndex
		return nil
	}
	last := iter.Key()
	l.rightIndex, err = l.decodeKey(last)
	if err != nil {
		return err
	}
	return nil
}

// 编码key
func (l *List) encodeKey(index int64) ([]byte, error) {
	w := bytes.NewBuffer(make([]byte, 9))
	if index < 0 {
		w.WriteRune('-')
	} else {
		w.WriteRune('=')
	}
	if err := binary.Write(w, binary.BigEndian, index); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// 解码key
func (l *List) decodeKey(b []byte) (int64, error) {
	var value int64
	r := bytes.NewReader(b[1:])
	if err := binary.Read(r, binary.BigEndian, &value); err != nil {
		return 0, err
	}
	return value, nil
}

// 添加到队列首
func (l *List) AddToFirst(value interface{}) error {
	l.leftIndex--
	k, err := l.encodeKey(l.leftIndex)
	if err != nil {
		return err
	}
	b, err := l.opts.Encoder(value)
	if err != nil {
		return err
	}
	return l.db.Put(k, b, nil)
}

// 添加到队列尾
func (l *List) AddToLast(value interface{}) error {
	l.rightIndex++
	k, err := l.encodeKey(l.rightIndex)
	if err != nil {
		return err
	}
	b, err := l.opts.Encoder(value)
	if err != nil {
		return err
	}
	return l.db.Put(k, b, nil)
}

// 删除队列首
func (l *List) RemoveFirst(value interface{}) (ok bool, err error) {
	for l.leftIndex <= l.rightIndex {
		k, err := l.encodeKey(l.leftIndex)
		if err != nil {
			return false, err
		}
		b, err := l.db.Get(k, nil)
		if err == leveldb.ErrNotFound {
			l.leftIndex++
			continue
		}
		if err != nil {
			return false, err
		}
		l.leftIndex++
		if err := l.opts.Decoder(b, value); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 删除队列尾
func (l *List) RemoveLast(value interface{}) (ok bool, err error) {
	for l.rightIndex >= l.leftIndex {
		k, err := l.encodeKey(l.rightIndex)
		if err != nil {
			return false, err
		}
		b, err := l.db.Get(k, nil)
		if err == leveldb.ErrNotFound {
			l.rightIndex--
			continue
		}
		if err != nil {
			return false, err
		}
		l.rightIndex--
		if err := l.opts.Decoder(b, value); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 队列元素数量
func (l *List) Size() (size int, err error) {
	iter := l.db.NewIterator(nil, nil)
	defer iter.Release()
	if !iter.First() {
		return 0, nil
	}
	for {
		size++
		if !iter.Next() {
			break
		}
	}
	return size, nil
}

// 关闭数据库
func (l *List) Close() error {
	return l.db.Close()
}
