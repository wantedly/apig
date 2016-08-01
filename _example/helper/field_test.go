package helper

import (
	"testing"
)

type User struct {
	ID      uint     `json:"id"`
	Jobs    []*Job   `json:"jobs"`
	Name    string   `json:"name"`
	Profile *Profile `json:"profile"`
}

type Profile struct {
	ID      uint  `json:"id"`
	UserID  uint  `json:"user_id"`
	User    *User `json:"user"`
	Engaged bool  `json:"engaged"`
}

type Job struct {
	ID     uint  `json:"id"`
	UserID uint  `json:"user_id"`
	User   *User `json:"user"`
	RoleCd uint  `json:"role_cd"`
}

func TestQueryFields_Wildcard(t *testing.T) {
	fields := map[string]interface{}{"*": nil}
	result := QueryFields(User{}, fields)
	expected := "*"

	if result != expected {
		t.Fatalf("result should be %s. actual: %s", expected, result)
	}
}

func TestQueryFields_Primitive(t *testing.T) {
	fields := map[string]interface{}{"name": nil}
	result := QueryFields(User{}, fields)
	expected := "name"

	if result != expected {
		t.Fatalf("result should be %s. actual: %s", expected, result)
	}
}

func TestQueryFields_Multiple(t *testing.T) {
	fields := map[string]interface{}{"id": nil, "name": nil}
	result := QueryFields(User{}, fields)
	expected := "id,name"

	if result != expected {
		t.Fatalf("result should be %s. actual: %s", expected, result)
	}
}

func TestQueryFields_BelongsTo(t *testing.T) {
	fields := map[string]interface{}{"user": nil}
	result := QueryFields(Profile{}, fields)
	expected := "user_id"

	if result != expected {
		t.Fatalf("result should be %s. actual: %s", expected, result)
	}
}

func TestQueryFields_HasOne(t *testing.T) {
	fields := map[string]interface{}{"profile": nil}
	result := QueryFields(User{}, fields)
	expected := "id"

	if result != expected {
		t.Fatalf("result should be %s. actual: %s", expected, result)
	}
}

func TestQueryFields_HasMany(t *testing.T) {
	fields := map[string]interface{}{"jobs": nil}
	result := QueryFields(User{}, fields)
	expected := "id"

	if result != expected {
		t.Fatalf("result should be %s. actual: %s", expected, result)
	}
}

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
