package main

import (
	"log"

	"github.com/bengal-dev/panel/internal/config"
)

func main() {
	if _, err := config.New(); err != nil { // пока не использую конфиг
		log.Fatalf("Cant init config: %v", err)
	}

	log.Println("Node Started Successfully")
}
