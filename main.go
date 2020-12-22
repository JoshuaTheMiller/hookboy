package main

import (
	"fmt"
	"os"

	"github.com/hookboy/source/cli"
	"github.com/hookboy/source/hookboy"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	var configuration, err = hookboy.GetDefaultConfiguration()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	if err := cli.RunApp(os.Args, os.Stdout, configuration); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
