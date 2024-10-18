package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Config struct {
    next *string
    prev *string
}

var config Config = Config{ next: nil, prev: nil }

func start() {
    fmt.Println("##################################")
    fmt.Println("#                                #")
    fmt.Println("#       Welkome to Pokedex       #")
    fmt.Println("#                                #")
    fmt.Println("##################################")

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("pokedex > ")
        scanner.Scan()
        input := scanner.Text()
        fmt.Println()

        err := execute(input)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func execute(input string) error {
    commands :=  getCommands()
    
    commandAndArgs := strings.Split(input, " ")
    command, ok := commands[commandAndArgs[0]]
    if !ok {
        return fmt.Errorf("Unknown command, use 'help' to see list of commands")
    }

    err := command.callback(commandAndArgs[1:], &config)
    if err != nil {
        return err
    }

    return nil
}
