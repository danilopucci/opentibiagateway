package mysql

import (
	"context"

	"github.com/danilopucci/opentibiagateway/internal/domain"
	"github.com/danilopucci/opentibiagateway/internal/provider/persistence"

	"gorm.io/gorm"
)

type MySQLPlayerRepository struct {
	db *gorm.DB
}

func NewMySQLPlayerRepository(db *gorm.DB) *MySQLPlayerRepository {
	return &MySQLPlayerRepository{db: db}
}

var _ persistence.PlayerRepository = (*MySQLPlayerRepository)(nil) // compile-time interface check

func (r *MySQLPlayerRepository) FindByID(ctx context.Context, id int) (*domain.Player, error) {
	var entity persistence.PlayerEntity
	result := r.db.WithContext(ctx).First(&entity, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return PlayerToDomain(&entity), nil
}
