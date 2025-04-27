package http

import (
	"fmt"
	"net/http"
	"time"

	playerpb "github.com/danilopucci/opentibiagateway/internal/protogen/v1"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type HttpServer struct {
	grpcConn     *grpc.ClientConn
	playerClient playerpb.PlayerServiceClient
}

// NewHttpServer initializes the HttpServer
func NewHttpServer(grpcAddress string) (*HttpServer, error) {
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return &HttpServer{
		grpcConn:     conn,
		playerClient: playerpb.NewPlayerServiceClient(conn),
	}, nil
}

// Start runs the HTTP server
func (s *HttpServer) Start(address string) error {
	router := mux.NewRouter()

	playerHandler := NewPlayerHandler(s.playerClient)

	// Setup routes
	router.HandleFunc("/players/{id}", playerHandler.GetPlayer).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("HTTP server running on", address)
	return srv.ListenAndServe()
}

// Stop gracefully closes resources
func (s *HttpServer) Stop() error {
	if s.grpcConn != nil {
		return s.grpcConn.Close()
	}
	return nil
}
