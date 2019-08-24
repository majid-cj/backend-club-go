package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Request ...
type Request struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

var requests []Request

func test(w http.ResponseWriter, r *http.Request) {
	var req Request
	if r.Method == "GET" {
		fmt.Fprint(w, requests)
	} else if r.Method == "POST" {
		json.NewDecoder(r.Body).Decode(&req)
		requests = append(requests, req)
		fmt.Fprint(w, req)
	} else {
		fmt.Fprint(w, "use valid method")
	}
}

func main() {
	http.HandleFunc("/test", test)

	http.ListenAndServe(":8080", nil)
}
