package grpc

import (
	gatewaypb "github.com/danilopucci/opentibiagateway/internal/protogen/v1"
	"github.com/danilopucci/opentibiagateway/internal/service"
)

type GrpcServer struct {
	gatewaypb.UnimplementedPlayerServiceServer
	playerService *service.PlayerService
}

// Constructor
func NewGrpcServer(playerService *service.PlayerService) *GrpcServer {
	return &GrpcServer{
		playerService: playerService,
	}
}
