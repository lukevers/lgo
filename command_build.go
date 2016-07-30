package main

import (
	"github.com/urfave/cli"
)

func CommandBuild(c *cli.Context) error {
	// Compile lgo files to go
	err := CommandCompile(c)
	if err != nil {
		return err
	}

	// TODO: build go files
	return nil
}
