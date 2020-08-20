package benchmark

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

type Result struct {
	Count int           `json:"count"`
	Spent time.Duration `json:"spent"`
	TPS   int           `json:"tps"`
}

// 执行测试
func Do(count int, loop func()) Result {
	st := time.Now()
	for n := 0; n < count; n++ {
		loop()
	}
	et := time.Now()
	spent := et.Sub(st) / time.Duration(count)
	return Result{Count: count, Spent: spent, TPS: int(time.Second / spent)}
}

// 执行测试并打印结果
func DoAndPrint(count int, title string, loop func()) {
	r := Do(count, loop)
	fmt.Printf("%s\n\t%+v TPS:%+v\n", title, r.Spent, r.TPS)
}

// 并行执行测试
func DoParallel(concurrent int, count int, loop func()) Result {
	if concurrent < 1 {
		concurrent = runtime.NumCPU()
	}
	parallelLoop := func(count int, wg *sync.WaitGroup) {
		for n := 0; n < count; n++ {
			loop()
		}
		wg.Done()
	}
	st := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(concurrent)
	pc := int(math.Ceil(float64(count) / float64(concurrent)))
	for c := 0; c < concurrent; c++ {
		go parallelLoop(pc, &wg)
	}
	wg.Wait()
	et := time.Now()
	spent := et.Sub(st) / time.Duration(pc*concurrent)
	return Result{Count: count, Spent: spent, TPS: int(time.Second / spent)}
}

// 并行执行测试并打印结果
func DoParallelAndPrint(concurrent int, count int, title string, loop func()) {
	r := DoParallel(concurrent, count, loop)
	fmt.Printf("%s\n\t%+v TPS:%+v\n", title, r.Spent, r.TPS)
}
