package continuation

type frameMessage int

const (
	frameMessageSleep frameMessage = iota
	frameMessageNext
	frameMessageReturn
	frameMessageThrow
)

// 帧栈
type Frame struct {
	continuation *Continuation
	channel      chan frameMessage

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

// 执行下一段
func (f *Frame) Next(value interface{}) {
	nextStep := f.step + 1
	if nextStep >= len(f.continuation.segments) {
		f.Return(value)
		return
	}

	f.step = nextStep
	f.local = value
	f.channel <- frameMessageNext
}

// 返回结果
func (f *Frame) Return(value interface{}) {
	f.done = true
	f.result = value
	f.channel <- frameMessageReturn
}

// 抛出异常
func (f *Frame) Throw(err error) {
	f.done = false
	f.error = err
	f.channel <- frameMessageThrow
}
