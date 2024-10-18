package main

import "fmt"

func commandHelp(args []string, config *Config) error {
    for _, command := range getCommands() {
        fmt.Printf("# %s: %s\n", command.name, command.description);    
    }
    
    return nil
}
