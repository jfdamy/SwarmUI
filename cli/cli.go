package cli

import (
	"fmt"
	"os"
	"path"
    "log"

	"github.com/codegangsta/cli"
)

//Run run the command
func Run() {

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "Swarm UI is a toolkit with an Api (REST), a WebUI and a autoscaling service for docker/docker swarm"
	app.Version = "0.0.1"

	app.Author = "Jeff"
	app.Email = "jeanfrancois.damy@gmail.com"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},

		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: fmt.Sprintf("Log level (options: debug, info, warn, error, fatal, panic)"),
		},
	}

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}