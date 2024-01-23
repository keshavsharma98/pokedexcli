# pokedexcli
A REPL in GO

# For executing Tests

 ```js
 go test ./...
 ```

# For build and run

 ```js
 go build && ./pokedexcli
```

The Pokedex CLI is a powerful command-line tool designed for Pokemon enthusiasts. Elevate your Pokemon journey with this user-friendly interface, offering a range of features to explore and manage your Pokemon encounters. Here's a concise overview of its capabilities:

  - `help`: Displays a help message, guiding command usage and features.
  - `exit`: Gracefully exit the Pokedex, concluding your Pokemon exploration session.
  - `map`: Displays names of the following location areas in the Pokemon world.
  - `mapb`: Displays names of previous location areas.
  - `explore <location_area>`: Lists Pokemon in a specific area, providing crucial information for your adventures.
  - `catch <pokemon_name>`: Attempt to catch a Pokemon and add it to your Pokedex for a comprehensive collection.
  - `inspect <pokemon_name>`: View detailed information about a caught Pokemon, including characteristics, abilities, and stats.
  - `pokedex`: View all Pokemon in your Pokedex, effortlessly managing and tracking your collection.


pokecache in an internal directory. This package will be responsible for all of our caching logic.
