package localpersistence

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"

	"github.com/leizongmin/go/typeutil"
)

func TestOpenList(t *testing.T) {
	file := generateTempPath()
	log.Println(file)
	list, err := OpenList(file, nil)
	defer list.Close()
	assert.NoError(t, err)
	litter.Dump(list)
	{
		size, err := list.Size()
		assert.NoError(t, err)
		assert.Equal(t, 0, size)
	}
	{
		i := 0
		for i < 10 {
			i++
			err := list.AddToLast(typeutil.H{"a": i})
			assert.NoError(t, err)
		}
		size, err := list.Size()
		assert.NoError(t, err)
		assert.Equal(t, 10, size)
	}
	{
		i := 0
		for i < 10 {
			i++
			err := list.AddToFirst(typeutil.H{"b": i})
			assert.NoError(t, err)
		}
		size, err := list.Size()
		assert.NoError(t, err)
		assert.Equal(t, 20, size)
	}
	{
		i := 0
		for i < 20 {
			i++
			value := typeutil.H{}
			ok, err := list.RemoveFirst(&value)
			assert.NoError(t, err)
			assert.True(t, ok)
			litter.Dump(value)
		}
	}
	{
		assert.NoError(t, list.AddToLast("a"))
		assert.NoError(t, list.AddToLast("b"))
		assert.NoError(t, list.AddToLast("c"))
		assert.NoError(t, list.AddToFirst("x"))
		assert.NoError(t, list.AddToFirst("y"))
		assert.NoError(t, list.AddToFirst("z"))
		var ret []string
		for {
			var v string
			ok, err := list.RemoveFirst(&v)
			assert.NoError(t, err)
			if !ok {
				break
			}
			ret = append(ret, v)
		}
		litter.Dump(ret)
		assert.Equal(t, []string{"z", "y", "x", "a", "b", "c"}, ret)
	}
}

func TestListConcurrency(t *testing.T) {
	file := generateTempPath()
	log.Println(file)
	list, err := OpenList(file, nil)
	assert.NoError(t, err)
	defer list.Close()

	count := 10000
	go func() {
		i := 0
		for i < count {
			i++
			assert.NoError(t, list.AddToFirst(fmt.Sprintf("key1_%d", i)))
		}
		fmt.Println("key1 done")
	}()
	go func() {
		i := 0
		for i < count {
			i++
			assert.NoError(t, list.AddToLast(fmt.Sprintf("key2_%d", i)))
		}
		fmt.Println("key2 done")
	}()
	go func() {
		i := 0
		for i < count {
			i++
			s := ""
			_, err := list.RemoveFirst(&s)
			assert.NoError(t, err)
		}
		fmt.Println("key3 done")
	}()
	go func() {
		i := 0
		for i < count {
			i++
			s := ""
			_, err := list.RemoveLast(&s)
			assert.NoError(t, err)
		}
		fmt.Println("key4 done")
	}()
	go func() {
		i := 0
		for i < count {
			i++
			_, err := list.Size()
			assert.NoError(t, err)
		}
		fmt.Println("key5 done")
	}()
	time.Sleep(10 * time.Second)
}
