package main

import (
	"fmt"
	"os"
)

func main() {
	e, err := InitializeEvent()

	if err != nil {
		fmt.Printf("failed to create event: %s", err)
		os.Exit(2)
	}

	e.Start()
}
