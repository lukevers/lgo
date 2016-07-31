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

	shared := []cli.Flag{
		cli.StringFlag{
			Name:   "extension, ext",
			Value:  "lgo",
			Usage:  "file extensions to parse",
			EnvVar: "FILE_EXT",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "build",
			Usage:  "compile lgo files and build go files",
			Action: CommandBuild,
			Flags:  shared,
		},
		{
			Name:   "compile",
			Usage:  "compile lgo files to go files",
			Action: CommandCompile,
			Flags:  shared,
		},
	}

	app.Run(os.Args)
}
