package main

type Model struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name        string
	Type        string
	Association string
}
