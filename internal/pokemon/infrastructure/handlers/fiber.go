package handlers

import "github.com/solrac97gr/go-hexa/internal/pokemon/domain/ports"

type PokemonHandlers struct {
	app ports.PokemonApplication
}

var _ ports.PokemonHandlers = &PokemonHandlers{}

func NewPokemonHandlers(app ports.PokemonApplication) *PokemonHandlers {
	return &PokemonHandlers{
		app: app,
	}
}
