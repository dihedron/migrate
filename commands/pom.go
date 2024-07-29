package commands

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/dihedron/migrate/xmlstream"
)

type POM struct {
	//Files []string `short:"f" long:"file" description:"The files to migrate" required:"yes"`
	// Example of positional arguments
	Args struct {
		Files []string
	} `positional-args:"yes" required:"yes"`
}

func DetectVersion(stack []string, tkn xml.Token) (xml.Token, error) {
	slog.Debug("handling token", "type", fmt.Sprintf("%T", tkn), "stack", stack)
	if len(stack) > 1 && stack[len(stack)-2] == "project" && stack[len(stack)-1] == "version" {
		token, ok := tkn.(xml.CharData)
		if ok {
			fmt.Fprintf(os.Stderr, "VERSION: %s\n", strings.TrimSpace(string(token)))
		}
		return xml.CharData("${revision}"), nil
	}
	return tkn, nil
}

func (cmd *POM) Execute(args []string) error {

	var result error

	for _, path := range cmd.Args.Files {
		err := xmlstream.ParseFile(path, os.Stdout, DetectVersion)
		if err != nil {
			result = errors.Join(result, err)
		}
	}

	return result
}
