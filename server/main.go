package main

import (
	"log"

	"beverage-grpc/server/server"
)

func main() {
	if err := server.NewServer().Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
