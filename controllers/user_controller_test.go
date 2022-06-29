package controllers_test

import (
	"TOGO/controllers"
	"TOGO/middleware"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetMe(t *testing.T) {
	req, _ := http.NewRequest("GET", "/me", nil)
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY1MjczMjcsImlkIjoiNjJiYWJkNTA0OTBjMmE0ODc4MTViYzcxIn0.m-Je6dZC73vm5GA8YyUEFDylakvUbA24G1L9K1BYZRI"
	req.Header.Set("Authorization", token)

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(controllers.GetMe()))
	handler.ServeHTTP(rr, req)
	var r controllers.Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	//check satatus code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//check user
	if r.Data["username"] != "tuanchoi1" {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Data["username"], "tuanchoi1")
	}
}

func TestGetUser(t *testing.T) {
	router := mux.NewRouter()
	req, err := http.NewRequest("GET", fmt.Sprintf("/user/%s", "62babd50490c2a487815bc71"), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(GetUser())
	router.ServeHTTP(rr, req)
	var r controllers.Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if status := r.Status; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if r.Data["username"] != "tuanchoi1" {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Data["username"], "tuanchoi1")
	}

}

func TestUpdateMe(t *testing.T) {
	var jsonStr = []byte(`{"name": "Test tuandz"}`)
	req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(jsonStr))
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY1MjczMjcsImlkIjoiNjJiYWJkNTA0OTBjMmE0ODc4MTViYzcxIn0.m-Je6dZC73vm5GA8YyUEFDylakvUbA24G1L9K1BYZRI"
	req.Header.Set("Authorization", token)

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(controllers.GetMe()))
	handler.ServeHTTP(rr, req)
	var r controllers.Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Status, http.StatusOK)
	}

	if r.Data["name"] != "Test tuandz" {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Data["name"], "Test tuandz")
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/user/%s", "62b978f9de0fcea8d82cce85"), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(GetUser())
	handler := http.HandlerFunc(controllers.GetUser())
	handler.ServeHTTP(rr, req)
	var r controllers.Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Status, http.StatusOK)
	}
	if r.Message != "success" {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Message, "success")
	}
}
