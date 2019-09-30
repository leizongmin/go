package continuation

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/leizongmin/go-common-libs/typeUtils"
)

func TestNew(t *testing.T) {
	doSomething := New().ContinueSegment(func(frame *Frame) {
		fmt.Println("step1")
		time.Sleep(time.Millisecond * 200)
		local := typeUtils.MustToMap(frame.Local())
		local["step1"] = true
		frame.Next(local)
	}).ContinueSegment(func(frame *Frame) {
		fmt.Println("step2")
		time.Sleep(time.Millisecond * 200)
		local := typeUtils.MustToMap(frame.Local())
		local["step2"] = true
		frame.Next(local)
	}).ContinueSegment(func(frame *Frame) {
		fmt.Println("step3")
		time.Sleep(time.Millisecond * 200)
		local := typeUtils.MustToMap(frame.Local())
		local["step3"] = true
		frame.Next(local)
	})

	local := map[string]interface{}{"hello": "world"}
	f, err := doSomething.Call(local)
	assert.NoError(t, err)

	go func() {
		time.Sleep(time.Millisecond * 300)
		result, err := doSomething.Sleep(f)
		fmt.Println("sleep", result, err)
		assert.Nil(t, result)
		assert.Nil(t, err)
	}()

	status, result, err := doSomething.Wait(f, nil)
	fmt.Println(status, result, err)
	assert.Equal(t, FrameStatusSleep, status)
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
	status, result, err = doSomething.Wait(f, nil)
	fmt.Println(status, result, err)
	assert.Equal(t, FrameStatusReturn, status)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	result2 := typeUtils.MustToMap(result)
	assert.Equal(t, "world", result2["hello"])
	assert.Equal(t, true, result2["step1"])
	assert.Equal(t, true, result2["step2"])
	assert.Equal(t, true, result2["step3"])
}

func TestFrame_Yield(t *testing.T) {
	doSomething := New().ContinueSegment(func(frame *Frame) {
		count := typeUtils.MustToInt(frame.Local())
		if count < 10 {
			frame.Yield(count + 1)
		} else {
			frame.Next(count)
		}
	}).ContinueSegment(func(frame *Frame) {
		frame.Return(frame.Local())
	})

	frame, err := doSomething.Call(0)
	assert.NoError(t, err)
	i := 0
	for i < 10 {
		status, result, err := doSomething.Wait(frame, nil)
		assert.NoError(t, err)
		assert.Equal(t, FrameStatusYield, status)
		assert.Equal(t, i+1, result)
		i++
		doSomething.Resume(frame)
	}

	status, result, err := doSomething.Wait(frame, nil)
	assert.NoError(t, err)
	assert.Equal(t, FrameStatusReturn, status)
	assert.Equal(t, 10, result)
}

func TestFrame_Throw(t *testing.T) {
	doSomething := New().ContinueSegment(func(frame *Frame) {
		frame.Next(nil)
	}).ContinueSegment(func(frame *Frame) {
		panic("some error")
	})

	frame, err := doSomething.Call(nil)
	assert.NoError(t, err)

	status, result, err := doSomething.Wait(frame, nil)
	assert.Equal(t, FrameStatusThrow, status)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "some error", err.Error())

	fmt.Println(err)
}

func TestFrame_Manual(t *testing.T) {
	doSomething := New().Segment(func(frame *Frame) {
		frame.Next(123)
	}).Segment(func(frame *Frame) {
		frame.Return(456)
	})

	frame, err := doSomething.Call(nil)
	assert.NoError(t, err)

	status, result, err := doSomething.Wait(frame, nil)
	assert.Equal(t, FrameStatusNext, status)
	assert.Nil(t, err)
	assert.Equal(t, 123, result)

	doSomething.Resume(frame)
	status, result, err = doSomething.Wait(frame, nil)
	assert.Equal(t, FrameStatusReturn, status)
	assert.Nil(t, err)
	assert.Equal(t, 456, result)
	assert.True(t, frame.IsDone())
}

func TestFrame_CheckContinue(t *testing.T) {
	doSomething := New().ContinueSegment(func(frame *Frame) {
		frame.Next(123)
	}).ContinueSegment(func(frame *Frame) {
		frame.Return(456)
	})

	frame, err := doSomething.Call(nil)
	assert.NoError(t, err)

	status, result, err := doSomething.Wait(frame, func(frame *Frame) bool {
		return false
	})
	assert.Equal(t, FrameStatusNext, status)
	assert.Nil(t, err)
	assert.Equal(t, 123, result)

	doSomething.Resume(frame)
	status, result, err = doSomething.Wait(frame, func(frame *Frame) bool {
		return false
	})
	assert.Equal(t, FrameStatusReturn, status)
	assert.Nil(t, err)
	assert.Equal(t, 456, result)
	assert.True(t, frame.IsDone())
}

func TestFrame_CheckContinue2(t *testing.T) {
	doSomething := New().ContinueSegment(func(frame *Frame) {
		frame.Next(123)
	}).ContinueSegment(func(frame *Frame) {
		frame.Return(456)
	})

	frame, err := doSomething.Call(nil)
	assert.NoError(t, err)

	status, result, err := doSomething.Wait(frame, func(frame *Frame) bool {
		return true
	})
	assert.Equal(t, FrameStatusReturn, status)
	assert.Nil(t, err)
	assert.Equal(t, 456, result)
	assert.True(t, frame.IsDone())
}
