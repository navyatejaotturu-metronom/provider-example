package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var port, url string
	var exists bool

	if url, exists = os.LookupEnv("DRP_CF_HTTP_ADDR"); !exists {
		url = "0.0.0.0"
	}

	if port, exists = os.LookupEnv("DRP_CF_HTTP_PORT"); !exists {
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
	var exists bool
	if backgroundColor, exists = os.LookupEnv("BACKGROUND_COLOR"); !exists {
		backgroundColor = "white"
	}

	if environment, exists = os.LookupEnv("BACKGROUND_COLOR"); !exists {
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
	var livenessResponse string
	var exists bool

	if livenessResponse, exists = os.LookupEnv("RESPONSE_CODE"); !exists {
		livenessResponse = "204"
	}

	response, err := strconv.Atoi(livenessResponse)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(response)
}

func Ready(w http.ResponseWriter, r *http.Request) {
	var readinessResponse string
	var exists bool

	if readinessResponse, exists = os.LookupEnv("RESPONSE_CODE"); !exists {
		readinessResponse = "204"
	}

	response, err := strconv.Atoi(readinessResponse)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(response)
}
