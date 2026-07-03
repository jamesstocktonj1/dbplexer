package main

import (
	"os"

	"github.com/jamesstocktonj1/dbplexer/internal/command"
)

func main() {
	if err := command.RootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
