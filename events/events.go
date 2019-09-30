package events

import (
	"reflect"
	"runtime"
	"sync"
)

type EventEmitter struct {
	mu     sync.RWMutex
	events map[string][]EventListener
	chans  map[string][]chan []interface{}
}

type EventListener = func(args ...interface{})

// 新建事件触发器
func New() *EventEmitter {
	return &EventEmitter{}
}

// 获取事件监听器的chan
func (e *EventEmitter) NewEventListenerChan(name string, size int) (c chan []interface{}, remove func()) {
	if e.chans == nil {
		e.chans = make(map[string][]chan []interface{})
	}
	c = make(chan []interface{}, size)
	e.chans[name] = append(e.chans[name], c)
	remove = func() {
		for i, v := range e.chans[name] {
			if v == c {
				e.chans[name] = append(e.chans[name][:i], e.chans[name][i+1:]...)
				close(c)
				break
			}
		}
		if len(e.chans[name]) < 1 {
			delete(e.chans, name)
		}
	}
	return c, remove
}

func (e *EventEmitter) ensureEventListeners(name string) []EventListener {
	if e.events == nil {
		e.events = make(map[string][]EventListener)
	}
	listeners, exist := e.events[name]
	if !exist {
		e.events[name] = make([]EventListener, 0)
		listeners = e.events[name]
	}
	return listeners
}

// 添加事件监听器
func (e *EventEmitter) AddEventListener(name string, listener EventListener) {
	e.mu.Lock()
	defer e.mu.Unlock()

	listeners := e.ensureEventListeners(name)
	e.events[name] = append(listeners, listener)
}

// 删除事件监听器
func (e *EventEmitter) RemoveEventListener(name string, listener EventListener) {
	e.mu.Lock()
	defer e.mu.Unlock()

	listeners := e.ensureEventListeners(name)
	listenerName := getFuncFullName(listener)
	for i, f := range listeners[:] {
		if getFuncFullName(f) == listenerName {
			listeners = append(listeners[:i], listeners[i+1:]...)
		}
	}
	if len(listeners) > 0 {
		e.events[name] = listeners
	} else {
		delete(e.events, name)
	}
}

// 删除所有事件监听器
func (e *EventEmitter) RemoveAllEventListener(name string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	delete(e.events, name)
}

// 触发事件，触发后立即返回，不等待监听器执行完毕。
// 但是保证同一事件的所有监听器一定是按照顺序执行的，如果前面的监听器阻塞了则会导致后面的事件也阻塞
func (e *EventEmitter) EmitEvent(name string, args ...interface{}) (count int) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.events != nil {
		listeners, exist := e.events[name]
		if exist {
			count += len(listeners)
			go func() {
				for _, f := range listeners {
					f(args...)
				}
			}()
		}
	}

	if e.chans != nil {
		chans, exist := e.chans[name]
		if exist {
			count += len(chans)
			go func() {
				for _, c := range chans {
					c <- args
				}
			}()
		}
	}

	return count
}

func getFuncFullName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
