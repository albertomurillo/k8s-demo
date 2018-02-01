package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/albertomurillo/k8s-demo/payload"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Port    int    `required:"true"`
	Message string `default:"World!"`
}

var config Configuration

func main() {
	err := envconfig.Process("backend", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/", ReturnMessage).Methods("GET")
	listenAddress := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(listenAddress, router))
}

func ReturnMessage(w http.ResponseWriter, r *http.Request) {
	var response = payload.Message{
		Message: config.Message,
	}
	json.NewEncoder(w).Encode(response)
}
