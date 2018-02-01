package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/albertomurillo/k8s-demo/payload"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Port       int    `required:"true"`
	Greeting   string `default:"Hello"`
	BackendUrl string `required:"true" envconfig:"BACKEND_URL"`
}

var config Configuration

func main() {
	err := envconfig.Process("frontend", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/", ReturnMessage).Methods("GET")
	listenAddress := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(listenAddress, router))
}

func ReturnMessage(w http.ResponseWriter, r *http.Request) {
	var backendMessage payload.Message
	var response payload.Message

	resp, err := http.Get(config.BackendUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(respData, &backendMessage)
	response = payload.Message{
		Message: config.Greeting + " " + backendMessage.Message,
	}

	json.NewEncoder(w).Encode(response)
}
