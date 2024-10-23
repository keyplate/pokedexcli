package main

type Command struct {
    name string
    description string
    callback func(args []string, config *Config) error
}

func getCommands() map[string]Command {
    return map[string]Command {
        "help" : {
            name: "help",
            description: "Display a help message",
            callback: commandHelp, 
        },
        "exit" : {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "map" : {
            name: "map",
            description: "Get a list of locations",
            callback: commandMap,
        },
        "mapb" : {
            name: "mapb",
            description: "Return to previous list of locations",
            callback: commandMapb,
        },
        "explore" : {
            name: "explore",
            description: "Explore location for pokemon encounters\n## args - [locationArea]",
            callback: commandExplore,
        },
        "catch": {
            name: "catch",
            description: "Tries to capture provided pokemon\n## args - [pokemonName]",
            callback: commandCatch,
        },
        "inspect": {
            name: "inspect",
            description: "Show information about the given pokemon, if it was previously encountered\n## args - [pokemonName]",
            callback: commandInspect,
        }, 
    }
}
