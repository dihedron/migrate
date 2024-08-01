package commands

import (
	"log/slog"
	"strings"

	"github.com/dihedron/rawdata"
)

type Format int8

const (
	Text Format = iota
	YAML
	JSON
)

func (f *Format) UnmarshalFlag(value string) error {
	switch strings.ToLower(value) {
	case "text":
		*f = Text
	case "json":
		*f = JSON
	case "yaml":
		*f = YAML
	}
	return nil
}

type Input struct {
	Data interface{}
}

func (i *Input) UnmarshalFlag(value string) error {
	var err error
	i.Data, err = rawdata.Unmarshal(value)
	if err != nil {
		slog.Error("cannot unmarshal input data", "error", err)
	}
	return err
}

// Command collects the base set of options in common to all commands.
type Command struct{}

type FormattedCommand struct {
	Command

	// Format represents the output format for the command
	Format Format `short:"F" long:"format" choice:"json" choice:"yaml" choice:"text" optional:"true" default:"text"` //lint:ignore SA5008 multiple choices
}
