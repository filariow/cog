package main

import (
	"os"

	"github.com/FrancescoIlario/gocg/cmd"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	return cmd.Execute()
}
