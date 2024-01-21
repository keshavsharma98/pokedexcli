package pokeapis

import "math/rand"

const baseExperienceThreshold = 30

func (c *Client) CatchPokemon(pokemonName string) (bool, error) {

	pokemonInfo, err := c.GetPokemonInfo(pokemonName)
	if err != nil {
		return false, err
	}

	if c.user.Pokedex == nil {
		c.user.Pokedex = make(map[string]Pokemon)
	}

	if rand.Intn(pokemonInfo.BaseExperience) > baseExperienceThreshold {
		c.user.Pokedex[pokemonName] = Pokemon{
			pokemonInfo.Name,
			pokemonInfo.Height,
			pokemonInfo.Weight,
			pokemonInfo.Stats,
			pokemonInfo.Types,
		}
		return true, nil
	}

	return false, nil
}

func (c *Client) InspectPokemon(pokemonName string) (bool, Pokemon, error) {
	pokemon, ok := c.user.Pokedex[pokemonName]
	return ok, pokemon, nil
}

func (c *Client) GetAllPokedexPokemons() ([]string, error) {
	pokemons := []string{}
	for _, pokemon := range c.user.Pokedex {
		pokemons = append(pokemons, pokemon.Name)
	}
	return pokemons, nil
}
