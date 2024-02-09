package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetWebsiteListHandler(t *testing.T) {
	// Test case: Valid JSON input
	jsonData := []byte(`{"websites":["https://example.com","https://google.com"]}`)
	request, err := http.NewRequest("POST", "/input", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	getWebsiteList(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test case: Empty JSON input
	emptyData := []byte(`{}`)
	request, err = http.NewRequest("POST", "/input", bytes.NewBuffer(emptyData))
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder = httptest.NewRecorder()
	getWebsiteList(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Test case: Invalid JSON input
	invalidData := []byte(`invalid json`)
	request, err = http.NewRequest("POST", "/input", bytes.NewBuffer(invalidData))
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder = httptest.NewRecorder()
	getWebsiteList(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestCheckWebsiteStatusHandler(t *testing.T) {
	// Test case: No websites checked yet
	request, err := http.NewRequest("GET", "/check", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	checkWebsiteStatusHandler(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Test case: Websites checked, expect JSON response
	websiteList := []string{"https://example.com", "https://google.com"}
	checkWebsiteStatus(websiteList)
	request, err = http.NewRequest("GET", "/check", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder = httptest.NewRecorder()
	checkWebsiteStatusHandler(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test case: Specific website status requested
	request, err = http.NewRequest("GET", "/check?name=https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder = httptest.NewRecorder()
	checkWebsiteStatusHandler(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGenericHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	genericHandler(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHTTPChecker_Check(t *testing.T) {
	// Test case: Valid website returns UP status
	checker := httpChecker{}
	ctx := context.Background()
	status, err := checker.Check(ctx, "https://example.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if status != Up {
		t.Errorf("expected status UP, got %v", status)
	}

	// Test case: Invalid website returns DOWN status
	status, err = checker.Check(ctx, "https://invalidurl")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if status != Down {
		t.Errorf("expected status DOWN, got %v", status)
	}
}

func TestCheckWebsiteStatus(t *testing.T) {
	// Test case: Check status for a list of websites
	websiteList := []string{"https://example.com", "https://google.com"}
	checkWebsiteStatus(websiteList)
	time.Sleep(6 * time.Second) // Wait for goroutines to finish
	if status := websiteMap["https://example.com"]; status != Up {
		t.Errorf("expected status Up, got %v", status)
	}
	if status := websiteMap["https://google.com"]; status != Up {
		t.Errorf("expected status Up, got %v", status)
	}
}
