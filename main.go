package main

import (
	"fmt"
	"github.com/rQxwX3/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println("failed to read config file:", err)
	}

	err = config.SetUser("rQxwX3")
	if err != nil {
		fmt.Println("failed to set current user:", err)
	}

	conf, err = config.Read()
	fmt.Println(conf)
}
