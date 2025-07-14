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
	catchChance := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if catchChance > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	fmt.Printf("You caught %s!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the enspect command.")
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}

