package commands

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"log/slog"
	"os"
)

type Pom struct {
	Files []string `short:"f" long:"file" description:"The files to migrate" required:"yes"`
}

func (cmd *Pom) Execute(args []string) error {

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
			var data bytes.Buffer
			// optionally, resize scanner's capacity for lines over 64K, see next example
			for scanner.Scan() {
				fmt.Println(scanner.Text())
				data.Write(scanner.Bytes())
			}

			if err := scanner.Err(); err != nil {
				slog.Error("error reading file", "path", path, "error", err)
				return err
			}

			p := POM{}

			xml.Unmarshal(data.Bytes(), &p)

			// j, _ := json.MarshalIndent(p, "", "  ")
			// fmt.Printf("JSON:\n %s\n", string(j))

			x, _ := xml.MarshalIndent(p, "", "  ")
			fmt.Printf("XML:\n %s\n", string(x))

			return nil
		}(path)

	}
	return nil

}

type POM struct {
	XMLName        xml.Name `xml:"project" json:"project,omitempty"`
	XMLNS          string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	XSI            string   `xml:"xsi,attr" json:"xsi,omitempty"`
	SchemaLocation string   `xml:"schemaLocation,attr" json:"schemaLocation,omitempty"`
	ModelVersion   struct {
		Text string `xml:",chardata" json:"text,omitempty" yaml:"modelVersion,omitempty"`
	} `xml:"modelVersion" json:"modelversion,omitempty" yaml:",omitempty,inline"`
	Parent *struct {
		// Text    string `xml:",chardata" json:"text,omitempty"`
		GroupId struct {
			Text string `xml:",chardata" json:"text,omitempty" yaml:"groupid,omitempty"`
		} `xml:"groupId,omitempty" json:"groupid,omitempty" yaml:",omitempty,inline"`
		ArtifactId struct {
			Text string `xml:",chardata" json:"text,omitempty" yaml:"artifactId,omitempty"`
		} `xml:"artifactId,omitempty" json:"artifactid,omitempty" yaml:",omitempty,inline"`
		Version struct {
			Text string `xml:",chardata" json:"text,omitempty" yaml:"version,omitempty"`
		} `xml:"version,omitempty" json:"version,omitempty" yaml:",omitempty,inline"`
		Packaging struct {
			Text string `xml:",chardata" json:"text,omitempty" yaml:"packaging,omitempty"`
		} `xml:"packaging,omitempty" json:"packaging,omitempty" yaml:",omitempty,inline"`
		RelativePath struct {
			Text string `xml:",chardata" json:"text,omitempty" yaml:"relativePath,omitempty"`
		} `xml:"relativePath,omitempty" json:"relativePath,omitempty" yaml:",omitempty,inline"`
	} `xml:"parent,omitempty" json:"parent,omitempty"`
	GroupId struct {
		Text string `xml:",chardata" json:"text,omitempty" yaml:"groupId,omitempty"`
	} `xml:"groupId" json:"groupid,omitempty" yaml:",omitempty,inline"`
	ArtifactId struct {
		Text string `xml:",chardata" json:"text,omitempty" yaml:"artifactId,omitempty"`
	} `xml:"artifactId" json:"artifactid,omitempty" yaml:",omitempty,inline"`
	Version struct {
		Text string `xml:",chardata" json:"text,omitempty" yaml:"version,omitempty"`
	} `xml:"version" json:"version,omitempty" yaml:",omitempty,inline"`
	Packaging struct {
		Text string `xml:",chardata" json:"text,omitempty" yaml:"packaging,omitempty"`
	} `xml:"packaging" json:"packaging,omitempty" yaml:",omitempty,inline"`
	Modules struct {
		// Text   string `xml:",chardata" json:"text,omitempty"`
		Module []struct {
			Text string `xml:",chardata" json:"text,omitempty" yaml:"name,omitempty"`
		} `xml:"module" json:"module,omitempty" yaml:"module,omitempty"`
	} `xml:"modules" json:"modules,omitempty" yaml:"modules,omitempty"`
}
