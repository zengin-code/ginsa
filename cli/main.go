package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zengin-code/ginsa"
	"os"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "ginsa"
	app.Usage = "Zengin Code data diff viewer"
	app.Version = ginsa.FullVersion()

	app.Before = cli.BeforeFunc(before)

	app.Commands = []cli.Command{
		diffCmd,
		tagsCmd,
	}
}

func before(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	return nil
}

func main() {
	app.Run(os.Args)
}
