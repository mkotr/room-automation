package main

import (
	"log"
	"net/http"
)

func main() {
	server := &ComputerServer{}

	if err := http.ListenAndServe(":1337", server); err != nil {
		log.Fatalf("could not listen on the port :1337 %v", err)
	}
}
