package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGET(t *testing.T) {

	server := &ComputerServer{}

	t.Run("returns 200 OK if there's a signal running with message to path '/'", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Successfully pinged the server."

		assertResponseBody(t, got, want)
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 when other paths with message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/randompath", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Nothing here. POST to /shutdown or GET to / for healthcheck"

		assertResponseBody(t, got, want)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestPOST(t *testing.T) {

	server := ComputerServer{}
	t.Run("returns 202 status accepted on /shutdown and message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/shutdown", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Successfully sent the shut down command"

		assertResponseBody(t, got, want)
		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("returns 404 status accepted all other paths /shutdown and message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/randompath", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Nothing here. POST to /shutdown or GET to / for healthcheck"

		assertResponseBody(t, got, want)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestOtherMethodsAnyPath(t *testing.T) {
	server := ComputerServer{}

	t.Run("Return 404 on other methods with a message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPut, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "Nothing here. POST to /shutdown or GET to / for healthcheck"

		assertResponseBody(t, got, want)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got code %v, want %v", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
