package main

import (
	"math/rand"
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You might give a pokemon you want to catch")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	catchChance := rand.Intn(100)
	if catchChance < 70 {
		fmt.Printf("Caught %v successfully!", pokemon.Species.Name)
	}
	return nil
}
