package app

import (
	"connectivity-checker/app/services"
	"fmt"

	"github.com/urfave/cli"
)

func Setup() *cli.App {
	fmt.Println()
	app := cli.NewApp()
	app.Name = "Connectivity Checker"
	app.Usage = "Check a server connection"
	app.Authors = []cli.Author{
		{
			Name:  "Fillipe Meireles",
			Email: "fillipehlealmeireles@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "check",
			Usage: "Check connection",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "luspew.com",
				},
				cli.StringFlag{
					Name:  "port",
					Value: "80",
				},
			},
			Action: services.CheckServer,
		},
	}

	return app
}
