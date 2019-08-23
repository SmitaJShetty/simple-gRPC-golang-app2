package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("app2 listening ...")
	addServe()
}

func addServe() {
	hostURL := ":8080"
	http.HandleFunc("/app2", handlerFunc)
	log.Fatal(http.ListenAndServe(hostURL, nil))
}
