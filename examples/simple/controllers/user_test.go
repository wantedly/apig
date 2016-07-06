package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/wantedly/api-server-generator/examples/simple/db"
	"github.com/wantedly/api-server-generator/examples/simple/models"
	"github.com/wantedly/api-server-generator/examples/simple/server"

	"testing"
)

var uuid string

func TestGetUsers(t *testing.T) {
	response := httptest.NewRecorder()
	database := db.Connect()
	s := server.Setup(database)
	req, err := http.NewRequest("GET", "http://localhost:8080/api/users", nil)
	if err != nil {
		t.Error(err)
	}
	s.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("Got error for GET request to /api/users")
	}
}

func TestCreateUser(t *testing.T) {
	response := httptest.NewRecorder()
	database := db.Connect()
	s := server.Setup(database)
	requestParams := `{
		"name": "NAME"
	}`
	req, err := http.NewRequest("POST", "http://localhost:8080/api/users", bytes.NewBuffer([]byte(requestParams)))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	s.ServeHTTP(response, req)
	if response.Code != http.StatusCreated {
		t.Errorf("Got error for POST request to /api/users")
	}

	body := json.NewDecoder(response.Body)
	var user models.User
	body.Decode(&user)

	if user.Name != "NAME" {
		t.Errorf("Create user failed.\nGot: %v", user)
	}

	uuid = strconv.Itoa(int(user.ID))
}

func TestGetUser(t *testing.T) {
	response := httptest.NewRecorder()
	database := db.Connect()
	s := server.Setup(database)
	req, err := http.NewRequest("GET", "http://localhost:8080/api/users/"+uuid, nil)
	if err != nil {
		t.Error(err)
	}
	s.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("Got error for GET request to /api/users/" + uuid)
	}
	body := json.NewDecoder(response.Body)
	var user models.User
	body.Decode(&user)

	if user.Name != "NAME" {
		t.Errorf("Show user failed.\nGot: %v", user)
	}
}

func TestUpdateUser(t *testing.T) {
	response := httptest.NewRecorder()
	database := db.Connect()
	s := server.Setup(database)
	requestParams := `{
		"name": "NAME_2"
	}`
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/users/"+uuid, bytes.NewBuffer([]byte(requestParams)))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	s.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("Got error for PUT request to /api/users/" + uuid)
	}
	body := json.NewDecoder(response.Body)
	var user models.User
	body.Decode(&user)

	if user.Name != "NAME_2" {
		t.Errorf("Update user failed.\nGot: %v", user)
	}
}

func TestDeleteUser(t *testing.T) {
	response := httptest.NewRecorder()
	database := db.Connect()
	s := server.Setup(database)
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/users/"+uuid, nil)
	if err != nil {
		t.Error(err)
	}
	s.ServeHTTP(response, req)
	if response.Code != http.StatusNoContent {
		t.Errorf("Got error for DELETE request to /api/users/" + uuid)
	}
}
