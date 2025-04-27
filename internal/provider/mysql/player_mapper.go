package mysql

import (
	"github.com/danilopucci/opentibiagateway/internal/domain"
	"github.com/danilopucci/opentibiagateway/internal/provider/persistence"
)

func PlayertoDomain(entity *persistence.PlayerEntity) *domain.Player {
	return &domain.Player{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func PlayertoEntity(p *domain.Player) *persistence.PlayerEntity {
	return &persistence.PlayerEntity{
		ID:   p.ID,
		Name: p.Name,
	}
}
