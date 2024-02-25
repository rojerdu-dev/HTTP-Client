package client

import (
	"errors"
	"fmt"
)

var (
	ErrFetchingPokemon = errors.New("failed to fetch Pokemon")
)

type PokemonFetchErr struct {
	Message    string
	StatusCode int
}

func (e PokemonFetchErr) Error() string {
	return fmt.Sprintf("Failed to fetch Pokemon: %s with statuscode %d", e.Message, e.StatusCode)
}
