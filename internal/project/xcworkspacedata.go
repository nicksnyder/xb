package project

import "github.com/nicksnyder/xb/internal/fs"

func (p *Project) xcworkspacedata() *fs.File {
	data := map[string]interface{}{
		"Project": p,
	}
	template := newTemplate(pbxprojTemplate, data)
	return &fs.File{
		Name:     "contents.xcworkspacedata",
		Contents: template,
	}
}

var xcworkspacedataTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<Workspace
   version = "1.0">
   <FileRef
      location = "self:{{.Project.Name}}.xcodeproj">
   </FileRef>
</Workspace>
`
