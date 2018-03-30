package main

import (
	"os"

	"github.com/andreiburuntia/cupsbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}