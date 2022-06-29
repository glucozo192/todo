package controllers

import (
	"TOGO/middleware"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	// send data to create
	var jsonStr = []byte(`{"name": "test task","content": "test content"}`)
	req, _ := http.NewRequest("POST", "/task", bytes.NewBuffer(jsonStr))

	// send token
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY1MjczMjcsImlkIjoiNjJiYWJkNTA0OTBjMmE0ODc4MTViYzcxIn0.m-Je6dZC73vm5GA8YyUEFDylakvUbA24G1L9K1BYZRI"
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(CreateTask()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusCreated {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}

	if r.Data["name"] != "test task" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["name"], "test task")
	}
	if r.Data["content"] != "test content" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["content"], "test content")
	}
}

func TestGetTask(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user-tasks", nil)

	// send token
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTY1MjczMjcsImlkIjoiNjJiYWJkNTA0OTBjMmE0ODc4MTViYzcxIn0.m-Je6dZC73vm5GA8YyUEFDylakvUbA24G1L9K1BYZRI"
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(GetTask()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}
}

func TestGetOneTask(t *testing.T) {

}
