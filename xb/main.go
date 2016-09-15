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

  xb project [config]
		Generates an Xcode project from the config file.
		If config is not provided, it defaults to "build.json".

`)
}

func main() {
	flag.Usage = usage
	// sourceLanguage := flag.String("sourceLanguage", "en-us", "")
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "project":
		configFile := flag.Arg(1)
		if configFile == "" {
			configFile = "build.json"
		}
		generateProject(configFile)
	case "":
		usage()
		os.Exit(1)
	}
}

func generateProject(configFile string) {
	proj, err := project.NewFromConfigFile(configFile)
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
