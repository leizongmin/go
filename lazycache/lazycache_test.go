package lazycache

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
)

func getData() (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(math.MaxInt32)
	time.Sleep(10 * time.Millisecond)
	return i, nil
}

func TestLazy(t *testing.T) {
	lazy := New(150*time.Millisecond, getData)
	c, _ := lazy.ForceLoad()
	fmt.Printf("Get c %d \n", c)
	time.Sleep(20 * time.Millisecond)
	a := lazy.Get()
	fmt.Printf("Get a %d \n", a)
	time.Sleep(20 * time.Millisecond)
	b := lazy.Get()
	fmt.Printf("Get b %d \n", b)
	if a != b {
		t.FailNow()
	}
}

func TestLazyAutoRefresh(t *testing.T) {
	lazy := New(150*time.Millisecond, getData)
	lazy.ForceLoad()
	lazy.StartAutoRefresh(time.Millisecond * 100)
	a1 := lazy.Get()
	litter.Dump(a1)
	time.Sleep(time.Millisecond * 120)
	a2 := lazy.Get()
	litter.Dump(a2)
	assert.NotEqual(t, a1, a2)
	time.Sleep(time.Millisecond * 120)
	a3 := lazy.Get()
	litter.Dump(a3)
	assert.NotEqual(t, a2, a3)
	lazy.StopAutoRefresh()
	time.Sleep(time.Millisecond * 120)
	a4 := lazy.Get()
	litter.Dump(a4)
	assert.Equal(t, a3, a4)
}

func BenchmarkLazyGet(b *testing.B) {
	lazy := New(10*time.Second, getData)
	a, _ := lazy.ForceLoad()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := lazy.Get()
		if c != a {
			b.FailNow()
		}
	}
}
