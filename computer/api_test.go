package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	t.Run("test get endpoint", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		handler := http.HandlerFunc(PingFunc)

		handler.ServeHTTP(res, req)

		got := res.Body.String()
		want := "Successfully pinged the computer."

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		if res.Code != http.StatusOK {
			t.Errorf("got %v, expected %v", res.Code, http.StatusOK)
		}
	})

	t.Run("test post endpoint to return status OK if successful", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/shutdown", nil)

		res := httptest.NewRecorder()

		handler := http.HandlerFunc(ShutdownFunc)
		handler.ServeHTTP(res, req)

		if res.Code != http.StatusAccepted {
			t.Errorf("got %v, want % v", res.Code, http.StatusAccepted)
		}
	})
}
