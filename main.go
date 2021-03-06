package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//Info - struct for api version data
type Info struct {
	Version        string `json:"version"`
	AppID          string `json:"appID"`
	InstanceID     string `json:"instanceID"`
	ServerSoftware string `json:"serverSoftware"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/info", infoHandler)
	appengine.Main()
}

//indexHandler - displays the index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := appengine.NewContext(r)
	// log.Debugf(ctx, "%#v", os.Environ())
	fmt.Fprintf(w, "hello from the index")
}

//infoHandler - prints api info in json
func infoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	info := Info{Version: "dev", AppID: "dev", InstanceID: "dev", ServerSoftware: "dev"}
	log.Debugf(ctx, "Is dev %v", appengine.IsDevAppServer())
	if !appengine.IsDevAppServer() {
		info.Version = appengine.VersionID(ctx)
		info.AppID = appengine.AppID(ctx)
		info.InstanceID = appengine.InstanceID()
		info.ServerSoftware = appengine.ServerSoftware()
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
