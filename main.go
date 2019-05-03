package main

import (
	"log"

	"pumpkin/factories"
)

func main() {
	server := factories.NewServer()
	log.Fatal(server.Run())
}
