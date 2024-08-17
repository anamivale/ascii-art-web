package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		url      string
		formData string
		expected int
	}{
		{
			name:     "GET Request to root",
			method:   http.MethodGet,
			url:      "/",
			expected: http.StatusOK,
		},
		
		{
			name:     "Request to nonexistent URL",
			method:   http.MethodGet,
			url:      "/nonexistent",
			expected: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.formData))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			// Set Content-Type for POST requests
			if tt.method == http.MethodPost {
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expected {
				t.Errorf("handler returned wrong status code for %s %s: got %v want %v\nResponse Body: %s",
					tt.method, tt.url, status, tt.expected, rr.Body.String())
			}
		})
	}
}

func TestErrorHandlingMiddleware(t *testing.T) {
	handler := ErrorHandlingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("Test panic")
	}))

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
