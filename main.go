package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

//Info - struct for api version data
type Info struct {
	Version string `json:"version"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/info", infoHandler)
	appengine.Main()
}

//indexHandler - displays the index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from the index")
}

//infoHandler - prints api info in json
func infoHandler(w http.ResponseWriter, r *http.Request) {
	var info Info
	info.Version = "0.1.0"
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}