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
	file := path.Join(os.TempDir(), fmt.Sprintf("localpersistence-list-%d", time.Now().Unix()))
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
		var values []string
		i := -100
		for i < 200 {
			i++
			var v string
			ok, err := list.First(int64(i), &v)
			assert.NoError(t, err)
			if ok {
				values = append(values, v)
				litter.Dump(v)
			}
		}
		litter.Dump(values)
		assert.Equal(t, []string{"bbb", "aaa", "ddd", "ccc", "fff", "xxx", "eee"}, values)
		size, err := list.Size()
		assert.NoError(t, err)
		assert.Equal(t, 0, size)
	}
}

func BenchmarkOpenSortedList(b *testing.B) {
	file := path.Join(os.TempDir(), fmt.Sprintf("localpersistence-list-%d", time.Now().UnixNano()))
	log.Println(file)
	list, err := OpenSortedList(file, nil)
	defer list.Close()
	assert.NoError(b, err)

	i := 0
	for i < b.N {
		i++
		assert.NoError(b, list.Add(randutil.Int63n(math.MaxInt64), time.Now().UnixNano()))
	}
}
