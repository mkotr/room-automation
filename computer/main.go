package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type InMemoryComputerStore struct{}

var execCommand = exec.Command

func (i *InMemoryComputerStore) Shutdown() (string, error) {
	cmd := exec.Command("shutdown", "-s", "-t", "10")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return out.String(), err
	}
	return out.String(), nil
}

func main() {
	fmt.Println("Running the computer api")
	store := InMemoryComputerStore{}
	server := &ComputerServer{&store}
	if err := http.ListenAndServe(":1337", server); err != nil {
		log.Fatalf("could not listen on the port :1337 %v", err)
	}
}
