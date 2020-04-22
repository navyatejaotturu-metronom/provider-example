package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var port string

	var url string
	if os.Getenv("DRP_CF_HTTP_ADDR") != "" {
		url = os.Getenv("DRP_CF_HTTP_ADDR")
	} else {
		url = "localhost"
	}
	if os.Getenv("DRP_CF_HTTP_PORT") != "" {
		port = os.Getenv("DRP_CF_HTTP_PORT")
	} else {
		port = "8085"
	}
	r := mux.NewRouter()
	r.HandleFunc("/user", User).Methods("Get")
	r.HandleFunc("/.well-known/live", Live).Methods("Get")
	r.HandleFunc("/.well-known/ready", Ready).Methods("Get")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", url, port), r))

}

func User(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["id"] = 1
	response["name"] = "john"
	val, err := json.Marshal(&response)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(val)
}

func Live(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func Ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}
