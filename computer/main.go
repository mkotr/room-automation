package computer

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(ComputerServer)
	if err := http.ListenAndServe(":1337", handler); err != nil {
		log.Fatalf("could not listen on the port :1337 %v", err)
	}
}
