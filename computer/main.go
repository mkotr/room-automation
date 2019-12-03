package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PingFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Successfully pinged the computer.")
}

func ShutdownFunc(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", PingFunc)

	fmt.Println("Running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
