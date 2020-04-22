package localpersistence

import (
	"bytes"
	"encoding/binary"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

const sortedListCacheSeconds = 1

type SortedList struct {
	db                     *DB
	opts                   *Options
	cacheIterator          iterator.Iterator
	cacheIteratorTimestamp int64
}

// 打开数据库
func OpenSortedList(file string, opts *Options) (*SortedList, error) {
	opts = fillDefaultOptions(opts)
	db, err := openDb(file, opts.DB)
	if err != nil {
		return nil, err
	}
	return &SortedList{db: db, opts: opts}, nil
}

// 编码key
func (l *SortedList) encodeKey(score int64, value interface{}) ([]byte, error) {
	vb, err := l.opts.Encoder(value)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer([]byte{})
	if score < 0 {
		w.WriteRune('-')
	} else {
		w.WriteRune('=')
	}
	if err := binary.Write(w, binary.BigEndian, score); err != nil {
		return nil, err
	}
	w.Write(vb)
	return w.Bytes(), nil
}

// 解码key
func (l *SortedList) decodeKey(b []byte, value interface{}) (score int64, err error) {
	score, _ = binary.Varint(b[1:9])
	vb := b[9:]
	if err := jsoniter.Unmarshal(vb, value); err != nil {
		return 0, err
	}
	return score, nil
}

func (l *SortedList) reopenCacheIterator() (iterator.Iterator, error) {
	if l.cacheIterator != nil {
		l.cacheIterator.Release()
		l.cacheIterator = nil
	}
	l.cacheIterator = l.db.NewIterator(nil, nil)
	l.cacheIterator.First()
	ts := time.Now().Unix()
	l.cacheIteratorTimestamp = ts - ts%sortedListCacheSeconds + sortedListCacheSeconds
	return l.cacheIterator, nil
}

func (l *SortedList) getCacheIterator() (iterator.Iterator, error) {
	if l.cacheIterator == nil {
		return l.reopenCacheIterator()
	} else if l.cacheIteratorTimestamp < time.Now().Unix() {
		return l.reopenCacheIterator()
	} else if !l.cacheIterator.Next() {
		return l.reopenCacheIterator()
	}
	return l.cacheIterator, nil
}

func (l *SortedList) Add(score int64, value interface{}) error {
	b, err := l.encodeKey(score, value)
	if err != nil {
		return err
	}
	return l.db.Put(b, nil, nil)
}

// 取得第一个元素
func (l *SortedList) First(maxScore int64, value interface{}) (ok bool, err error) {
	iter, err := l.getCacheIterator()
	if err != nil {
		return false, err
	}
	if !iter.Valid() {
		return false, nil
	}
	key := iter.Key()
	score, err := l.decodeKey(key, value)
	if err != nil {
		l.cacheIterator.Prev()
		return false, err
	}
	if score > maxScore {
		l.cacheIterator.Prev()
		return false, nil
	}
	if err := l.db.Delete(key, nil); err != nil {
		l.cacheIterator.Prev()
		return false, nil
	}
	return true, nil
}

// 队列元素数量
func (l *SortedList) Size() (size int, err error) {
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
func (l *SortedList) Close() error {
	if l.cacheIterator != nil {
		l.cacheIterator.Release()
	}
	return l.db.Close()
}
