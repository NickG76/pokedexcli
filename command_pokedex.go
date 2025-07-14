package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("You have not caught any pokemon yet")
	}
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println(" - ", pokemon.Name)
	}
	return nil
}
