package eventemitter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/leizongmin/go-common-libs/typeutil"
)

func TestNew(t *testing.T) {
	e := New()
	f1 := func(args ...interface{}) {}
	f2 := func(args ...interface{}) {}
	f3 := func(args ...interface{}) {}

	e.AddEventListener("f1", f1)
	e.AddEventListener("f2", f2)
	e.AddEventListener("f3", f3)
	assert.Equal(t, 3, len(e.events))

	e.RemoveAllEventListener("f1")
	assert.Equal(t, 2, len(e.events))
	e.RemoveEventListener("f1", f1)
	assert.Equal(t, 2, len(e.events))
	e.RemoveEventListener("f2", f1)
	assert.Equal(t, 2, len(e.events))
	e.RemoveEventListener("f2", f2)
	assert.Equal(t, 1, len(e.events))

	isCall := false
	var argv interface{}
	e.AddEventListener("call", func(args ...interface{}) {
		isCall = true
		argv = args[len(args)-1]
	})
	e.EmitEvent("call", 123456)
	time.Sleep(time.Millisecond * 100)
	assert.Equal(t, true, isCall)
	assert.Equal(t, 123456, argv)
}

func TestNew2(t *testing.T) {
	e := New()

	c1, remove1 := e.NewEventListenerChan("f1", 0)
	c2, remove2 := e.NewEventListenerChan("f1", 0)

	var c1r []int
	var c2r []int
	go func() {
		for {
			v := typeutil.MustToIntArray(<-c1)
			if len(v) < 1 {
				break
			}
			c1r = append(c1r, v[0])
		}
	}()
	go func() {
		for {
			v := typeutil.MustToIntArray(<-c2)
			if len(v) < 1 {
				break
			}
			c2r = append(c2r, v[0])
		}
	}()

	time.Sleep(time.Millisecond * 10)
	assert.Equal(t, 2, e.EmitEvent("f1", 123))
	assert.Equal(t, 2, e.EmitEvent("f1", 456))
	assert.Equal(t, 2, e.EmitEvent("f1", 789))
	assert.Equal(t, 2, e.EmitEvent("f1", 666))

	time.Sleep(time.Millisecond * 100)
	remove1()
	remove2()

	fmt.Println(c1r)
	fmt.Println(c2r)

	assert.Equal(t, []int{123, 456, 789, 666}, c1r)
	assert.Equal(t, []int{123, 456, 789, 666}, c2r)
}
