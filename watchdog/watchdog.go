package watchdog

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
	"time"
)

var CommandName = "watchdog"

// 使用默认方法，直接从命令行提取参数
// Examples:
// func main() {
//     watchdog.Main()
// }
func Main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: %s start|stop example_program", CommandName)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("get working directory failed: %s", err)
	}

	opType := os.Args[1]
	programPath := path.Join(workingDir, os.Args[2])

	Run(workingDir, opType, programPath)
}

// 执行watchdog
// workingDir 当前工作目录
// opType 命令，start 或者 stop
// programPath 二进制程序文件路径
// 执行start成功，在当前工作目录生成 watchdog.pid 文件
func Run(workingDir string, opType string, programPath string) {

	log.Printf("program path: %s", programPath)
	binPidFile := programPath + ".pid"
	log.Printf("program pid file: %s", binPidFile)
	watchdogPidFile := path.Join(workingDir, CommandName+".pid")

	if opType == "stop" {
		tryKillProcess(readPidFromFile(watchdogPidFile))
		tryKillProcess(readPidFromFile(binPidFile))
		log.Printf("done")
		return
	}
	if opType != "start" {
		log.Fatalf("usage: %s start|stop example_program", CommandName)
	}

	watchdogPid := fmt.Sprintf("%d", os.Getpid())
	log.Printf("watchdog pid: %s", watchdogPid)
	log.Printf("watchdog pid file: %s", watchdogPidFile)
	if err := ioutil.WriteFile(watchdogPidFile, []byte(watchdogPid), 0644); err != nil {
		log.Printf("write pid file failed: %s", err)
	}

	for {
		log.Printf("starting %s...\n\n\n", programPath)
		cmd := exec.Command(programPath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			log.Printf("start failed: %s", err)
		}

		pid := fmt.Sprintf("%d", cmd.Process.Pid)
		log.Printf("current pid: %s", pid)
		if err := ioutil.WriteFile(binPidFile, []byte(pid), 0644); err != nil {
			log.Printf("write pid file failed: %s", err)
		}

		if err := cmd.Wait(); err != nil {
			log.Printf("run failed: %s", err)
		}

		restartAfter := time.Second * 2
		log.Printf("restart after %s\n\n\n", restartAfter)
		time.Sleep(restartAfter)
	}
}

func readPidFromFile(file string) int {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("read pid from file %s failed: %s", file, err)
		return 0
	}
	pid, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		log.Printf("read pid from file %s failed: %s", file, err)
		return 0
	}
	return int(pid)
}

func tryKillProcess(pid int) {
	err := syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		log.Printf("try to kill process #%d failed: %s", pid, err)
	}
}
