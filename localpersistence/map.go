package localpersistence

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Map struct {
	db   *DB
	opts *Options
}

// 打开数据库
func OpenMap(file string, opts *Options) (*Map, error) {
	opts = fillDefaultOptions(opts)
	db, err := openDb(file, opts.DB)
	if err != nil {
		return nil, err
	}
	return &Map{db: db, opts: opts}, nil
}

// 获取元素值
func (m *Map) Get(key string, value interface{}) (ok bool, err error) {
	b, err := m.db.Get([]byte(key), nil)
	if err == leveldb.ErrNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if err := m.opts.Decoder(b, value); err != nil {
		return false, err
	}
	return true, nil
}

// 设置元素值
func (m *Map) Put(key string, value interface{}) (err error) {
	b, err := m.opts.Encoder(value)
	if err != nil {
		return err
	}
	if err := m.db.Put([]byte(key), b, nil); err != nil {
		return err
	}
	return nil
}

// 设置元素值
func (m *Map) Remove(key string) (err error) {
	return m.db.Delete([]byte(key), nil)
}

// 遍历所有key
func (m *Map) ForEachKey(fn func(key string) bool) (err error) {
	iter := m.db.NewIterator(nil, nil)
	defer iter.Release()
	if !iter.First() {
		return nil
	}
	for {
		if !fn(string(iter.Key())) {
			break
		}
		if !iter.Next() {
			break
		}
	}
	return nil
}

// 元素数量
func (m *Map) Size() (size int, err error) {
	iter := m.db.NewIterator(nil, nil)
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
func (m *Map) Close() error {
	return m.db.Close()
}
