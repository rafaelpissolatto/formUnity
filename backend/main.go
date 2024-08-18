package main

import (
	"context"
	"log"

	cmdserver "github.com/rafaelpissolatto/formUnity/backend/cmd/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Starting server...")

	cmdserver.RunAPIServer(ctx)

}
