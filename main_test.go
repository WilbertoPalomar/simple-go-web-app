package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionHandler(t *testing.T) {

	// A request will be created passing to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder creation section
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(versionHandler)

	// Calling ServeHTTP method directly to pass our request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the expected status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the expected response body.
	expected := version
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
