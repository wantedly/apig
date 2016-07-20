package main

type Detail struct {
	VCS       string
	User      string
	Project   string
	Models    []*Model
	Model     *Model
	ImportDir string
}
