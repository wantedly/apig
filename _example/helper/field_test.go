package helper

import (
	"testing"
)

func TestParseFields_Wildcard(t *testing.T) {
	fields := "*"
	result := ParseFields(fields)

	if _, ok := result["*"]; !ok {
		t.Fatalf("result[*] should exist: %#v", result)
	}

	if result["*"] != nil {
		t.Fatalf("result[*] should be nil: %#v", result)
	}
}

func TestParseFields_Flat(t *testing.T) {
	fields := "profile"
	result := ParseFields(fields)

	if _, ok := result["profile"]; !ok {
		t.Fatalf("result[profile] should exist: %#v", result)
	}

	if result["profile"] != nil {
		t.Fatalf("result[profile] should be nil: %#v", result)
	}
}

func TestParseFields_Nested(t *testing.T) {
	fields := "profile.nation"
	result := ParseFields(fields)

	if _, ok := result["profile"]; !ok {
		t.Fatalf("result[profile] should exist: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{}); !ok {
		t.Fatalf("result[profile] should be map: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["nation"]; !ok {
		t.Fatalf("result[profile][nation] should exist: %#v", result)
	}

	if result["profile"].(map[string]interface{})["nation"] != nil {
		t.Fatalf("result[profile][nation] should be nil: %#v", result)
	}
}

func TestParseFields_NestedDeeply(t *testing.T) {
	fields := "profile.nation.name"
	result := ParseFields(fields)

	if _, ok := result["profile"]; !ok {
		t.Fatalf("result[profile] should exist: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{}); !ok {
		t.Fatalf("result[profile] should be map: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["nation"]; !ok {
		t.Fatalf("result[profile][nation] should exist: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["nation"].(map[string]interface{}); !ok {
		t.Fatalf("result[profile][nation] should be map: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["nation"].(map[string]interface{})["name"]; !ok {
		t.Fatalf("result[profile][nation][name] should exist: %#v", result)
	}

	if result["profile"].(map[string]interface{})["nation"].(map[string]interface{})["name"] != nil {
		t.Fatalf("result[profile][nation][name] should be nil: %#v", result)
	}
}

func TestParseFields_MultipleFields(t *testing.T) {
	fields := "profile.nation.name,emails"
	result := ParseFields(fields)

	if _, ok := result["profile"]; !ok {
		t.Fatalf("result[profile] should exist: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{}); !ok {
		t.Fatalf("result[profile] should be map: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["nation"]; !ok {
		t.Fatalf("result[profile][nation] should exist: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["nation"].(map[string]interface{}); !ok {
		t.Fatalf("result[profile][nation] should be map: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["nation"].(map[string]interface{})["name"]; !ok {
		t.Fatalf("result[profile][nation][name] should exist: %#v", result)
	}

	if result["profile"].(map[string]interface{})["nation"].(map[string]interface{})["name"] != nil {
		t.Fatalf("result[profile][nation][name] should be nil: %#v", result)
	}

	if _, ok := result["emails"]; !ok {
		t.Fatalf("result[emails] should exist: %#v", result)
	}

	if result["emails"] != nil {
		t.Fatalf("result[emails] should be map: %#v", result)
	}
}

func TestParseFields_Included(t *testing.T) {
	fields := "profile.nation.name,profile"
	result := ParseFields(fields)

	if _, ok := result["profile"]; !ok {
		t.Fatalf("result[profile] should exist: %#v", result)
	}

	if result["profile"] != nil {
		t.Fatalf("result[profile] should be nil: %#v", result)
	}
}
