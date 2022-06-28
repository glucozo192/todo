package test

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/go-playground/assert/v2"
// )

// func TestSignup(t *testing.T) {
// 	var jsonStr = []byte(`{"username":"test5","password":"123456","name":"nguyen tuan"}`)
// 	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(Signup())
// 	handler.ServeHTTP(rr, req)

// 	assert.Equal(t, rr.Code, http.StatusOK)
// }
