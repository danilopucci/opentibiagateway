package persistence

import (
	"context"

	"github.com/danilopucci/opentibiagateway/internal/domain"
)

type PlayerRepository interface {
	FindByID(ctx context.Context, id int) (*domain.Player, error)
}
