package main

import (
	"fmt"
	"os"

	"github.com/hookboy/hookboy/cli"
	"github.com/hookboy/hookboy/runner"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	var configuration, err = runner.GetDefaultConfiguration()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	if err := cli.RunApp(os.Args, os.Stdout, configuration); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
