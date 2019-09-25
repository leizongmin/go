package daemon

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
)

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
	go onStart()

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case killSignal := <-interrupt:
			log.Printf("Got signal: %v", killSignal)
			log.Println("Stopping server...")
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

func Run(serviceName string, serviceDescription string, dependencies []string, onStart func(), onShutdown func() error) {
	srv, err := daemon.New(serviceName, serviceDescription, dependencies...)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	service := &service{srv, serviceName, serviceDescription, dependencies}
	status, err := service.Manage(onStart, onShutdown)
	if err != nil {
		log.Fatalf("Error with status %s: %s", status, err)
	}
	log.Println(status)
}
