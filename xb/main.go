package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nicksnyder/xb/internal/project"
)

func usage() {
	fmt.Printf(`xb generates and builds xcode projects.
Usage:
    xb [options] [command] [targets...]

Commands:
    project
Options:
    -sourceLanguage tag
        goi18n uses the strings from this language to seed the translations for other languages.
        Default: en-us
`)
}

func main() {
	flag.Usage = usage
	// sourceLanguage := flag.String("sourceLanguage", "en-us", "")
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "project":
		generateProject("build.json")
	case "":
		usage()
		os.Exit(1)
	}
}

func generateProject(configFilename string) {
	proj, err := project.NewFromConfig(configFilename)
	check("read project configuration file", err)
	err = proj.ExportTo(".")
	check("export project", err)
}

func check(action string, err error) {
	if err != nil {
		fmt.Printf("%s failed %q\n", action, err)
		os.Exit(1)
	}
}
