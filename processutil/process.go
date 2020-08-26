package processutil

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// 无限等待
func InfiniteWait() {
	c := make(chan bool, 0)
	<-c
}

// 结束程序
func Exit(code int) {
	os.Exit(code)
}

// 查询当前源文件的信息
func SourceFileInfo(skip ...int) (funcName string, fileName string, fileLine int, err error) {
	sk := 1
	if len(skip) > 0 && skip[0] > 1 {
		sk = skip[0]
	}
	var pc uintptr
	var ok bool
	pc, fileName, fileLine, ok = runtime.Caller(sk)
	if !ok {
		return "", "", 0, fmt.Errorf("N/A")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	ix := strings.LastIndex(name, ".")
	if ix > 0 && (ix+1) < len(name) {
		name = name[ix+1:]
	}
	funcName = name
	// nd, nf := filepath.Split(fileName)
	// fileName = filepath.Join(filepath.Base(nd), nf)
	return funcName, fileName, fileLine, nil
}
