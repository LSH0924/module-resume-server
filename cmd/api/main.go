package main

import (
	"log"

	"module.resume/internal/container"
)

func main() {
	c, err := container.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
