package main

import (
	"fmt"
	"github.com/rQxwX3/gator/internal/config"
	"os"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		os.Exit(1)
	}

	state := state{&conf}
	commands := commands{cmdmap{}}

	commands.register("login", handlerLogin)

	osArgs := os.Args
	if len(osArgs) < 2 {
		fmt.Println("Usage: gator <command> [arguments]")
		os.Exit(1)
	}

	fmt.Println(osArgs)

	cmdname := osArgs[1]
	cmdargs := osArgs[2:]
	cmd := command{cmdname, cmdargs}

	err = commands.run(&state, cmd)
	if err != nil {
		fmt.Println("Error when running command:", err)
		os.Exit(1)
	}
}
