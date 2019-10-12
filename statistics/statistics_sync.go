package statistics

import (
	"math"
	"sync"
	"time"
)

// 并发安全版本
type SyncStatistics struct {
	tags sync.Map
}

func NewSync() *SyncStatistics {
	return &SyncStatistics{}
}

// 获取指定标签的数据
func (s *SyncStatistics) Get(tag string) (item TagItem, ok bool) {
	t, ok := s.tags.Load(tag)
	if !ok {
		return item, false
	}
	return t.(TagItem), true
}

// 初始化标签
func (s *SyncStatistics) Init(itemType string, tag string, title string) {
	s.tags.Store(tag, TagItem{
		Type:    itemType,
		Title:   title,
		Counter: 0,
		Min:     math.MaxInt32,
		Max:     math.MinInt32,
		Avg:     0,
		Data:    nil,
	})
}

// 增加计数
func (s *SyncStatistics) Incr(tag string) *SyncStatistics {
	return s.IncrN(tag, 1)
}

// 增加计数
func (s *SyncStatistics) IncrN(tag string, n int32) *SyncStatistics {
	item, ok := s.Get(tag)
	if ok {
		item.Counter += n
		s.tags.Store(tag, item)
	}
	return s
}

// 减计数
func (s *SyncStatistics) Decr(tag string) *SyncStatistics {
	return s.DecrN(tag, 1)
}

// 减计数
func (s *SyncStatistics) DecrN(tag string, n int32) *SyncStatistics {
	item, ok := s.Get(tag)
	if ok {
		item.Counter -= n
		s.tags.Store(tag, item)
	}
	return s
}

// 添加采样数据
func (s *SyncStatistics) Add(tag string, n int32) *SyncStatistics {
	item, ok := s.Get(tag)
	if ok {
		if n < item.Min {
			item.Min = n
		}
		if n > item.Max {
			item.Max = n
		}
		item.Avg = (float64(n)-item.Avg)/float64(item.Counter+1) + item.Avg
		item.Counter++
		s.tags.Store(tag, item)
	}
	return s
}

// 设置数据
func (s *SyncStatistics) Set(tag string, data interface{}) *SyncStatistics {
	item, ok := s.Get(tag)
	if ok {
		item.Data = data
		s.tags.Store(tag, item)
	}
	return s
}

// 获得当前报告
func (s *SyncStatistics) Report(flush bool) []ReportItem {
	list := make([]ReportItem, 0)
	s.tags.Range(func(key, value interface{}) bool {
		tag := key.(string)
		item := value.(TagItem)
		if item.Type == "counter" {
			list = append(list, ReportItem{
				Tag:     tag,
				Type:    item.Type,
				Counter: item.Counter,
			})
		} else if item.Type == "samples" {
			if item.Counter == 0 {
				list = append(list, ReportItem{
					Tag:     tag,
					Type:    item.Type,
					Counter: item.Counter,
				})
			} else if item.Counter == 1 {
				list = append(list, ReportItem{
					Tag:     tag,
					Type:    item.Type,
					Counter: item.Counter,
					Min:     int32(item.Avg),
					Max:     int32(item.Avg),
					Avg:     item.Avg,
				})
			} else {
				list = append(list, ReportItem{
					Tag:     tag,
					Type:    item.Type,
					Counter: item.Counter,
					Min:     item.Min,
					Max:     item.Max,
					Avg:     math.Round(item.Avg*10000) / 10000,
				})
			}
		} else if item.Type == "data" {
			list = append(list, ReportItem{
				Tag:  tag,
				Type: item.Type,
				Data: item.Data,
			})
		}
		return true
	})
	if flush {
		s.Flush()
	}
	return list
}

// 清空统计信息（一般与 report() 配合使用）
func (s *SyncStatistics) Flush() {
	s.tags.Range(func(key, value interface{}) bool {
		item := value.(TagItem)
		item.Counter = 0
		item.Min = math.MaxInt32
		item.Max = math.MinInt32
		item.Avg = 0
		item.Data = nil
		s.tags.Store(key, item)
		return true
	})
}

// 监听数据报告
func (s *SyncStatistics) Watch(interval time.Duration, callback func(list []ReportItem)) (cancel func()) {
	ch := make(chan bool)
	cancel = func() {
		ch <- true
	}
	go func() {
		for {
			select {
			case <-time.After(interval):
				callback(s.Report(true))
			case <-ch:
				callback(s.Report(true))
				break
			}
		}
	}()
	return cancel
}
