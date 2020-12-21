package cli

import (
	"io"

	"github.com/hookboy/source/runner"
	"github.com/urfave/cli"
)

// RunApp starts Hookboy
func RunApp(args []string, stdout io.Writer, application runner.HookBoy) error {
	// var grappleConfiguration, configurationError = getDefaultConfiguration()

	// if configurationError != nil {
	// 	return configurationError
	// }

	app := &cli.App{
		Writer: stdout,
		Name:   "Grapple",
		Usage:  "Git Hooks made easy!",
		Commands: []cli.Command{
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
					message, err := application.Install()
					stdout.Write([]byte(message))
					return err
				},
			},
		},
	}

	return app.Run(args)
}
