package main

import (
	"log"
	"os"
	"server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags|log.Lshortfile)
	srv := server.CreateServer(logger)

	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}
}
