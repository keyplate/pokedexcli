package main

import "os"

func commandExit(args []string, config *Config) error {
    os.Exit(0)
    return nil
}
