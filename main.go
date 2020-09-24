package main

import (
	"fmt"
	"os"

	"github.com/FrancescoIlario/gocg/cmd"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func run() error {
	return cmd.Execute()
}
