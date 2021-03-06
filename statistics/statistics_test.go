package statistics

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func compute(list []int32) (int32, int32, float64) {
	var max float64 = math.MinInt32
	var min float64 = math.MaxInt32
	var avg float64
	for i, v := range list {
		max = math.Max(max, float64(v))
		min = math.Min(min, float64(v))
		avg = (float64(v)-avg)/float64(i+1) + avg
	}
	avg = math.Round(avg*10000) / 10000
	return int32(min), int32(max), avg
}

func TestStatistics_Add(t *testing.T) {
	s := New()
	s.Init(TypeSamples, "a", "")
	list := []int32{8, 2, 3, 4, 3, 2, 1, 7, 6, 9, 20, 15, 12}
	min, max, avg := compute(list)
	for _, v := range list {
		s.Add("a", v)
	}
	ret := s.Report(true)
	item := ret[0]
	assert.Equal(t, TypeSamples, item.Type)
	assert.Equal(t, "a", item.Tag)
	assert.Equal(t, int32(len(list)), item.Counter)
	assert.Equal(t, max, item.Max)
	assert.Equal(t, min, item.Min)
	assert.Equal(t, avg, item.Avg)
}

func TestStatistics_Add2(t *testing.T) {
	s := New()
	s.Init(TypeSamples, "a", "")
	list := []int32{8, 2}
	min, max, avg := compute(list)
	for _, v := range list {
		s.Add("a", v)
	}
	ret := s.Report(true)
	item := ret[0]
	assert.Equal(t, TypeSamples, item.Type)
	assert.Equal(t, "a", item.Tag)
	assert.Equal(t, int32(len(list)), item.Counter)
	assert.Equal(t, max, item.Max)
	assert.Equal(t, min, item.Min)
	assert.Equal(t, avg, item.Avg)
}

func TestStatistics_Init(t *testing.T) {
	s := New()
	s.Init(TypeCounter, "aa", "")
	i := 0
	for i < 100 {
		i++
		s.Incr("aa")
	}
	ret := s.Report(true)
	item := ret[0]
	assert.Equal(t, TypeCounter, item.Type)
	assert.Equal(t, "aa", item.Tag)
	assert.Equal(t, int32(100), item.Counter)
}

func TestStatistics_Flush(t *testing.T) {
	s := New()
	s.Init(TypeCounter, "aa", "")

	s.Incr("aa")
	s.Incr("aa")
	ret1 := s.Report(true)
	item1 := ret1[0]
	assert.Equal(t, TypeCounter, item1.Type)
	assert.Equal(t, "aa", item1.Tag)
	assert.Equal(t, int32(2), item1.Counter)

	s.Incr("aa")
	s.Incr("aa")
	ret2 := s.Report(true)
	item2 := ret2[0]
	assert.Equal(t, TypeCounter, item2.Type)
	assert.Equal(t, "aa", item2.Tag)
	assert.Equal(t, int32(2), item2.Counter)
}

func TestStatistics_Watch(t *testing.T) {
	s := New()
	s.Init(TypeCounter, "aa", "")

	go func() {
		for {
			s.Incr("aa")
			s.Incr("aa")
			time.Sleep(time.Millisecond * 100)
		}
	}()

	cancel := s.Watch(time.Millisecond*500, func(list []ReportItem) {
		fmt.Printf("%+v\n", list)
	})
	time.Sleep(time.Second * 2)
	cancel()
}

func BenchmarkStatistics_Counter(b *testing.B) {
	s := New()
	s.Init(TypeCounter, "request_count", "请求次数")
	i := 0
	for i < b.N {
		i++
		s.Incr("request_count")
	}
}

func BenchmarkStatistics_Samples(b *testing.B) {
	s := New()
	s.Init(TypeSamples, "response_time", "响应时间")
	i := 0
	for i < b.N {
		i++
		s.Add("response_time", int32(i))
	}
}

func BenchmarkStatistics_Data(b *testing.B) {
	s := New()
	s.Init(TypeData, "any_data", "任意数据")
	i := 0
	data := map[string]interface{}{"a": 123, "b": 456}
	for i < b.N {
		i++
		s.Set("any_data", data)
	}
}
