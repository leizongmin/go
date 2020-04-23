package lrucache

import (
	"sync/atomic"
)

// return status of Cache
type Status struct {
	Gets    uint64 `json:"gets"` // 查询次数
	Hits    uint64 `json:"hits"` // 命中次数
	MaxSize int    `json:"max"`  // 历史最大缓存数量
	Size    int    `json:"size"` // 当前缓存数量
}

// this is a interface which defines some common functions
type Cache interface {
	Set(key string, value interface{})  // 设置缓存
	Get(key string) (interface{}, bool) // 获取缓存
	Delete(key string)                  // 删除缓存
	Status() *Status                    // 当前状态
	GC()                                // 垃圾回收
}

// An atomicUint is an int64 to be accessed atomically.
type atomicUint uint64

func (i *atomicUint) add(n uint64) {
	atomic.AddUint64((*uint64)(i), n)
}

func (i *atomicUint) get() uint64 {
	return atomic.LoadUint64((*uint64)(i))
}
