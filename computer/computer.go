package main

import (
	"fmt"
	"net/http"
)

//ComputerServer
type ComputerServer struct {
}

func (c *ComputerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.processGet(w, r)
	case http.MethodPost:
		c.processPost(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Nothing here. POST to /shutdown or GET to / for healthcheck")
	}

}

func (c *ComputerServer) processGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Successfully pinged the server.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Nothing here. POST to /shutdown or GET to / for healthcheck")
	}
}

func (c *ComputerServer) processPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/shutdown" {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Successfully sent the shut down command")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Nothing here. POST to /shutdown or GET to / for healthcheck")
	}

}
