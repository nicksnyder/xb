package project

import "github.com/nicksnyder/xb/internal/fs"

func (p *Project) pbxproj() *fs.File {
	data := map[string]interface{}{
		"Project":                  p,
		"ProductsID":               "0B9773641D8AE37900017700",
		"MainGroupID":              "0BA02E491D87571500F1E8D3",
		"ProjectObjectID":          "0BA02E4A1D87571500F1E8D3",
		"BuildConfigurationListID": "0BA02E4D1D87571500F1E8D3",
		"BuildConfigurations": []BuildConfiguration{
			{Name: "Debug"},
			{Name: "Release"},
		},
	}
	template := newTemplate(pbxprojTemplate, data)
	return &fs.File{
		Name:     "project.pbxproj",
		Contents: template,
	}
}

var pbxprojTemplate = `// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 48;
	objects = {

/* Begin PBXGroup section */
		{{.ProductsID}} /* Products */ = {
			isa = PBXGroup;
			children = (
			);
			name = Products;
			sourceTree = "<group>";
		};
		{{.MainGroupID}} = {
			isa = PBXGroup;
			children = (
				{{.ProductsID}} /* Products */,
			);
			sourceTree = "<group>";
		};
/* End PBXGroup section */

/* Begin PBXProject section */
		{{.ProjectObjectID}} /* Project object */ = {
			isa = PBXProject;
			attributes = {
				LastUpgradeCheck = 0800;
			};
			buildConfigurationList = {{.BuildConfigurationListID}} /* Build configuration list for PBXProject "{{.Project.Name}}" */;
			compatibilityVersion = "Xcode 8.0";
			developmentRegion = English;
			hasScannedForEncodings = 0;
			knownRegions = (
				en,
			);
			mainGroup = {{.MainGroupID}};
			productRefGroup = {{.ProductsID}} /* Products */;
			projectDirPath = "";
			projectRoot = "";
			targets = (
			);
		};
/* End PBXProject section */

/* Begin XCBuildConfiguration section */
		{{- range .BuildConfigurations}}
		{{.XcodeID}} /* {{.Name}} */ = {
			isa = XCBuildConfiguration;
			buildSettings = {
			};
			name = {{.Name}};
		};
		{{- end}}
/* End XCBuildConfiguration section */

/* Begin XCConfigurationList section */
		{{.BuildConfigurationListID}} /* Build configuration list for PBXProject "{{.Project.Name}}" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				{{- range .BuildConfigurations}}
				{{.XcodeID}} /* {{.Name}} */
				{{- end}}
			);
			defaultConfigurationIsVisible = 0;
			defaultConfigurationName = Release;
		};
/* End XCConfigurationList section */
	};
	rootObject = {{.ProjectObjectID}} /* Project object */;
}
`
