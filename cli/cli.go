package cli

import (
	"fmt"
	"os"
	"path"


	log "github.com/Sirupsen/logrus"
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
    
    // logs
	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stderr)
		level, err := log.ParseLevel(c.String("log-level"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.SetLevel(level)

		// If a log level wasn't specified and we are running in debug mode,
		// enforce log-level=debug.
		if !c.IsSet("log-level") && !c.IsSet("l") && c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}