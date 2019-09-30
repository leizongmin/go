package continuation

import (
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

// 函数段
type SegmentFunc func(frame *Frame)

// 函数段
type Segment struct {
	Func     SegmentFunc
	Continue bool
}

type Continuation struct {
	segments []Segment
}

// 创建分段式函数
func New() *Continuation {
	return &Continuation{}
}

// 增加函数分段，每执行完一段自动暂停，需要手动 Resume()
func (c *Continuation) Segment(f SegmentFunc) *Continuation {
	c.segments = append(c.segments, Segment{Func: f, Continue: false})
	return c
}

// 增加函数分段，每执行完一段会自动执行下一段
func (c *Continuation) ContinueSegment(f SegmentFunc) *Continuation {
	c.segments = append(c.segments, Segment{Func: f, Continue: true})
	return c
}

// 从头开始执行函数
func (c *Continuation) Call(local interface{}) (*Frame, error) {
	return c.CallFromStep(0, local)
}

// 调用指定分段的函数
func (c *Continuation) CallFromStep(step int, local interface{}) (*Frame, error) {
	if step >= len(c.segments) {
		return nil, fmt.Errorf("invalid step %d (0~%d)", step, len(c.segments))
	}
	frame := c.NewFrame(step, local)
	segment := c.segments[step]
	go segment.Func(frame)
	return frame, nil
}

// 等待执行结果，如果 checkContinue 不为空则通过此函数判断是否自动继续执行下一分段（忽略分段原本的continue属性属性）
func (c *Continuation) Wait(frame *Frame, checkContinue func(frame *Frame) bool) (status FrameMessage, result interface{}, err error) {
	for {
		status = <-frame.channel
		switch status {
		case FrameStatusSleep:
			return status, nil, nil
		case FrameStatusYield:
			return status, frame.local, nil
		case FrameStatusNext:
			segment := c.segments[frame.step]
			continueNext := false
			if checkContinue != nil {
				continueNext = checkContinue(frame)
			} else if segment.Continue {
				continueNext = true
			}
			if continueNext {
				callCurrentStep(frame)
			} else {
				return status, frame.local, nil
			}
		case FrameStatusReturn:
			return status, frame.result, frame.error
		case FrameStatusThrow:
			return status, frame.result, frame.error
		}
	}
}

// 调用指定分段
func callCurrentStep(frame *Frame) {
	segment := frame.continuation.segments[frame.step]
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				frame.Throw(xerrors.Errorf("%s", err))
			}
		}()
		segment.Func(frame)
	}()
}

// 创建新的帧栈
func (c *Continuation) NewFrame(step int, local interface{}) *Frame {
	return &Frame{continuation: c, step: step, local: local, channel: make(chan FrameMessage, 0)}
}

// 休眠
func (c *Continuation) Sleep(frame *Frame) (result interface{}, err error) {
	frame.channel <- FrameStatusSleep
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
