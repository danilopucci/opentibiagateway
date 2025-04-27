package service

import (
	"context"

	"github.com/danilopucci/opentibiagateway/internal/domain"
	"github.com/danilopucci/opentibiagateway/internal/provider/persistence"
)

type PlayerService struct {
	PlayerRepo persistence.PlayerRepository
}

func NewPlayerService(repo persistence.PlayerRepository) *PlayerService {
	return &PlayerService{PlayerRepo: repo}
}

func (s *PlayerService) GetPlayerByID(ctx context.Context, id int) (*domain.Player, error) {
	return s.PlayerRepo.FindByID(ctx, id)
}
