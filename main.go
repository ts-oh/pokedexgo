package main

import "github.com/ts-oh/pokedexgo/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)

}
