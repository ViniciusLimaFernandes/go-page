package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	applicationPort := ":8000"

	fmt.Println("Application started! Listening on port", applicationPort)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(applicationPort, nil))
}
