package main

import (
	"fmt"
	"log"

	"github.com/ts-oh/pokedexgo/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
