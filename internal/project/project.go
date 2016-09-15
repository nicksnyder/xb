package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/nicksnyder/xb/internal/fs"
)

type Project struct {
	ProjectName  string   `json:"projectName"`
	XcodeVersion string   `json:"xcodeVersion"`
	Targets      []Target `json:"targets"`
}

type Target struct {
	TargetName           string   `json:"name"`
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
	var pbxproj bytes.Buffer
	if err := pbxprojTemplate.Execute(&pbxproj, p); err != nil {
		return err
	}
	var xcworkspacedata bytes.Buffer
	if err := xcworkspacedataTemplate.Execute(&xcworkspacedata, p); err != nil {
		return err
	}

	projectFiles := fs.Directory{
		Name: p.directory(),
		Children: []fs.Exportable{
			&fs.File{
				Name:     "project.pbxproj",
				Contents: &pbxproj,
			},
			&fs.Directory{
				Name: "project.xcworkspace",
				Children: []fs.Exportable{
					&fs.File{
						Name:     "contents.xcworkspacedata",
						Contents: &xcworkspacedata,
					},
				},
			},
		},
	}

	return projectFiles.ExportTo(destination)
}

func (p *Project) directory() string {
	return p.ProjectName + ".xcodeproj"
}

func (p *Project) Clean() error {
	return os.RemoveAll(p.directory())
}

func parseTemplate(source string) *template.Template {
	return template.Must(template.New("").Parse(source))
}

var _ fs.Exportable = (*Project)(nil)
