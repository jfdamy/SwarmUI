package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:      "serve",
			ShortName: "s",
			Usage:     "Serve the api and the web uid",
			Flags: []cli.Flag{},
			Action: serve,
		},
		{
			Name:      "autoscaling",
			ShortName: "a",
			Usage:     "run the autoscaling service",
			Flags: []cli.Flag{},
			Action: runautoscaling,
		},
	}
)