package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint: homePage")
}

func testEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
	fmt.Println("Endpoint: test")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/test", testEndPoint)
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	handleRequests()
}
