package daemon

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
)

var logf = log.Printf

// 设置日志记录函数
func SetLogFunc(f func(format string, a ...interface{})) {
	logf = f
}

// service has embedded daemon
type service struct {
	daemon.Daemon

	serviceName        string
	serviceDescription string
	dependencies       []string
}

// Manage by daemon commands or run the daemon
func (s *service) Manage(onStart func(), onShutdown func() error) (string, error) {
	usage := fmt.Sprintf("Usage: %s install | remove | start | stop | status", s.serviceName)

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return s.Install()
		case "remove":
			return s.Remove()
		case "start":
			return s.Start()
		case "stop":
			return s.Stop()
		case "status":
			return s.Status()
		default:
			return usage, nil
		}
	}

	// Do something, call your goroutines, etc

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// 启动服务
	onStart()

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case killSignal := <-interrupt:
			logf("Got signal: %v", killSignal)
			logf("Stopping server...")
			if err := onShutdown(); err != nil {
				return "Daemon shutdown server failed", err
			}
			if killSignal == os.Interrupt {
				return "Daemon was interrupted by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
}

// 启动服务，一般在 main() 函数中执行
// serviceName 服务名称
// serviceDescription 服务介绍
// dependencies 服务依赖项，一般为 nil
// onStart 启动服务函数，如果函数内有阻塞的代码（比如监听服务器），需要自己创建 goroutine
// onShutdown 关闭服务函数，用于接收到关闭信号后执行相应的清理，当函数结束时进程将退出
func Run(serviceName string, serviceDescription string, dependencies []string, onStart func(), onShutdown func() error) {
	srv, err := daemon.New(serviceName, serviceDescription, dependencies...)
	if err != nil {
		logf("Error: %s", err)
		os.Exit(1)
	}
	service := &service{srv, serviceName, serviceDescription, dependencies}
	status, err := service.Manage(onStart, onShutdown)
	if err != nil {
		logf("Status %s: %s", status, err)
		os.Exit(1)
	}
	logf(status)
}
