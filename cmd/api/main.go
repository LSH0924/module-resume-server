package main

import (
	"log"

	"module.resume/internal/api"
)

func main() {
	r := api.MakeRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
