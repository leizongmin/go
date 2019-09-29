package continuation

import (
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// 函数段
type Segment func(frame *Frame)

type Continuation struct {
	segments []Segment
}

// 创建分段式函数
func New(segments ...Segment) *Continuation {
	return &Continuation{segments: segments}
}

// 增加函数分段
func (c *Continuation) Segment(segment Segment) *Continuation {
	c.segments = append(c.segments, segment)
	return c
}

// 调用指定分段的函数
func (c *Continuation) Call(step int, local interface{}) (*Frame, error) {
	if step >= len(c.segments) {
		return nil, fmt.Errorf("invalid step %d (0~%d)", step, len(c.segments))
	}
	frame := c.NewFrame(step, local)
	segment := c.segments[step]
	go segment(frame)
	return frame, nil
}

// 等待执行结果
func (c *Continuation) Wait(frame *Frame) (sleep bool, result interface{}, err error) {
	for {
		switch <-frame.channel {
		case frameMessageSleep:
			return true, nil, nil
		case frameMessageNext:
			callCurrentStep(frame)
		case frameMessageReturn:
			return false, frame.result, frame.error
		case frameMessageThrow:
			return false, frame.result, frame.error
		}
	}
}

// 调用指定分段
func callCurrentStep(frame *Frame) {
	segment := frame.continuation.segments[frame.step]
	go segment(frame)
}

// 创建新的帧栈
func (c *Continuation) NewFrame(step int, local interface{}) *Frame {
	return &Frame{continuation: c, step: step, local: local, channel: make(chan frameMessage, 0)}
}

// 休眠
func (c *Continuation) Sleep(frame *Frame) (result interface{}, err error) {
	frame.channel <- frameMessageSleep
	return frame.result, frame.error
}

// 继续执行
func (c *Continuation) Resume(frame *Frame) {
	callCurrentStep(frame)
}

// 输出当前执行状态数据
func (c *Continuation) Dump(frame *Frame) ([]byte, error) {
	if frame.IsDone() {
		return nil, AlreadyDoneError
	}
	d := dumpFrameData{Step: frame.step, Local: frame.local}
	return jsoniter.Marshal(&d)
}

// 恢复执行状态数据并继续执行
func (c *Continuation) Restore(data []byte) error {
	d := dumpFrameData{}
	err := jsoniter.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	frame := c.NewFrame(d.Step, d.Local)
	c.Resume(frame)
	return nil
}

type dumpFrameData struct {
	Step  int         `json:"step"`
	Local interface{} `json:"local"`
}

// 已经执行完成
var AlreadyDoneError = errors.New("already done")
