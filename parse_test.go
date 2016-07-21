package main

import (
	"path/filepath"
	"testing"
)

func TestParseModel(t *testing.T) {
	path := filepath.Join("_fixtures", "models.go")

	models, err := parseModel(path)

	if err != nil {
		t.Fatalf("Failed to parse model file. error: %s", err)
	}

	if len(models) != 2 {
		t.Fatalf("Number of parsed models is incorrect. expected: 2, actual: %d", len(models))
	}

	user := models[0]

	if user.Name != "User" {
		t.Fatalf("Incorrect model name. expected: User, actual: %s", user.Name)
	}

	expectedFields := map[string]string{
		"ID":        "uint",
		"Name":      "string",
		"CreatedAt": "*time.Time",
		"UpdatedAt": "*time.Time",
	}

	fmap := convertMap(models[0].Fields)

	for k, v := range fmap {
		if _, ok := expectedFields[k]; !ok {
			t.Fatalf("Invalid field name: %s", k)
		}

		if v != expectedFields[k] {
			t.Fatalf("Invalid field type. expected: %s, actual: %s", expectedFields[k], v)
		}
	}
}

func convertMap(fields []*Field) map[string]string {
	fmap := make(map[string]string)

	for _, field := range fields {
		fmap[field.Name] = field.Type
	}

	return fmap
}
