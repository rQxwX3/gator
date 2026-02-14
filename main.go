package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rQxwX3/gator/internal/config"
	"github.com/rQxwX3/gator/internal/database"
	"os"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", conf.DBurl)
	dbQueries := database.New(db)
	state := state{&conf, dbQueries}

	commands := commands{cmdmap{}}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)

	osArgs := os.Args
	if len(osArgs) < 2 {
		fmt.Println("Usage: gator <command> [arguments]")
		os.Exit(1)
	}

	cmd := command{osArgs[1], osArgs[2:]}

	err = commands.run(&state, cmd)
	if err != nil {
		fmt.Println("Error when running command:", err)
		os.Exit(1)
	}
}
