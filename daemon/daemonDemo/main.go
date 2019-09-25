package main

import (
	"log"

	"github.com/leizongmin/go-common-libs/daemon"
)

func main() {
	daemon.Run("my-example-service", "no description", nil, func() {
		log.Println("Server started...")
	}, func() error {
		log.Println("Shutdown...")
		return nil
	})
}
