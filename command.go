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
            description: "Explore location for pokemon encounters\n## args - [locationArea] (required)",
            callback: commandExplore,
        },
    }
}
