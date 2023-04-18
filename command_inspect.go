package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no input: please provide pokemon name")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("🔍 Inspecting: %s\n", pokemon.Name)
	fmt.Println("")
	fmt.Printf("📏 Height: %v\n", pokemon.Height)
	fmt.Printf("🏋️ Weight: %v\n", pokemon.Weight)
	fmt.Printf("📊 Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s", typ.Type.Name)
	}

	return nil
}
