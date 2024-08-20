package repositories

import "github.com/solrac97gr/go-hexa/internal/pokemon/domain/ports"

type PokemonRepo struct {
}

var _ ports.PokemonRepository = &PokemonRepo{}

func NewPokemonHandlers() *PokemonRepo {
	return &PokemonRepo{}
}
