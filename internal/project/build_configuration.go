package project

type BuildConfiguration struct {
	Name string
}

func (bc *BuildConfiguration) XcodeID() string {
	return xcodeID(bc.Name)
}
