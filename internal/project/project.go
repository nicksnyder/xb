package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
		Name: p.ProjectName + ".xcodeproj",
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

var _ fs.Exportable = (*Project)(nil)

func parseTemplate(source string) *template.Template {
	return template.Must(template.New("").Parse(source))
}

var xcworkspacedataTemplate = parseTemplate(`<?xml version="1.0" encoding="UTF-8"?>
<Workspace
   version = "1.0">
   <FileRef
      location = "self:{{.ProjectName}}.xcodeproj">
   </FileRef>
</Workspace>
`)

var pbxprojTemplate = parseTemplate(`// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 48;
	objects = {

/* Begin PBXGroup section */
		0BA02E491D87571500F1E8D3 = {
			isa = PBXGroup;
			children = (
			);
			sourceTree = "<group>";
		};
/* End PBXGroup section */

/* Begin PBXProject section */
		0BA02E4A1D87571500F1E8D3 /* Project object */ = {
			isa = PBXProject;
			attributes = {
				LastUpgradeCheck = 0800;
			};
			buildConfigurationList = 0BA02E4D1D87571500F1E8D3 /* Build configuration list for PBXProject "{{.ProjectName}}" */;
			compatibilityVersion = "Xcode 8.0";
			developmentRegion = English;
			hasScannedForEncodings = 0;
			knownRegions = (
				en,
			);
			mainGroup = 0BA02E491D87571500F1E8D3;
			projectDirPath = "";
			projectRoot = "";
			targets = (
			);
		};
/* End PBXProject section */

/* Begin XCBuildConfiguration section */
		0BA02E4E1D87571500F1E8D3 /* Debug */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
			};
			name = Debug;
		};
		0BA02E4F1D87571500F1E8D3 /* Release */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
			};
			name = Release;
		};
/* End XCBuildConfiguration section */

/* Begin XCConfigurationList section */
		0BA02E4D1D87571500F1E8D3 /* Build configuration list for PBXProject "{{.ProjectName}}" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				0BA02E4E1D87571500F1E8D3 /* Debug */,
				0BA02E4F1D87571500F1E8D3 /* Release */,
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
/* End XCConfigurationList section */
	};
	rootObject = 0BA02E4A1D87571500F1E8D3 /* Project object */;
}
`)
