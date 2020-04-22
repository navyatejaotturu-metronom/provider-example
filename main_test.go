package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type payload map[string]interface{}

func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(User)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := payload{
		"id":   1,
		"name": "john",
	}
	if rr.Body.String() != expected.toJSON() {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func (p payload) toJSON() string {
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(j)
}
