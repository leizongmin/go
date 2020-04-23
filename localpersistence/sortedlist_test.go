package localpersistence

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"testing"
	"time"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"

	"github.com/leizongmin/go/randutil"
)

func TestOpenSortedList(t *testing.T) {
	file := path.Join(os.TempDir(), fmt.Sprintf("localpersistence-sortedlist-%d", time.Now().UnixNano()))
	log.Println(file)
	list, err := OpenSortedList(file, nil)
	defer list.Close()
	assert.NoError(t, err)
	litter.Dump(list)
	{
		size, err := list.Size()
		assert.NoError(t, err)
		assert.Equal(t, 0, size)
	}
	{
		i := -10
		for i <= 10 {
			s1 := int64(i)
			b, err := list.encodeKey(s1, i)
			assert.NoError(t, err)
			var v int
			s2, err := list.decodeKey(b, &v)
			assert.NoError(t, err)
			fmt.Printf("%d: %+v s1=%d, s2=%d, v=%d\n", len(b), b, s1, s2, v)
			assert.Equal(t, s1, s2)
			i++
		}
	}
	{
		assert.NoError(t, list.Add(0, "aaa"))
		assert.NoError(t, list.Add(-1, "bbb"))
		assert.NoError(t, list.Add(3, "ccc"))
		assert.NoError(t, list.Add(2, "ddd"))
		assert.NoError(t, list.Add(99, "eee"))
		assert.NoError(t, list.Add(88, "fff"))
		assert.NoError(t, list.Add(88, "xxx"))
		size, err := list.Size()
		assert.NoError(t, err)
		assert.Equal(t, 7, size)
	}
	{
		type tValue struct {
			Score int
			Value string
		}
		var values []tValue
		i := -100
		for i < 200 {
			i++
			var v string
			s, ok, err := list.First(int64(i), &v)
			assert.NoError(t, err)
			if ok {
				values = append(values, tValue{Score: int(s), Value: v})
			}
		}
		litter.Dump(values)
		assert.Equal(t, []tValue{{-1, "bbb"}, {0, "aaa"}, {2, "ddd"}, {3, "ccc"}, {88, "fff"}, {88, "xxx"}, {99, "eee"}}, values)
		size, err := list.Size()
		assert.NoError(t, err)
		assert.Equal(t, 0, size)
	}
}

func BenchmarkOpenSortedList(b *testing.B) {
	file := path.Join(os.TempDir(), fmt.Sprintf("localpersistence-sortedlist-%d", time.Now().UnixNano()))
	log.Println(file)
	list, err := OpenSortedList(file, nil)
	defer list.Close()
	assert.NoError(b, err)

	i := 0
	for i < b.N {
		i++
		list.Add(randutil.Int63n(math.MaxInt32), time.Now().UnixNano())
	}
	var v int64
	for i < b.N {
		i++
		list.First(time.Now().UnixNano(), &v)
	}
}
