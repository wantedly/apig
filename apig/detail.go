package apig

type Detail struct {
	VCS       string
	User      string
	Project   string
	Namespace string
	Models    []*Model
	Model     *Model
	ImportDir string
}
