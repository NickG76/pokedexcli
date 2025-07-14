package main

import (
	"time"
	"github.com/NickG76/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Mimnute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
