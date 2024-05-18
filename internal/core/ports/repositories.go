package ports

import (
	"github.com/rkperes/pokitalo/internal/core/domain"
)

type Repository interface {
	GetPokemonByID(int) (domain.Pokemon, error)
}

