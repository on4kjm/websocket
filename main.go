package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/on4kjm/websocket/serial_mock"
	"github.com/on4kjm/websocket/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// our new stats function which will expose any YouTube
// stats via a websocket connection
func stats(w http.ResponseWriter, r *http.Request) {
	// we call our new websocket package Upgrade
	// function in order to upgrade the connection
	// from a standard HTTP connection to a websocket one
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	// we then call our Writer function
	// which continually polls and writes the results
	// to this websocket connection
	go websocket.Writer(ws)
}

func setupRoutes(messageBuffer chan string) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/morserino", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Websocket test")

	messageBuffer := make(chan string)

	go serial_mock.Receive_loop()

	setupRoutes(messageBuffer)
}
