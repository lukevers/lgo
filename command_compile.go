package main

import (
	"fmt"
	"github.com/lukevers/lgo/parser"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"strings"
)

func CommandCompile(c *cli.Context) error {
	// Grab all files in the current directory
	files, err := filepath.Glob(fmt.Sprintf("*.%s", c.String("extension")))
	if err != nil {
		return err
	}

	// Compile files from lgo to go
	for _, path := range files {
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			return err
		}

		parsed, err := parser.NewParser(file).Parse()
		if err != nil {
			return err
		}

		npath := strings.Replace(path, ".lgo", ".go", 1)
		nfile, err := os.Create(npath)
		if err != nil {
			return err
		}

		nfile.WriteString(parsed)
		nfile.Close()
	}

	return nil
}
