package statistics

import (
	"math"
)

//  计数器，得到 counter
const TypeCounter = "counter"

// 采样，得到 counter, max, min, avg
const TypeSamples = "samples"

// 任意数据
const TypeData = "data"

type TagItem struct {
	Type    string      `json:"type"`    // 统计类型
	Title   string      `json:"title"`   // 标题
	Counter int32       `json:"counter"` // 计数
	Min     int32       `json:"min"`     // 最小值
	Max     int32       `json:"max"`     // 最大值
	Avg     float64     `json:"avg"`     // 平均值
	Data    interface{} `json:"data"`    // 数据
}

type ReportItem struct {
	Type    string      `json:"type"`    // 统计类型
	Tag     string      `json:"tag"`     // 标签
	Counter int32       `json:"counter"` // 计数
	Min     int32       `json:"min"`     // 最小值
	Max     int32       `json:"max"`     // 最大值
	Avg     float64     `json:"avg"`     // 平均值
	Data    interface{} `json:"data"`    // 数据
}

type Statistics struct {
	tags map[string]TagItem
}

func New() *Statistics {
	return &Statistics{tags: make(map[string]TagItem)}
}

// 获取指定标签的数据
func (s *Statistics) Get(tag string) (TagItem, bool) {
	t, ok := s.tags[tag]
	return t, ok
}

// 初始化标签
func (s *Statistics) Init(itemType string, tag string, title string) {
	s.tags[tag] = TagItem{
		Type:    itemType,
		Title:   title,
		Counter: 0,
		Min:     math.MaxInt32,
		Max:     math.MinInt32,
		Avg:     0,
		Data:    nil,
	}
}

// 增加计数
func (s *Statistics) Incr(tag string) *Statistics {
	return s.IncrN(tag, 1)
}

// 增加计数
func (s *Statistics) IncrN(tag string, n int32) *Statistics {
	item, ok := s.tags[tag]
	if ok {
		item.Counter += n
		s.tags[tag] = item
	}
	return s
}

// 减计数
func (s *Statistics) Decr(tag string) *Statistics {
	return s.DecrN(tag, 1)
}

// 减计数
func (s *Statistics) DecrN(tag string, n int32) *Statistics {
	item, ok := s.tags[tag]
	if ok {
		item.Counter -= n
		s.tags[tag] = item
	}
	return s
}

// 添加采样数据
func (s *Statistics) Add(tag string, n int32) *Statistics {
	item, ok := s.tags[tag]
	if ok {
		if n < item.Min {
			item.Min = n
		}
		if n > item.Max {
			item.Max = n
		}
		item.Avg = (float64(n)-item.Avg)/float64(item.Counter+1) + item.Avg
		item.Counter++
		s.tags[tag] = item
	}
	return s
}

// 设置数据
func (s *Statistics) Set(tag string, data interface{}) *Statistics {
	item, ok := s.tags[tag]
	if ok {
		item.Data = data
		s.tags[tag] = item
	}
	return s
}

// 获得当前报告
func (s *Statistics) Report() []ReportItem {
	list := make([]ReportItem, 0)
	for tag, item := range s.tags {
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
	}
	return list
}

// 清空统计信息（一般与 report() 配合使用）
func (s *Statistics) Flush(tag string, data interface{}) {
	for _, item := range s.tags {
		item.Counter = 0
		item.Min = math.MaxInt32
		item.Max = math.MinInt32
		item.Avg = 0
		item.Data = nil
	}
}
