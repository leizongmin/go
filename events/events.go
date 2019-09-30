package events

import (
	"reflect"
	"runtime"
	"sync"
)

type EventEmitter struct {
	mu     sync.RWMutex
	events map[string][]EventListener
}

type EventListener = func(args ...interface{})

// 新建事件触发器
func New() *EventEmitter {
	return &EventEmitter{}
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

// 触发事件
func (e *EventEmitter) EmitEvent(name string, args ...interface{}) int {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.events == nil {
		return 0
	}
	listeners, exist := e.events[name]
	if !exist {
		return 0
	}
	for _, f := range listeners {
		f(args...)
	}
	return len(listeners)
}

func getFuncFullName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
