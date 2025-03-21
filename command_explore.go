package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area name")
	}
	areaName := &args[0]
	locationResp, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *areaName)
	fmt.Println("Found Pokemon:")

	for _, pe := range locationResp.PokemonEncounters {
		fmt.Println(" -", pe.Pokemon.Name)
	}
	return nil
}
