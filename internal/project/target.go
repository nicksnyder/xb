package project

type Target struct {
	Name                 string   `json:"name"`
	Type                 string   `json:"type"`
	Platform             string   `json:"platform"`
	Sources              []string `json:"sources"`
	Dependencies         []string `json:"dependencies"`
	TestableDependencies []string `json:"testableDependencies"`
}

func (t *Target) XcodeID() string {
	return xcodeID(t.Name)
}
