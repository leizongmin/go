package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, true, isCall)
	assert.Equal(t, 123456, argv)
}
