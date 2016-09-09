package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

//VERSION is the version for this software
var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "recipes"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = func(c *cli.Context) error {
		logrus.Info("start")
		CreateDB()
		return nil
	}

	app.Run(os.Args)
}
