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

type Settings struct {
	Command
	Args struct {
		Files []string
	} `positional-args:"yes" required:"yes"`
}

func HandleSettingsXml(ctx context.Context, stack xmlstream.Stack, tkn xml.Token) (context.Context, []xml.Token, error) {
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

func (cmd *Settings) Execute(args []string) error {

	var result error

	for _, path := range cmd.Args.Files {
		err := xmlstream.ParseFile(path, os.Stdout, HandleSettingsXml)
		if err != nil {
			result = errors.Join(result, err)
		}
	}

	return result
}
