package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Println("")
	fmt.Println("Bye 👋")
	fmt.Println("")

	os.Exit(0)

	return nil
}
