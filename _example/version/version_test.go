package version

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAcceptHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json;version= 1.0.0 ; more information; more information")
	c := &gin.Context{
		Request: req,
	}
	ver, _ := New(c)
	if ver != "1.0.0" {
		t.Errorf("Accept header should be `1.0.0`. actual: %#v", ver)
	}
}

func TestEmptyAcceptHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json; more information; more information")
	c := &gin.Context{
		Request: req,
	}
	ver, _ := New(c)
	if ver != "-1" {
		t.Errorf("Accept header should be the latest version `-1`. actual: %#v", ver)
	}
}

func TestUndefinedAcceptHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	c := &gin.Context{
		Request: req,
	}
	ver, _ := New(c)
	if ver != "-1" {
		t.Errorf("No accept header should be the latest version `-1`. actual: %#v", ver)
	}
}

func TestQuery(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?v=1.0.1", nil)
	req.Header.Add("Accept", "application/json;version= 1.0.0 ; more information; more information")
	c := &gin.Context{
		Request: req,
	}
	ver, _ := New(c)
	if ver != "1.0.1" {
		t.Errorf("URL Query should be `1.0.1`. actual: %#v", ver)
	}
}

func TestRange(t *testing.T) {

	if Range("1.2.3", "<", "0.9") {
		t.Errorf("defect in <")
	}

	if Range("1.2.3", "<", "0.9.1") {
		t.Errorf("defect in <")
	}

	if Range("1.2.3", "<", "1.2.2") {
		t.Errorf("defect in <")
	}

	if Range("1.2.3", "<", "1.2.3") {
		t.Errorf("defect in <")
	}

	if !Range("1.2.3", "<", "1.2.4") {
		t.Errorf("defect in <")
	}

	if !Range("1.2.3", "<", "1.2") {
		t.Errorf("defect in <")
	}

	if !Range("1.2.3", "<", "1.5") {
		t.Errorf("defect in <")
	}

	if !Range("1.2.3", "<", "-1") {
		t.Errorf("defect in <")
	}

	if Range("1.2.3", "<=", "0.9") {
		t.Errorf("defect in <=")
	}

	if Range("1.2.3", "<=", "0.9.1") {
		t.Errorf("defect in <=")
	}

	if Range("1.2.3", "<=", "1.2.2") {
		t.Errorf("defect in <=")
	}

	if !Range("1.2.3", "<=", "1.2.3") {
		t.Errorf("defect in <=")
	}

	if !Range("1.2.3", "<=", "1.2.4") {
		t.Errorf("defect in <=")
	}

	if !Range("1.2.3", "<=", "1.2") {
		t.Errorf("defect in <=")
	}

	if !Range("1.2.3", "<=", "1.5") {
		t.Errorf("defect in <=")
	}

	if !Range("1.2.3", "<=", "-1") {
		t.Errorf("defect in <=")
	}

	if !Range("1.2.3", ">", "0.9") {
		t.Errorf("defect in >")
	}

	if !Range("1.2.3", ">", "0.9.1") {
		t.Errorf("defect in >")
	}

	if !Range("1.2.3", ">", "1.2.2") {
		t.Errorf("defect in >")
	}

	if Range("1.2.3", ">", "1.2.3") {
		t.Errorf("defect in >")
	}

	if Range("1.2.3", ">", "1.2.4") {
		t.Errorf("defect in >")
	}

	if Range("1.2.3", ">", "1.2") {
		t.Errorf("defect in >")
	}

	if Range("1.2.3", ">", "1.5") {
		t.Errorf("defect in >")
	}

	if Range("1.2.3", ">", "-1") {
		t.Errorf("defect in >")
	}

	if !Range("1.2.3", ">=", "0.9") {
		t.Errorf("defect in >=")
	}

	if !Range("1.2.3", ">=", "0.9.1") {
		t.Errorf("defect in >=")
	}

	if !Range("1.2.3", ">=", "1.2.2") {
		t.Errorf("defect in >=")
	}

	if !Range("1.2.3", ">=", "1.2.3") {
		t.Errorf("defect in >=")
	}

	if Range("1.2.3", ">=", "1.2.4") {
		t.Errorf("defect in >=")
	}

	if Range("1.2.3", ">=", "1.2") {
		t.Errorf("defect in >=")
	}

	if Range("1.2.3", ">=", "1.5") {
		t.Errorf("defect in >=")
	}

	if Range("1.2.3", ">=", "-1") {
		t.Errorf("defect in >=")
	}

	if Range("1.2.3", "==", "0.9") {
		t.Errorf("defect in ==")
	}

	if Range("1.2.3", "==", "0.9.1") {
		t.Errorf("defect in ==")
	}

	if Range("1.2.3", "==", "1.2.2") {
		t.Errorf("defect in ==")
	}

	if !Range("1.2.3", "==", "1.2.3") {
		t.Errorf("defect in ==")
	}

	if Range("1.2.3", "==", "1.2.4") {
		t.Errorf("defect in ==")
	}

	if Range("1.2.3", "==", "1.2") {
		t.Errorf("defect in ==")
	}

	if Range("1.2.3", "==", "1.5") {
		t.Errorf("defect in ==")
	}

	if Range("1.2.3", "==", "-1") {
		t.Errorf("defect in ==")
	}
}
