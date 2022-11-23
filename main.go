package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type message struct {
	DeviceID string `json:"deviceId"`
	Payload  string `json:"payload"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From OpenShift %s!", r.URL.Path[1:])
	//fmt.Println("RESTfulServ. on:8093, Controller:", r.URL.Path[1:])

	var msg message
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Print("ID: ", msg.DeviceID)
	fmt.Println(" Payload: ", msg.Payload)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting Restful services v1.10...")
	fmt.Println("Using port:8093")
	err := http.ListenAndServe(":8093", nil)
	log.Print(err)
	errorHandler(err)
}
func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
