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

const Version Key = "version"

type POM struct {
	//Files []string `short:"f" long:"file" description:"The files to migrate" required:"yes"`
	// Example of positional arguments
	Args struct {
		Files []string
	} `positional-args:"yes" required:"yes"`
}

func DetectVersion(ctx context.Context, stack xmlstream.Stack, tkn xml.Token) (context.Context, []xml.Token, error) {
	slog.Debug("handling token", "type", fmt.Sprintf("%T", tkn), "stack", stack)
	if stack.Len() == 2 && stack.At(-2) == "project" && stack.At(-1) == "version" {
		token, ok := tkn.(xml.CharData)
		if ok {
			return context.WithValue(ctx, Version, strings.TrimSpace(string(token))), []xml.Token{xml.CharData("${revision}")}, nil
		}
	} else if stack.Len() == 3 && stack.At(-3) == "project" && stack.At(-2) == "parent" && stack.At(-1) == "version" {
		token, ok := tkn.(xml.CharData)
		if ok {
			return context.WithValue(ctx, Version, strings.TrimSpace(string(token))), []xml.Token{xml.CharData("${revision}")}, nil
		}
	} else if stack.Len() > 1 && stack.At(-2) == "project" && stack.At(-1) == "properties" {
		token, ok := tkn.(xml.EndElement)
		if ok {
			//fmt.Fprintf(os.Stderr, "ADDING PROPERTY: %s\n", strings.TrimSpace(string(token)))
			t := ctx.Value(Version)
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

// func AddProperty(ctx context.Context, stack []string, tkn xml.Token) (context.Context, []xml.Token, error) {
// 	slog.Debug("handling token", "type", fmt.Sprintf("%T", tkn), "stack", stack)
// 	if len(stack) > 1 && stack[len(stack)-2] == "project" && stack[len(stack)-1] == "properties" {
// 		token, ok := tkn.(xml.EndElement)
// 		if ok && token.Name.Local == "properties" {
// 			fmt.Fprintf(os.Stderr, "VERSION: %s\n", strings.TrimSpace(string(token)))
// 			return context.WithValue(ctx, "version", strings.TrimSpace(string(token))), xml.CharData("${revision}"), nil
// 		}
// 	}
// 	return ctx, []xml.Token{tkn}, nil
// }

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
