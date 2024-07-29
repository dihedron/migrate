package main

import (
	"log/slog"
	"os"

	"github.com/dihedron/migrate_snapshot_repo/commands"
	"github.com/jessevdk/go-flags"
)

func main() {
	options := commands.All{}

	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		slog.Error("failure parsing command line", "error", err)
		os.Exit(1)
	}
}
