package main

import (
    "fmt"
    "os"
)

type Command struct {
    name string
    description string
    callback func() error
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
    }
}

func commandHelp() error {
    for _, command := range getCommands() {
        fmt.Printf("# %s: %s\n", command.name, command.description);    
    }
    return nil
}

func commandExit() error {
    os.Exit(0)
    return nil
}
