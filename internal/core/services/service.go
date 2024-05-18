package services

import (
	"errors"

	"github.com/rkperes/pokitalo/internal/core/domain"
	"github.com/rkperes/pokitalo/internal/core/ports"
)

// enforce interface is implemented
var _ ports.Service = &Service{}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) GetPokemonByID(int) (domain.Pokemon, error) {
	return domain.Pokemon{}, errors.New("not implemented")

}
