package sqlutil

import "log"

var isDebug = false

func EnableDebug() {
	isDebug = true
}

func DisableDebug() {
	isDebug = false
}

func debugf(format string, args ...interface{}) {
	if isDebug {
		logDebugf(format, args...)
	}
}

func warnf(format string, args ...interface{}) {
	if isDebug {
		logWarnf(format, args...)
	}
}

type LogFunction = func(format string, args ...interface{})

var logDebugf LogFunction
var logWarnf LogFunction

func init() {
	logDebugf = log.Printf
	logWarnf = log.Printf
}

// 设置 DEBUG 打印日志函数
func SetLogDebugFunction(f LogFunction) {
	logDebugf = f
}

// 设置 WARN 打印日志函数
func SetLogWarnFunction(f LogFunction) {
	logWarnf = f
}
