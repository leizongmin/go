package continuation

import (
	"fmt"
	"testing"
	"time"

	"github.com/leizongmin/go-common-libs/typeUtils"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	doSomething := New().Segment(func(frame *Frame) {
		fmt.Println("step1")
		time.Sleep(time.Millisecond * 200)
		local := typeUtils.MustToMap(frame.Local())
		local["step1"] = true
		frame.Next(local)
	}).Segment(func(frame *Frame) {
		fmt.Println("step2")
		time.Sleep(time.Millisecond * 200)
		local := typeUtils.MustToMap(frame.Local())
		local["step2"] = true
		frame.Next(local)
	}).Segment(func(frame *Frame) {
		fmt.Println("step3")
		time.Sleep(time.Millisecond * 200)
		local := typeUtils.MustToMap(frame.Local())
		local["step3"] = true
		frame.Next(local)
	})

	local := map[string]interface{}{"hello": "world"}
	f, err := doSomething.Call(0, local)
	assert.NoError(t, err)

	go func() {
		time.Sleep(time.Millisecond * 300)
		result, err := doSomething.Sleep(f)
		fmt.Println("sleep", result, err)
		assert.Nil(t, result)
		assert.Nil(t, err)
	}()

	sleep, result, err := doSomething.Wait(f)
	fmt.Println(sleep, result, err)
	assert.True(t, sleep)
	assert.Nil(t, result)
	assert.Nil(t, err)
	assert.False(t, f.IsDone())

	fmt.Printf("%+v\n", f)
	fmt.Printf("%+v\n", local)

	data, err := doSomething.Dump(f)
	assert.NoError(t, err)
	fmt.Println(string(data))
	assert.Equal(t, `{"step":1,"local":{"hello":"world","step1":true}}`, string(data))

	err = doSomething.Restore(data)
	assert.NoError(t, err)
	sleep, result, err = doSomething.Wait(f)
	fmt.Println(sleep, result, err)
	assert.False(t, sleep)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	result2 := typeUtils.MustToMap(result)
	assert.Equal(t, "world", result2["hello"])
	assert.Equal(t, true, result2["step1"])
	assert.Equal(t, true, result2["step2"])
	assert.Equal(t, true, result2["step3"])
}
