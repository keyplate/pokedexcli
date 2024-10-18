package main

import "fmt"

func commandHelp(config *Config) error {
    for _, command := range getCommands() {
        fmt.Printf("# %s: %s\n", command.name, command.description);    
    }
    return nil
}
