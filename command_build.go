package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func CommandBuild(c *cli.Context) error {
	file, err := os.Open("examples/basic/main.lgo")
	if err != nil {
		return err
	}

	p, err := NewParser(file).Parse()
	if err != nil {
		return err
	}

	fmt.Println(p)

	return nil
}
