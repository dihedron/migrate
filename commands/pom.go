package commands

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/dihedron/migrate/xmlstream"
)

type Key string

const VersionKey Key = "version"

type POM struct {
	Command
	DryRun    bool   `short:"d" long:"dry-run" description:"Whether to simulate the execution by outputting the modified file(s) to STDOUT." optional:"yes"`
	Extension string `short:"e" long:"extension" description:"The file extension for the backup copy of the file(s)." optional:"yes" default:"bkp"`
	Args      struct {
		Files []string
	} `positional-args:"yes" required:"yes"`
}

func HandlePOM(ctx context.Context, stack xmlstream.Stack, tkn xml.Token) (context.Context, []xml.Token, error) {
	slog.Debug("handling token", "type", fmt.Sprintf("%T", tkn), "stack", stack)
	if stack.Len() == 2 && stack.At(-2) == "project" && stack.At(-1) == "version" {
		token, ok := tkn.(xml.CharData)
		if ok {
			return context.WithValue(ctx, VersionKey, strings.TrimSpace(string(token))), []xml.Token{xml.CharData("${revision}")}, nil
		}
	} else if stack.Len() == 3 && stack.At(-3) == "project" && stack.At(-2) == "parent" && stack.At(-1) == "version" {
		token, ok := tkn.(xml.CharData)
		if ok {
			return context.WithValue(ctx, VersionKey, strings.TrimSpace(string(token))), []xml.Token{xml.CharData("${revision}")}, nil
		}
	} else if stack.Len() > 1 && stack.At(-2) == "project" && stack.At(-1) == "properties" {
		token, ok := tkn.(xml.EndElement)
		if ok {
			t := ctx.Value(VersionKey)
			if v, ok := t.(string); ok {
				return ctx, []xml.Token{
					xml.CharData("    "),
					xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "revision",
						},
					},
					xml.CharData(v + "${versionSuffix}"),
					xml.EndElement{
						Name: xml.Name{
							Space: "",
							Local: "revision",
						},
					},
					xml.CharData("\n    "),
					token,
				}, nil
			}
		}
	}
	return ctx, []xml.Token{tkn}, nil
}

const separator = "----------------------------------------------------------------"

func (cmd *POM) Execute(args []string) error {

	var result error

	cmd.Extension = strings.TrimLeft(cmd.Extension, ".")

	for i, original := range cmd.Args.Files {
		if !cmd.DryRun {
			backup := fmt.Sprintf("%s.%s", original, cmd.Extension)
			slog.Info("creating backup copy of file", "original", original, "backup", backup)
			os.Rename(original, backup)
			func(original, backup string) error {
				file, err := os.Create(original)
				if err != nil {
					slog.Error("error opening file", "path", original, "error", err)
					return err
				}
				defer file.Close()
				err = xmlstream.ParseFile(backup, file, HandlePOM)
				if err != nil {
					result = errors.Join(result, err)
				}
				return err
			}(original, backup)
		} else {
			if i == 0 {
				fmt.Fprintln(os.Stdout, separator)
			}
			err := xmlstream.ParseFile(original, os.Stdout, HandlePOM)
			if err != nil {
				slog.Error("error parsing file", "path", original, "error", err)
				result = errors.Join(result, err)
			}
			fmt.Fprintln(os.Stdout, separator)
		}
	}

	return result
}
