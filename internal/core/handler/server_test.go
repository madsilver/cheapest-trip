package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const msg = "got %v want %v given"

func TestServerMethodNotAllowed(t *testing.T) {
	expect := http.StatusMethodNotAllowed

	req, err := http.NewRequest("PUT", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routerHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != expect {
		t.Errorf(msg, status, expect)
	}
}

func TestServerRouteNotGiven(t *testing.T) {
	expect := http.StatusBadRequest
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routerHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != expect {
		t.Errorf(msg, status, expect)
	}
}

func TestServerGetRoute(t *testing.T) {
	expect := http.StatusOK
	req, err := http.NewRequest("GET", "/routes?route=ABC-DEF", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routerHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != expect {
		t.Errorf(msg, status, expect)
	}

	var res map[string]interface{}
	json.NewDecoder(rr.Body).Decode(&res)

	assert.Equal(t, res["departure"], "ABC")
	assert.Equal(t, res["arrival"], "DEF")
	assert.Equal(t, res["cost"], 0.)
	assert.Equal(t, res["best_route"], []interface{}([]interface{}{}))
}

func TestServerPostRouteRequiredField(t *testing.T) {
	expect := http.StatusBadRequest
	json := []byte(`{"departure": "ABC"}`)

	req, err := http.NewRequest("POST", "/routes", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routerHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != expect {
		t.Errorf(msg, status, expect)
	}
}
