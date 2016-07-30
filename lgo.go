package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lgo"
	app.Usage = "literate go"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:   "build",
			Usage:  "compile lgo files and build go files",
			Action: CommandBuild,
		},
		{
			Name:   "compile",
			Usage:  "compile lgo files to go files",
			Action: CommandCompile,
		},
	}

	app.Run(os.Args)
}
