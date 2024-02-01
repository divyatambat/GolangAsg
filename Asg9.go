package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// representing status of a website
type WebsiteStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// interface for checking website status
type StatusChecker interface {
	Check(name string) (status bool, err error)
}

// implements StatusChecker
type httpChecker struct{}

func (h httpChecker) Check(name string) (status bool, err error) {
	resp, err := http.Get(name)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp != nil {
		return resp.StatusCode == 200, nil
	} else {
		return false, errors.New("Invalid response!")
	}
}

// websiteMap --> storing website names and their statuses
var websiteMap = sync.Map{}

func monitorWebsites(checker StatusChecker) {
	for {
		var wg sync.WaitGroup
		var errs []error

		websiteMap.Range(func(key, value interface{}) bool {
			name, ok := key.(string) // assert key type
			if !ok {
				fmt.Println("Unexpected key type in websiteMap:", key)
				return true
			}
			wg.Add(1)
			go func(name string) {
				defer wg.Done()
				_, err := checker.Check(name)
				if err != nil {
					errs = append(errs, fmt.Errorf("error checking website %s: %v", name, err))
					return
				}
				websiteMap.Store(name, "UP")
			}(name)
			return true
		})

		wg.Wait()

		if len(errs) > 0 {
			fmt.Println("Errors encountered:", errs)
		}

		time.Sleep(5 * time.Second)
	}
}

func submitWebsitesHandler(w http.ResponseWriter, r *http.Request) {
	var websites []string
	if err := json.NewDecoder(r.Body).Decode(&websites); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, website := range websites {
		websiteMap.Store(website, "UNKNOWN")
	}

	// monitoring in goroutine
	go monitorWebsites(httpChecker{})
	w.WriteHeader(http.StatusOK)
}

func getWebsitesHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Printf("Received request for website: %s\n", name)

	var statuses []WebsiteStatus
	if name != "" {
		// Get status for specific website
		status, ok := websiteMap.Load(name)
		if !ok {
			fmt.Printf("Website not found in websiteMap: %s\n", name)
			http.Error(w, "Website not found", http.StatusNotFound)
			return
		}
		statuses = append(statuses, WebsiteStatus{Name: name, Status: status.(string)})
	} else {
		// Get status for all websites
		websiteMap.Range(func(key, value interface{}) bool {
			if status, ok := value.(string); ok {
				statuses = append(statuses, WebsiteStatus{Name: key.(string), Status: status})
			} else {
				fmt.Printf("Unexpected value type in websiteMap: %v\n", value) // Log for debugging
			}
			return true
		})
	}

	// checking websiteMap is not empty before encoding
	if len(statuses) == 0 {
		fmt.Println("No website statuses found in websiteMap")
		http.Error(w, "No website statuses available", http.StatusNotFound)
		return
	}
	// encode and send response
	json.NewEncoder(w).Encode(statuses)
}

func main() {
	http.HandleFunc("/websites/submit", submitWebsitesHandler)
	http.HandleFunc("/websites/status", getWebsitesHandler)

	// start website monitoring in a separate goroutine
	go monitorWebsites(httpChecker{})

	fmt.Println("Starting server on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
}

