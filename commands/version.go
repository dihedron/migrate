package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"runtime/debug"

	"github.com/dihedron/migrate/output"
)

// Version is the command that prints information about the application
// or plugin to the console.
type Version struct {
	FormattedCommand
}

func (cmd *Version) Execute(args []string) error {
	switch cmd.Format {
	case Text:
		fmt.Printf("%s v%s.%s.%s (%s/%s built with %s on %s)\n", Name, VersionMajor, VersionMinor, VersionPatch, GoOS, GoArch, GoVersion, BuildTime)
	default:
		type (
			git struct {
				Commit   string `json:"commit" yaml:"commit"`
				Time     string `json:"time" yaml:"time"`
				Modified string `json:"modified" yaml:"modified"`
			}
			golang struct {
				Version string `json:"version" yaml:"version"`
				OS      string `json:"os" yaml:"os"`
				Arch    string `json:"arch" yaml:"arch"`
			}
			version struct {
				Major string `json:"major" yaml:"major"`
				Minor string `json:"minor" yaml:"minor"`
				Patch string `json:"patch" yaml:"patch"`
			}
			info struct {
				Name        string  `json:"name" yaml:"name"`
				Description string  `json:"description" yaml:"description"`
				Copyright   string  `json:"copyright" yaml:"copyright"`
				License     string  `json:"license" yaml:"license"`
				LicenseURL  string  `json:"licenseurl" yaml:"license-url"`
				Git         git     `json:"git" yaml:"git"`
				Go          golang  `json:"golang" yaml:"golang"`
				Version     version `json:"version" yaml:"version"`
				BuildTime   string  `json:"buildtime" yaml:"buildtime"`
			}
		)

		result := info{
			Name:        Name,
			Description: Description,
			Copyright:   Copyright,
			License:     License,
			LicenseURL:  LicenseURL,
			Git: git{
				Commit:   GitCommit,
				Time:     GitTime,
				Modified: GitModified,
			},
			Go: golang{
				Version: GoVersion,
				OS:      GoOS,
				Arch:    GoArch,
			},
			Version: version{
				Major: VersionMajor,
				Minor: VersionMinor,
				Patch: VersionPatch,
			},
			BuildTime: BuildTime,
		}
		switch cmd.Format {
		case JSON:
			s, _ := output.ToPrettyJSON(result)
			fmt.Println(s)
		case YAML:
			s, _ := output.ToYAML(result)
			fmt.Print(s)
		}
	}

	return nil
}

// NOTE: some of these variables are populated at compile time by using the -ldflags
// linker flag:
//
//	$> go build -ldflags "-X github.com/dihedron/migrate/commands.VersionMajor=$(major_ver)"
//
// in order to get the package path to the GitHash variable to use in the
// linker flag, use the nm utility and look for the variable in the built
// application symbols, then use its path in the linker flag:
//
//	$> nm ./migrate | grep VersionMajor
//	0000000000677fe0 B github.com/dihedron/migrate/commands.VersionMajor
var (
	// BuildTime is the time at which the application was built.
	BuildTime string
	// VersionMajor is the major version of the application.
	VersionMajor = "0"
	// VersionMinor is the minor version of the application.
	VersionMinor = "0"
	// VersionPatch is the patch or revision level of the application.
	VersionPatch = "0"
)

var (
	// Name is the name of the application or plugin.
	Name string = "migrate"
	// Description is a one-liner description of the application or plugin.
	Description string = "A simple configuration management experiment."
	// Copyright is the copyright clause of the application or plugin.
	Copyright string = "2024 © Andrea Funtò"
	// License is the license under which the code is released.
	License string = "MIT"
	// LicenseURL is the URL at which the license is available.
	LicenseURL string = "https://opensource.org/license/mit/"
	// GitCommit is the commit of this version of the application.
	GitCommit string
	// GitTime is the modification time associated with the Git commit.
	GitTime string
	// GitModified reports whether the repository had outstanding local changes at time of build.
	GitModified string
	// GoVersion is the version of the Go compiler used in the build process.
	GoVersion string
	// GoOS is the target operating system of the application.
	GoOS string
	// GoOS is the target architecture of the application.
	GoArch string
)

func init() {

	if Name == "" {
		Name = path.Base(os.Args[0])
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		slog.Error("no build info available")
		return
	}

	GoVersion = bi.GoVersion

	for _, setting := range bi.Settings {
		switch setting.Key {
		case "GOOS":
			GoOS = setting.Value
		case "GOARCH":
			GoArch = setting.Value
		case "vcs.revision":
			GitCommit = setting.Value
		case "vcs.time":
			GitTime = setting.Value
		case "vcs.modified":
			GitModified = setting.Value
		}
	}

	// fmt.Printf("name        : %s\n", Name)
	// fmt.Printf("description : %s\n", Description)
	// fmt.Printf("copyright   : %s\n", Copyright)
	// fmt.Printf("license     : %s\n", License)
	// fmt.Printf("license url : %s\n", LicenseURL)
	// fmt.Printf("git commit  : %s\n", GitCommit)
	// fmt.Printf("git time    : %s\n", GitTime)
	// fmt.Printf("git modified: %s\n", GitModified)
	// fmt.Printf("go version  : %s\n", GoVersion)
	// fmt.Printf("go os       : %s\n", GoOS)
	// fmt.Printf("go arch     : %s\n", GoArch)
	// fmt.Printf("major       : %s\n", VersionMajor)
	// fmt.Printf("minor       : %s\n", VersionMinor)
	// fmt.Printf("patch       : %s\n", VersionPatch)
	// fmt.Printf("build time  : %s\n", BuildTime)
}
