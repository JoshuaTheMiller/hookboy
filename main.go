package main

import (
	"fmt"
	"os"

	"github.com/hookboy/source/cli"
	"github.com/hookboy/source/hookboy"

	_ "github.com/hookboy/source/hookboy/aply"
	_ "github.com/hookboy/source/hookboy/conf/source"
	_ "github.com/hookboy/source/hookboy/prep"
	_ "github.com/hookboy/source/hookboy/prep/generators"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	builder := hookboy.GetBuilder()

	if err := cli.RunApp(os.Args, os.Stdout, builder); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
