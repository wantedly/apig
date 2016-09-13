package helper

import (
	"testing"
)

type User struct {
	ID      uint     `json:"id" form:"id"`
	Jobs    []*Job   `json:"jobs,omitempty" form:"jobs"`
	Name    string   `json:"name" form:"name"`
	Profile *Profile `json:"profile,omitempty" form:"profile"`
}

type Profile struct {
	ID      uint  `json:"id" form:"id"`
	UserID  uint  `json:"user_id" form:"user_id"`
	User    *User `json:"user" form:"user"`
	Engaged bool  `json:"engaged" form:"engaged"`
}

type Job struct {
	ID     uint  `json:"id" form:"id"`
	UserID uint  `json:"user_id" form:"user_id"`
	User   *User `json:"user" form:"user"`
	RoleCd uint  `json:"role_cd" form:"role_cd"`
}

type Company struct {
	ID           uint              `json:"id,omitempty" form:"id"`
	Name         string            `json:"name,omitempty" form:"name"`
	List         bool              `json:"list,omitempty" form:"list"`
	Subsidiary   []*Company        `json:"company,omitempty" form:"company"`
	Organization map[string]string `json:"organization,omitempty" form:"organization"`
	User         *User             `json:"user,omitempty" form:"user"`
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

var profile = Profile{
	ID:      1,
	UserID:  1,
	User:    nil,
	Engaged: true,
}

var job = Job{
	ID:     1,
	UserID: 1,
	User:   nil,
	RoleCd: 1,
}

func TestFieldToMap_Wildcard(t *testing.T) {
	user := User{
		ID:      1,
		Jobs:    []*Job{&job},
		Name:    "Taro Yamada",
		Profile: &profile,
	}

	fields := map[string]interface{}{
		"*": nil,
	}
	result, err := FieldToMap(user, fields)

	if err != nil {
		t.Fatalf("FieldToMap return an error. detail: %#v", err.Error())
	}

	for _, key := range []string{"id", "jobs", "name", "profile"} {
		if _, ok := result[key]; !ok {
			t.Fatalf("%s should exist. actual: %#v", key, result)
		}
	}

	if result["jobs"].([]*Job) == nil {
		t.Fatalf("jobs should not be nil. actual: %#v", result["jobs"])
	}

	if result["profile"].(*Profile) == nil {
		t.Fatalf("profile should not be nil. actual: %#v", result["profile"])
	}
}

func TestFieldToMap_OmitEmpty(t *testing.T) {
	user := User{
		ID:      1,
		Jobs:    nil,
		Name:    "Taro Yamada",
		Profile: nil,
	}

	fields := map[string]interface{}{
		"*": nil,
	}
	result, err := FieldToMap(user, fields)

	if err != nil {
		t.Fatalf("FieldToMap return an error. detail: %#v", err.Error())
	}

	for _, key := range []string{"id", "name"} {
		if _, ok := result[key]; !ok {
			t.Fatalf("%s should exist. actual: %#v", key, result)
		}
	}

	for _, key := range []string{"jobs", "profile"} {
		if _, ok := result[key]; ok {
			t.Fatalf("%s should not exist. actual: %#v", key, result)
		}
	}
}

func TestFieldToMap_OmitEmptyWithField(t *testing.T) {
	user := User{
		ID:      1,
		Jobs:    nil,
		Name:    "Taro Yamada",
		Profile: nil,
	}

	fields := map[string]interface{}{
		"id":   nil,
		"name": nil,
		"jobs": nil,
	}
	result, err := FieldToMap(user, fields)

	if err != nil {
		t.Fatalf("FieldToMap return an error. detail: %#v", err.Error())
	}

	for _, key := range []string{"id", "name", "jobs"} {
		if _, ok := result[key]; !ok {
			t.Fatalf("%s should exist. actual: %#v", key, result)
		}
	}

	for _, key := range []string{"profile"} {
		if _, ok := result[key]; ok {
			t.Fatalf("%s should not exist. actual: %#v", key, result)
		}
	}
}

func TestFieldToMap_OmitEmptyAllTypes(t *testing.T) {
	company := Company{
		ID:           0,
		Name:         "",
		List:         false,
		Subsidiary:   []*Company{},
		Organization: make(map[string]string),
		User:         nil,
	}

	fields := map[string]interface{}{
		"*": nil,
	}
	result, err := FieldToMap(company, fields)

	if err != nil {
		t.Fatalf("FieldToMap return an error. detail: %#v", err.Error())
	}

	for _, key := range []string{"id", "name", "list", "subsidiary", "organization", "user"} {
		if _, ok := result[key]; ok {
			t.Fatalf("%s should not exist. actual: %#v", key, result)
		}
	}
}

func TestFieldToMap_SpecifyField(t *testing.T) {
	user := User{
		ID:      1,
		Jobs:    nil,
		Name:    "Taro Yamada",
		Profile: nil,
	}

	fields := map[string]interface{}{
		"id":   nil,
		"name": nil,
	}
	result, err := FieldToMap(user, fields)

	if err != nil {
		t.Fatalf("FieldToMap return an error. detail: %#v", err.Error())
	}

	for _, key := range []string{"id", "name"} {
		if _, ok := result[key]; !ok {
			t.Fatalf("%s should exist. actual: %#v", key, result)
		}
	}

	for _, key := range []string{"jobs", "profile"} {
		if _, ok := result[key]; ok {
			t.Fatalf("%s should not exist. actual: %#v", key, result)
		}
	}
}

func TestFieldToMap_NestedField(t *testing.T) {
	user := User{
		ID:      1,
		Jobs:    []*Job{&job},
		Name:    "Taro Yamada",
		Profile: &profile,
	}

	fields := map[string]interface{}{
		"profile": map[string]interface{}{
			"id": nil,
		},
		"name": nil,
	}
	result, err := FieldToMap(user, fields)

	if err != nil {
		t.Fatalf("FieldToMap return an error. detail: %#v", err.Error())
	}

	for _, key := range []string{"name", "profile"} {
		if _, ok := result[key]; !ok {
			t.Fatalf("%s should exist. actual: %#v", key, result)
		}
	}

	for _, key := range []string{"id", "jobs"} {
		if _, ok := result[key]; ok {
			t.Fatalf("%s should not exist. actual: %#v", key, result)
		}
	}

	if result["profile"].(map[string]interface{}) == nil {
		t.Fatalf("profile should not be nil. actual: %#v", result)
	}

	if _, ok := result["profile"].(map[string]interface{})["id"]; !ok {
		t.Fatalf("profile.id should exist. actual: %#v", result)
	}

	for _, key := range []string{"id"} {
		if _, ok := result["profile"].(map[string]interface{})[key]; !ok {
			t.Fatalf("profile.%s should exist. actual: %#v", key, result)
		}
	}

	for _, key := range []string{"user_id", "user", "engaged"} {
		if _, ok := result["profile"].(map[string]interface{})[key]; ok {
			t.Fatalf("profile.%s should not exist. actual: %#v", key, result)
		}
	}
}
