package main

type Command struct {
    name string
    description string
    callback func(config *Config) error
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
    }
}
