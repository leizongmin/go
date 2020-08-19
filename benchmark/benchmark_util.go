package benchmark

import (
	"log"
	"os"
	"time"
)

var spentTimeLog *log.Logger

func init() {
	spentTimeLog = log.New(os.Stderr, "[benchmark.SpentTime] ", log.LstdFlags)
}

// 打印执行函数消耗的时间
func SpentTime(msg string, f func()) time.Duration {
	spentTimeLog.Println(msg)
	start := time.Now()
	f()
	spent := time.Now().Sub(start)
	spentTimeLog.Printf("%s\t(%s)", msg, spent)
	return spent
}
