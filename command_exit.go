package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config) error {
	fmt.Println("Bye! 👋")
	fmt.Println("")

	os.Exit(0)

	return nil
}
