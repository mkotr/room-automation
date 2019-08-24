package computer

import (
	"fmt"
	"net/http"
)

//ComputerServer - The server
func ComputerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Successfully shutdown the computer.")
}
