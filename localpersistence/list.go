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
func (list *List) updateIndex() (err error) {
	iter := list.db.NewIterator(nil, nil)
	defer iter.Release()
	if !iter.First() {
		return nil
	}
	first := iter.Key()
	list.leftIndex, err = list.decodeKey(first)
	if err != nil {
		return err
	}
	if !iter.Last() {
		list.rightIndex = list.leftIndex
		return nil
	}
	last := iter.Key()
	list.rightIndex, err = list.decodeKey(last)
	if err != nil {
		return err
	}
	return nil
}

// 编码key
func (list *List) encodeKey(index int64) ([]byte, error) {
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
func (list *List) decodeKey(b []byte) (int64, error) {
	var value int64
	r := bytes.NewReader(b[1:])
	if err := binary.Read(r, binary.BigEndian, &value); err != nil {
		return 0, err
	}
	return value, nil
}

// 添加到队列首
func (list *List) AddToFirst(value interface{}) error {
	list.leftIndex--
	k, err := list.encodeKey(list.leftIndex)
	if err != nil {
		return err
	}
	b, err := list.opts.Encoder(value)
	if err != nil {
		return err
	}
	return list.db.Put(k, b, nil)
}

// 添加到队列尾
func (list *List) AddToLast(value interface{}) error {
	list.rightIndex++
	k, err := list.encodeKey(list.rightIndex)
	if err != nil {
		return err
	}
	b, err := list.opts.Encoder(value)
	if err != nil {
		return err
	}
	return list.db.Put(k, b, nil)
}

// 删除队列首
func (list *List) RemoveFirst(value interface{}) (ok bool, err error) {
	for list.leftIndex <= list.rightIndex {
		k, err := list.encodeKey(list.leftIndex)
		if err != nil {
			return false, err
		}
		b, err := list.db.Get(k, nil)
		if err == leveldb.ErrNotFound {
			list.leftIndex++
			continue
		}
		if err != nil {
			return false, err
		}
		list.leftIndex++
		if err := list.opts.Decoder(b, value); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 删除队列尾
func (list *List) RemoveLast(value interface{}) (ok bool, err error) {
	for list.rightIndex >= list.leftIndex {
		k, err := list.encodeKey(list.rightIndex)
		if err != nil {
			return false, err
		}
		b, err := list.db.Get(k, nil)
		if err == leveldb.ErrNotFound {
			list.rightIndex--
			continue
		}
		if err != nil {
			return false, err
		}
		list.rightIndex--
		if err := list.opts.Decoder(b, value); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 队列元素数量
func (list *List) Size() (size int, err error) {
	iter := list.db.NewIterator(nil, nil)
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
func (list *List) Close() error {
	return list.db.Close()
}
