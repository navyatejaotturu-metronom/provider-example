package main

import (
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
		url = "0.0.0.0"
	}

	if os.Getenv("DRP_CF_HTTP_PORT") != "" {
		port = os.Getenv("DRP_CF_HTTP_PORT")
	} else {
		port = "80"
	}

	r := mux.NewRouter()
	r.HandleFunc("/user", User).Methods("Get")
	r.HandleFunc("/.well-known/live", Live).Methods("Get")
	r.HandleFunc("/.well-known/ready", Ready).Methods("Get")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", url, port), r))

}

func User(w http.ResponseWriter, r *http.Request) {
	var backgroundColor, environment string
	if os.Getenv("BACKGROUND_COLOR") != "" {
		backgroundColor = os.Getenv("BACKGROUND_COLOR")
	} else {
		backgroundColor = "white"
	}

	if os.Getenv("ENV") != "" {
		environment = os.Getenv("ENV")
	} else {
		environment = "default"
	}

	htmlData := `<!DOCTYPE html>
	<html>
	<body style="background-color:` + backgroundColor + `;">
	<h2>User: John</h2>
	<h2>ID: 1</h2>
	<h2>Environment: ` + environment + `</h2>
	</body>
	</html>`

	htmlDataInBytes := []byte(htmlData)
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlDataInBytes)
}

func Live(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func Ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}
