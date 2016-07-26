package apig

type Model struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name        string
	JSONName    string
	Type        string
	Tag         string
	Association *Association
}

type Association struct {
	Type  int
	Model *Model
}
