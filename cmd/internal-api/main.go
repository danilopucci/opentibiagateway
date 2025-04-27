package main

import (
	"log"

	"github.com/danilopucci/opentibiagateway/internal/transport/http"
)

func main() {
	httpServer, err := http.NewHttpServer("localhost:50051") // gRPC server address
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	err = httpServer.Start(":8080") // HTTP server address
	if err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}
