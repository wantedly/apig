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

func (f *Field) IsAssociation() bool {
	return f.Association != nil && f.Association.Type != AssociationNone
}

func (f *Field) IsBelongsTo() bool {
	return f.Association != nil && f.Association.Type == AssociationBelongsTo
}

type Association struct {
	Type  int
	Model *Model
}
