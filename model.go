package main

const (
	AssociationNone      = 0
	AssociationBelongsTo = 1
	AssociationHasMany   = 2
	AssociationHasOne    = 3
)

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
