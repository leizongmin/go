package lazycache

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type LazyFunc func() (interface{}, error)

type Lazy struct {
	statusLock        sync.Mutex
	maxAge            time.Duration
	lastUpdate        time.Time
	updating          atomic.Value
	data              atomic.Value
	loadFn            LazyFunc
	autoRefreshCtx    context.Context
	autoRefreshCancel context.CancelFunc
}

// 创建LazyCache
func New(maxAge time.Duration, loadFn LazyFunc) *Lazy {
	c := &Lazy{
		maxAge: maxAge,
		loadFn: loadFn,
	}
	c.updating.Store(false)
	return c
}

// 加载数据
func (c *Lazy) loadData() (interface{}, error) {
	data, err := c.loadFn()
	if err != nil {
		return nil, err
	}
	c.statusLock.Lock()
	defer c.statusLock.Unlock()
	c.data.Store(data)
	c.lastUpdate = time.Now()
	return data, nil
}

// 尝试刷新数据
func (c *Lazy) tryRefresh(useMaxAge bool) {
	if c.updating.Load().(bool) {
		return
	}
	if useMaxAge && time.Now().Sub(c.lastUpdate) < c.maxAge {
		return
	}
	c.updating.Store(true)
	go func() {
		defer c.updating.Store(false)
		_, _ = c.loadData()
	}()
}

// 强制重新加载数据
func (c *Lazy) ForceLoad() (interface{}, error) {
	return c.loadData()
}

// 获取数据，如果没有缓存则返回nil
func (c *Lazy) Get() interface{} {
	go c.tryRefresh(true)
	return c.data.Load()
}

// 自动刷新
func (c *Lazy) StartAutoRefresh(interval time.Duration) {
	c.autoRefreshCtx, c.autoRefreshCancel = context.WithCancel(context.TODO())
	go func() {
		for {
			select {
			case <-time.After(interval):
				c.tryRefresh(false)
			case <-c.autoRefreshCtx.Done():
				break
			}
		}
	}()
}

// 停止自动刷新
func (c *Lazy) StopAutoRefresh() {
	if c.autoRefreshCancel != nil {
		c.autoRefreshCancel()
	}
}
