package main

import (
	"github.com/urfave/cli"
	"os/exec"
)

func CommandBuild(c *cli.Context) error {
	// Compile lgo files to go
	err := CommandCompile(c)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "build")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
