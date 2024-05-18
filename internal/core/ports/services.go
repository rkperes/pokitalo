package ports

import "github.com/rkperes/pokitalo/internal/core/domain"

type Service interface {
	GetPokemonByID(int) (domain.Pokemon, error)
}
