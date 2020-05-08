package localpersistence

import (
	"fmt"
	"log"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
)

func TestOpenMap(t *testing.T) {
	file := generateTempPath()
	log.Println(file)
	m, err := OpenMap(file, nil)
	defer m.Close()
	assert.NoError(t, err)
	litter.Dump(m)
	{
		size, err := m.Size()
		assert.NoError(t, err)
		assert.Equal(t, 0, size)
	}
	{
		i := 0
		for i < 10 {
			i++
			assert.NoError(t, m.Put(fmt.Sprintf("index%2d", i), i))
		}
		size, err := m.Size()
		assert.NoError(t, err)
		assert.Equal(t, 10, size)
	}
	{
		var values []int
		err := m.ForEachKey(func(key string) bool {
			var v int
			ok, err := m.Get(key, &v)
			assert.NoError(t, err)
			assert.True(t, ok)
			values = append(values, v)
			assert.NoError(t, m.Remove(key))
			return true
		})
		assert.NoError(t, err)

		litter.Dump(values)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, values)

		size, err := m.Size()
		assert.NoError(t, err)
		assert.Equal(t, 0, size)
	}
}
