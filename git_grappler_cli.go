package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli"
)

var recognizedHooks = [...]string{
	"applypatch-msg",
	"commit-msg",
	"fsmonitor-watchman",
	"post-update",
	"pre-applypatch",
	"pre-commit",
	"pre-merge-commit",
	"pre-push",
	"pre-rebase",
	"pre-receive",
	"prepare-commit-msg",
	"update"}

var actualGitHooksDir = ".git/hooks/"
var grappleCacheDir = ".grapple-cache"

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := runApp(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func runApp(args []string, stdout io.Writer) error {
	var grappleConfiguration, configurationError = getDefaultConfiguration()

	if configurationError != nil {
		log.Fatal(configurationError)
		return configurationError
	}

	app := &cli.App{
		Writer: stdout,
		Name:   "Grapple",
		Usage:  "Git Hooks made easy!",
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "Says hello!",
				Action: func(c *cli.Context) error {
					var message = "Hello! We hope you are enjoying Grapple!"
					_, err := stdout.Write([]byte(message))
					return err
				},
			},
			{
				Name:  "install",
				Usage: "Configures local Git Hooks to adhere to the '.grapple' configuration file",
				Action: func(c *cli.Context) error {
					message, err := grappleConfiguration.Install()
					stdout.Write([]byte(message))
					return err
				},
			},
		},
	}

	err := app.Run(args)
	if err != nil {
		return err
	}

	return nil
}
