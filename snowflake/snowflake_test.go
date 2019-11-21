package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	// 生成节点实例
	worker, err := NewWorker(1, DefaultOptions())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", worker)

	ch := make(chan int64)
	count := 10000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := worker.GetId()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			t.Error(fmt.Sprintf("ID is not unique! id=%d", id))
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All", count, "snowflake ID Get successed!")

	for i := range m {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
