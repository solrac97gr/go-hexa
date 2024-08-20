package application

import "github.com/solrac97gr/go-hexa/internal/pokemon/domain/ports"

type PokemonApp struct {
	repo ports.PokemonRepository
}

var _ ports.PokemonApplication = &PokemonApp{}

func NewPokemonApp(repo ports.PokemonRepository) *PokemonApp {
	return &PokemonApp{
		repo: repo,
	}
}
