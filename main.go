package main

import (
	"fmt"
	"os"

	"github.com/hookboy/source/cli"

	_ "github.com/hookboy/source/hookboy/aply"
	"github.com/hookboy/source/hookboy/app"
	_ "github.com/hookboy/source/hookboy/conf/source"
	_ "github.com/hookboy/source/hookboy/prep"
	_ "github.com/hookboy/source/hookboy/prep/generators"
	_ "github.com/hookboy/source/hookboy/prep/generators/explicit"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	builder := app.GetBuilder()

	if err := cli.RunApp(os.Args, os.Stdout, builder); err != nil {
		fmt.Fprintf(os.Stderr, "Hookboy Error :(\n| ==> %s\n", err)
		os.Exit(exitFail)
	}
}
