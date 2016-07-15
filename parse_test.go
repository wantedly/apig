package main

import (
	"path/filepath"
	"testing"
)

func TestParseFile(t *testing.T) {
	path := filepath.Join("_fixtures", "models.go")

	models, err := parseFile(path)

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
		"CreatedAt": "time.Time",
		"UpdatedAt": "time.Time",
	}

	for k, v := range user.Fields {
		if _, ok := expectedFields[k]; !ok {
			t.Fatalf("Invalid field name: %s", k)
		}

		if v != expectedFields[k] {
			t.Fatalf("Invalid field type. expected: %s, actual: %s", expectedFields[k], v)
		}
	}
}
