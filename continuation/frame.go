package continuation

type FrameMessage int

const (
	FrameStatusSleep FrameMessage = iota
	FrameStatusYield
	FrameStatusNext
	FrameStatusReturn
	FrameStatusThrow
)

// 帧栈
type Frame struct {
	continuation *Continuation
	channel      chan FrameMessage

	step  int
	local interface{}

	done   bool
	result interface{}
	error  error
}

// 本地变量
func (f *Frame) Local() interface{} {
	return f.local
}

// 是否已执行完成
func (f *Frame) IsDone() bool {
	return f.done
}

// 产生结果，下一步继续执行当前段
func (f *Frame) Yield(value interface{}) {
	f.local = value
	f.channel <- FrameStatusYield
}

// 执行下一段
func (f *Frame) Next(value interface{}) {
	nextStep := f.step + 1
	if nextStep >= len(f.continuation.segments) {
		f.Return(value)
		return
	}

	f.step = nextStep
	f.local = value
	f.channel <- FrameStatusNext
}

// 返回结果
func (f *Frame) Return(value interface{}) {
	f.done = true
	f.result = value
	f.channel <- FrameStatusReturn
}

// 抛出异常
func (f *Frame) Throw(err error) {
	f.done = false
	f.error = err
	f.channel <- FrameStatusThrow
}
