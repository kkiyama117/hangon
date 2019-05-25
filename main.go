package main

import (
	"log"

	"pumpkin/codes/factories"
)

func main() {
	server := factories.NewServer()
	log.Fatal(server.Run(":3000"))
}
