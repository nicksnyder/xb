package project

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nicksnyder/xb/internal/fs"
)

type Project struct {
	Name         string   `json:"projectName"`
	XcodeVersion string   `json:"xcodeVersion"`
	Targets      []Target `json:"targets"`
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

func xcodeID(name string) string {
	digest := md5.Sum([]byte(name))
	return strings.ToUpper(hex.EncodeToString(digest[:12]))
}

var _ fs.Exportable = (*Project)(nil)
