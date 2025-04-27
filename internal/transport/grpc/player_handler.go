package grpc

import (
	"context"
	"fmt"

	playerpb "github.com/danilopucci/opentibiagateway/internal/protogen/v1"
)

func (s *GrpcServer) GetPlayer(ctx context.Context, req *playerpb.GetPlayerRequest) (*playerpb.GetPlayerResponse, error) {
	player, err := s.playerService.GetPlayerByID(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, fmt.Errorf("player not found")
	}

	return &playerpb.GetPlayerResponse{
		Player: PlayerToProto(player),
	}, nil
}
