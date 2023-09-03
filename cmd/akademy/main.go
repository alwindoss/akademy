package main

import (
	"log"

	"github.com/alwindoss/akademy"
	"github.com/alwindoss/akademy/internal/server"
)

func main() {
	cfg := akademy.Config{
		Port: 8080,
	}
	err := server.Run(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("exiting akademy")
}
