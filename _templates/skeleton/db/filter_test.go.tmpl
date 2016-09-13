package db

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID      uint   `json:"id,omitempty" form:"id"`
	Name    string `json:"name,omitempty" form:"name"`
	Engaged bool   `json:"engaged,omitempty" form:"engaged"`
}

func contains(ss map[string]string, s string) bool {
	_, ok := ss[s]

	return ok
}

func TestFilterToMap(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?q[id]=1,5,100&q[name]=hoge,fuga&q[unexisted_field]=null", nil)
	c := &gin.Context{
		Request: req,
	}
	value := filterToMap(c, User{})

	if !contains(value, "id") {
		t.Fatalf("Filter should have `id` key.")
	}

	if !contains(value, "name") {
		t.Fatalf("Filter should have `name` key.")
	}

	if contains(value, "unexisted_field") {
		t.Fatalf("Filter should not have `unexisted_field` key.")
	}

	if value["id"] != "1,5,100" {
		t.Fatalf("filters[\"id\"] expected: `1,5,100`, actual: %s", value["id"])
	}

	if value["name"] != "hoge,fuga" {
		t.Fatalf("filters[\"name\"] expected: `hoge,fuga`, actual: %s", value["id"])
	}
}
