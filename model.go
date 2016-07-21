package main

type Model struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name        string
	Type        string
	Tag         string
	Association *Association
}

type Association struct {
	Type  string
	Model *Model
}
