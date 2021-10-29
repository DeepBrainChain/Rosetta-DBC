package main

import (
	"os"

	"github.com/fatih/color"
	"rosetta-dbc/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}
