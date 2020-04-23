package localpersistence

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type SortedList struct {
	mu                     sync.Mutex
	db                     *DB
	opts                   *SortedListOptions
	cacheIterator          iterator.Iterator
	cacheIteratorTimestamp int64
	seq                    uint64
}

type SortedListOptions struct {
	*Options
	IteratorMaxMilliseconds int64 // 迭代器缓存时间
}

var sortedListSeqKey = []byte("~seq")

// 打开数据库
func OpenSortedList(file string, opts *SortedListOptions) (*SortedList, error) {
	opts = fillDefaultSortedListOptions(opts)
	db, err := openDb(file, opts.DB)
	if err != nil {
		return nil, err
	}
	return &SortedList{db: db, opts: opts}, nil
}

func fillDefaultSortedListOptions(opts *SortedListOptions) *SortedListOptions {
	if opts == nil {
		opts = &SortedListOptions{}
	}
	opts.Options = fillDefaultOptions(opts.Options)
	if opts.IteratorMaxMilliseconds <= 0 {
		opts.IteratorMaxMilliseconds = int64(time.Millisecond * 100)
	}
	return opts
}

// 更新全局计数器
func (l *SortedList) updateSeq() error {
	b, err := l.db.Get(sortedListSeqKey, nil)
	if err == leveldb.ErrNotFound {
		return l.saveSeq()
	}
	if err != nil {
		return err
	}
	l.seq = binary.BigEndian.Uint64(b)
	return nil
}

// 保存全局计数器
func (l *SortedList) saveSeq() error {
	return l.db.Put(sortedListSeqKey, l.encodeSeq(), nil)
}

func (l *SortedList) encodeSeq() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, l.seq)
	return b
}

func (l *SortedList) isSeqKey(key []byte) bool {
	return len(key) == len(sortedListSeqKey) && bytes.Compare(key, sortedListSeqKey) == 0
}

// 编码key
func (l *SortedList) encodeKey(score int64, value interface{}) ([]byte, error) {
	vb, err := l.opts.Encoder(value)
	if err != nil {
		return nil, err
	}
	flag := byte('-')
	if score >= 0 {
		flag = byte('=')
	}
	sb := make([]byte, 8)
	binary.PutVarint(sb, score)
	b := append([]byte{flag}, sb...)
	b = append(b, l.encodeSeq()...)
	b = append(b, vb...)
	return b, nil
}

// 解码key
func (l *SortedList) decodeKey(b []byte, value interface{}) (score int64, err error) {
	if len(b) < 17 {
		return 0, fmt.Errorf("SortedList.decodeKey() fail: %+v", b)
	}
	score, _ = binary.Varint(b[1:9])
	vb := b[17:]
	if err := jsoniter.Unmarshal(vb, value); err != nil {
		return 0, err
	}
	return score, nil
}

func (l *SortedList) releaseCacheIterator() {
	if l.cacheIterator != nil {
		l.cacheIterator.Release()
		l.cacheIterator = nil
	}
}

func (l *SortedList) reopenCacheIterator() (iterator.Iterator, error) {
	l.releaseCacheIterator()
	l.cacheIterator = l.db.NewIterator(nil, nil)
	l.cacheIterator.First()
	ts := getCurrentMilliseconds()
	l.cacheIteratorTimestamp = ts + l.opts.IteratorMaxMilliseconds
	return l.cacheIterator, nil
}

func (l *SortedList) getCacheIterator() (iterator.Iterator, error) {
	if l.cacheIterator == nil {
		return l.reopenCacheIterator()
	} else if l.cacheIteratorTimestamp < getCurrentMilliseconds() {
		return l.reopenCacheIterator()
	} else if !l.cacheIterator.Next() {
		return l.reopenCacheIterator()
	}
	return l.cacheIterator, nil
}

func (l *SortedList) Add(score int64, value interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.seq++
	b, err := l.encodeKey(score, value)
	if err != nil {
		return err
	}
	if err := l.db.Put(b, nil, nil); err != nil {
		return err
	}
	return l.saveSeq()
}

// 取得第一个元素
func (l *SortedList) First(maxScore int64, value interface{}) (score int64, ok bool, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	iter, err := l.getCacheIterator()
	if err != nil {
		return 0, false, err
	}
	if !iter.Valid() {
		return 0, false, nil
	}
	key := iter.Key()
	if l.isSeqKey(key) {
		return 0, false, nil
	}
	score, err = l.decodeKey(key, value)
	if err != nil {
		l.cacheIterator.Prev()
		return 0, false, err
	}
	if score <= maxScore {
		if err := l.db.Delete(key, nil); err != nil {
			l.cacheIterator.Prev()
			return 0, false, nil
		}
		return score, true, nil
	} else {
		l.cacheIterator.Prev()
		return 0, false, nil
	}
}

// 队列元素数量
func (l *SortedList) Size() (size int, err error) {
	iter := l.db.NewIterator(nil, nil)
	defer iter.Release()
	if !iter.First() {
		return 0, nil
	}
	for {
		if !l.isSeqKey(iter.Key()) {
			size++
		}
		if !iter.Next() {
			break
		}
	}
	return size, nil
}

// 关闭数据库
func (l *SortedList) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.releaseCacheIterator()
	return l.db.Close()
}

// 获得当前毫秒时间戳
func getCurrentMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
