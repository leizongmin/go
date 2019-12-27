package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/leizongmin/go/daemon"
)

func main() {
	daemon.SetLogFunc(func(format string, a ...interface{}) {
		fmt.Printf("[Custom Log Func] "+format+"\n", a...)
	})
	daemon.Run("my-example-service", "no description", nil, func() {
		go func() {
			log.Println("Server started...")
			for {
				log.Println("Running...")
				time.Sleep(time.Second)
			}
		}()
	}, func() error {
		log.Println("Stopping...")
		time.Sleep(time.Second * 3)
		return errors.New("something")
	})
}
