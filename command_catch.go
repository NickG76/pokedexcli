package main

import (
	"math/rand"
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	catchChance := rand.Intn(100)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchChance > 40 {
		fmt.Printf("Caught %v successfully!\n", pokemon.Name)
	} else {
		fmt.Printf("Oh no! %v got away!\n", pokemon.Name)
	}
	return nil
}

