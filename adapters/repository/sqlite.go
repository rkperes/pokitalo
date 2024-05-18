package repository

import (
	"errors"

	"github.com/rkperes/pokitalo/internal/core/domain"
	"github.com/rkperes/pokitalo/internal/core/ports"
)

// enforce interface is implemented
var _ ports.Repository = &SQLite{}

type SQLite struct {
}

func NewSQLite() *SQLite {
	return &SQLite{}
}

func (s *SQLite) GetPokemonByID(int) (domain.Pokemon, error) {
	return domain.Pokemon{}, errors.New("not implemented")
}
