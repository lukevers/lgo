package main

import (
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"strings"
)

func CommandCompile(c *cli.Context) error {
	files, err := filepath.Glob("*.lgo")
	if err != nil {
		return err
	}

	for _, path := range files {
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			return err
		}

		parsed, err := NewParser(file).Parse()
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
