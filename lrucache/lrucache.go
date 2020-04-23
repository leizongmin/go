package lrucache

import (
	"container/list"
	"context"
	"sync"
	"time"
)

// LRU缓存
type LRUCache struct {
	mu                sync.RWMutex
	maxSize           int                      // 最大缓存数量
	defaultTtl        time.Duration            // 默认缓存时间
	cacheList         *list.List               // 缓存数据列表
	cache             map[string]*list.Element // 缓存数据KV表
	hits              atomicUint               // 命中次数统计
	gets              atomicUint               // 查询次数统计
	autoRefreshCtx    context.Context
	autoRefreshCancel context.CancelFunc
}

type entry struct {
	key    string
	value  interface{}
	expiry time.Time
}

// 如果 maxSize = 0 则表示不限制
func New(maxSize int, ttl time.Duration) *LRUCache {
	return &LRUCache{
		maxSize:    maxSize,
		cacheList:  list.New(),
		cache:      make(map[string]*list.Element),
		defaultTtl: ttl,
	}
}

// 缓存状态
func (c *LRUCache) Status() *Status {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return &Status{
		MaxSize: c.maxSize,
		Size:    c.cacheList.Len(),
		Gets:    c.gets.get(),
		Hits:    c.hits.get(),
	}
}

// 判断元素是否未过期
func (c *LRUCache) isAlive(ent *entry) bool {
	return ent.expiry.IsZero() || ent.expiry.After(time.Now())
}

// 查询缓存
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	needToDelete := false
	defer func() {
		c.mu.RUnlock()
		// 必须在RLock解锁后再执行删除
		if needToDelete {
			c.Delete(key)
		}
	}()
	c.gets.add(1)
	if ele, hit := c.cache[key]; hit {
		ent := ele.Value.(*entry)
		if c.isAlive(ent) {
			c.hits.add(1)
			c.cacheList.MoveToFront(ele)
			return ele.Value.(*entry).value, true
		}
		needToDelete = true
	}
	return nil, false
}

// 设置缓存，指定过期时间
func (c *LRUCache) SetEx(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var expiry time.Time
	if ttl > 0 {
		expiry = time.Now().Add(ttl)
	} else {
		expiry = time.Time{}
	}

	if ele, ok := c.cache[key]; ok {
		c.cacheList.MoveToFront(ele)
		ele.Value.(*entry).value = value
		ele.Value.(*entry).expiry = expiry
		return
	}

	ele := c.cacheList.PushFront(&entry{key: key, value: value, expiry: expiry})
	c.cache[key] = ele
	if c.maxSize != 0 && c.cacheList.Len() > c.maxSize {
		c.removeOldest()
	}
}

// 设置缓存
func (c *LRUCache) Set(key string, value interface{}) {
	c.SetEx(key, value, c.defaultTtl)
}

func (c *LRUCache) remove(ele *list.Element) {
	c.cacheList.Remove(ele)
	key := ele.Value.(*entry).key
	delete(c.cache, key)
}

// 删除缓存
func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		return
	}
	if ele, ok := c.cache[key]; ok {
		c.remove(ele)
		return
	}
}

func (c *LRUCache) removeOldest() {
	if c.cache == nil {
		return
	}
	ele := c.cacheList.Back()
	if ele != nil {
		c.remove(ele)
	}
}

// 缓存垃圾回收
func (c *LRUCache) GC() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, ele := range c.cache {
		ent := ele.Value.(*entry)
		if !c.isAlive(ent) {
			c.remove(ele)
		}
	}
}

// 自动刷新
func (c *LRUCache) StartAutoGC(interval time.Duration) {
	c.autoRefreshCtx, c.autoRefreshCancel = context.WithCancel(context.TODO())
	go func() {
		for {
			select {
			case <-time.After(interval):
				c.GC()
			case <-c.autoRefreshCtx.Done():
				break
			}
		}
	}()
}

// 停止自动刷新
func (c *LRUCache) StopAutoGC() {
	if c.autoRefreshCancel != nil {
		c.autoRefreshCancel()
	}
}
