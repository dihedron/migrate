package commands

import (
	"errors"
	"os"

	"github.com/dihedron/migrate/xmlstream"
)

type Settings struct {
	//Files []string `short:"f" long:"file" description:"The files to migrate" required:"yes"`
	// Example of positional arguments
	Args struct {
		Files []string
	} `positional-args:"yes" required:"yes"`
}

func (cmd *Settings) Execute(args []string) error {

	var result error

	for _, path := range cmd.Args.Files {
		err := xmlstream.ParseFile(path, os.Stdout, nil)
		if err != nil {
			result = errors.Join(result, err)
		}
	}

	return result
}
