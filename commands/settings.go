package commands

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

type Settings struct {
	Files []string `short:"f" long:"file" description:"The files to migrate" required:"yes"`
}

func (cmd *Settings) Execute(args []string) error {

	//var err error

	for _, path := range cmd.Files {
		func(path string) error {
			file, err := os.Open(path)
			if err != nil {
				slog.Error("error opening file", "path", path, "error", err)
				return err
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			// optionally, resize scanner's capacity for lines over 64K, see next example
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				slog.Error("error reading file", "path", path, "error", err)
				return err
			}

			return nil
		}(path)

	}
	return nil

}
