package project

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/nicksnyder/xb/internal/fs"
)

type Project struct {
	Name         string   `json:"projectName"`
	XcodeVersion string   `json:"xcodeVersion"`
	Targets      []Target `json:"targets"`
}

type Target struct {
	Name                 string   `json:"name"`
	Type                 string   `json:"type"`
	Platform             string   `json:"platform"`
	Sources              []string `json:"sources"`
	Dependencies         []string `json:"dependencies"`
	TestableDependencies []string `json:"testableDependencies"`
}

func NewFromConfigFile(filename string) (*Project, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var project Project
	if err = json.Unmarshal(bytes, &project); err != nil {
		return nil, err
	}
	return &project, err
}

func (p *Project) ExportTo(destination string) error {
	projectFiles := fs.Directory{
		Name: p.directory(),
		Children: []fs.Exportable{
			p.pbxproj(),
			&fs.Directory{
				Name: "project.xcworkspace",
				Children: []fs.Exportable{
					p.xcworkspacedata(),
				},
			},
		},
	}

	return projectFiles.ExportTo(destination)
}

func (p *Project) directory() string {
	return p.Name + ".xcodeproj"
}

func (p *Project) Clean() error {
	return os.RemoveAll(p.directory())
}

var _ fs.Exportable = (*Project)(nil)
