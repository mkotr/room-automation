package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

func TestShutDownIntegration(t *testing.T) {
	store := InMemoryComputerStore{}
	server := ComputerServer{&store}

	request, err := http.NewRequest(http.MethodPost, "/shutdown", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assertStatus(t, response.Code, http.StatusAccepted)
	assertResponseBody(t, response.Body.String(), "Successfully sent the shutdown command")

	if err != nil {
		t.Errorf("got err %v when we expected none", err)
	}
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

const shutdownRunResult = "foo!"

// This test will make your comp shut down lol  - dunno how to test this command.
func TestInMemoryStore_Shutdown(t *testing.T) {
	store := InMemoryComputerStore{}

	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := store.Shutdown()

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if string(out) != shutdownRunResult {
		t.Errorf("Expected %q, got %q", shutdownRunResult, out)
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// some code here to check arguments perhaps?
	fmt.Fprintf(os.Stdout, shutdownRunResult)
	os.Exit(0)
}
