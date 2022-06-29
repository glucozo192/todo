package controllers

import (
	"TOGO/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Token   string                 `json:"token"`
	Data    map[string]interface{} `json:"data"`
}

func TestSignup(t *testing.T) {
	var jsonStr = []byte(`{"username": "tuanchoitest17", "password": "123456","name":"Nguyen tuan"}`)

	req, err := http.NewRequest("POST", "/user/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Signup())
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if r.Data["username"] != "tuanchoitest17" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["username"], "tuanchoitest17")
	}
	if r.Data["name"] != "Nguyen tuan" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["name"], "Nguyen tuan")
	}
}

func TestLogin(t *testing.T) {
	var jsonStr = []byte(`{"username": "tuanchoi1", "password": "123456"}`)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login())
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if r.Data["username"] != "tuanchoi1" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["username"], "tuanchoi1")
	}

	if !models.CheckPasswordHash("123456", r.Data["password"].(string)) {
		t.Errorf("handler returned wrong data")
	}
}
