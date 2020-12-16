package main

import (
	"fmt"
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

func main() {
	var grappleConfiguration, configurationError = getDefaultConfiguration()

	if configurationError != nil {
		log.Fatal(configurationError)
		return
	}

	app := &cli.App{
		Name:  "Grapple",
		Usage: "Git Hooks made easy!",
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "Says hello!",
				Action: func(c *cli.Context) error {
					_, err := fmt.Println("Hello! We hope you are enjoying Grapple!")
					return err
				},
			},
			{
				Name:  "install",
				Usage: "Configures local Git Hooks to adhere to the '.grapple' configuration file",
				Action: func(c *cli.Context) error {
					message, err := grappleConfiguration.Install()
					fmt.Println(message)
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
