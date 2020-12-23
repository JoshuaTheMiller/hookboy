package cli

import (
	"io"

	"github.com/hookboy/source/hookboy"
	"github.com/urfave/cli/v2"
)

// RunApp starts Hookboy
func RunApp(args []string, stdout io.Writer, ab hookboy.Builder) error {
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
					options := cliOptions{}

					application, err := retrieveApplication(options, ab)

					if err != nil {
						return err
					}

					var message, installErr = application.Install()

					if installErr != nil {
						return installErr
					}

					stdout.Write([]byte(message))

					return nil
				},
			},
			{
				Name:  "config",
				Usage: "Displays the current configuration",
				Subcommands: []*cli.Command{
					{
						Name:  "current",
						Usage: "Displays the current configuration",
						Action: func(c *cli.Context) error {
							options := cliOptions{}

							application, err := retrieveApplication(options, ab)

							if err != nil {
								return err
							}

							configuration, err := application.CurrentConfiguration()

							if err != nil {
								return err
							}

							niceJSON, err := serializeToNiceJSON(configuration)

							if err != nil {
								return err
							}

							stdout.Write([]byte(niceJSON))

							return nil
						},
					},
					{
						Name:  "source",
						Usage: "Displays the current configuration location",
						Action: func(c *cli.Context) error {
							options := cliOptions{}

							application, err := retrieveApplication(options, ab)

							if err != nil {
								return err
							}

							configurationLocation, err := application.ConfigurationLocation()

							if err != nil {
								return err
							}

							stdout.Write([]byte(configurationLocation))

							return nil
						},
					},
				},
			},
		},
	}

	return app.Run(args)
}

type cliOptions struct {
}

func retrieveApplication(options cliOptions, ab hookboy.Builder) (hookboy.Application, error) {
	application, err := ab.Construct()

	return application, err
}
