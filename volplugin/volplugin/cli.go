package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/contiv/volplugin/volplugin"
)

var host string

func init() {
	var err error
	host, err = os.Hostname()
	if err != nil {
		panic("Could not retrieve hostname")
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "volplugin"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Turn on debug logging",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:   "master",
			Usage:  "Set the volmaster host:port",
			EnvVar: "MASTER",
			Value:  "localhost:8080",
		},
		cli.StringFlag{
			Name:   "host-label",
			Usage:  "Set the internal hostname",
			EnvVar: "HOSTLABEL",
			Value:  host,
		},
	}
	app.Action = run

	app.Run(os.Args)
}

func run(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		fmt.Printf("\nUsage: %s [tenant/driver name]\n\n", os.Args[0])
		cli.ShowAppHelp(ctx)
		os.Exit(1)
	}

	volplugin.Daemon(ctx.Args()[0], ctx.Bool("debug"), ctx.String("master"), ctx.String("host-label"))
}
