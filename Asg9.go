package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Status string

const (
	Up   Status = "UP"
	Down Status = "DOWN"
)

type Websites struct {
	List []string `json:"websites"`
}

var websiteMap map[string]Status = make(map[string]Status)
var list []string

// interface --> for checking website status
type StatusChecker interface {
	Check(ctx context.Context, name string) (status Status, err error)
}

// httpChecker implements StatusChecker --> using HTTP calls
type httpChecker struct{}

func (h httpChecker) Check(ctx context.Context, name string) (Status, error) { // context for cancellation and timeouts
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // timeout
	defer cancel()

	response, err := http.Get("https://" + name)
	if err != nil {
		return Down, err
	}
	defer response.Body.Close()
	return Up, nil
}

var checker StatusChecker = httpChecker{}

func getWebsiteList(w http.ResponseWriter, r *http.Request) {
	var website Websites
	err := json.NewDecoder(r.Body).Decode(&website)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list = website.List
	w.WriteHeader(http.StatusOK)
}

func checkWebsiteStatus(web []string) {
	fmt.Println("running checkWebsiteStatus")
	for _, data := range web {
		go func(data string) {
			status, err := checker.Check(context.Background(), data)
			if err != nil {
				fmt.Printf("Error checking %s: %v\n", data, err)
				websiteMap[data] = Down
			} else {
				websiteMap[data] = status
			}
		}(data)
	}
}

func checkWebsiteStatusHandler(w http.ResponseWriter, r *http.Request) {
	if len(websiteMap) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("The map is empty"))
		return
	}

	name := r.URL.Query().Get("name")

	if name != "" {
		status, ok := websiteMap[name]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("name not in the list"))
			return
		}

		respJson, err := json.Marshal(map[string]string{name: string(status)})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
		return
	}

	respJson, err := json.Marshal(websiteMap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}

func genericHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the server, the page you are looking for does not exist!")
}

func main() {
	fmt.Println("Starting server")

	r := mux.NewRouter()

	r.HandleFunc("/", genericHandler)

	r.HandleFunc("/input", getWebsiteList).Methods("POST")

	r.HandleFunc("/check", checkWebsiteStatusHandler).Methods("GET")

	go func() {
		for {
			if list != nil {
				//fmt.Print(list)
				checkWebsiteStatus(list)
				time.Sleep(time.Minute)
			}
		}
	}()

	http.ListenAndServe(":8080", r)
}

