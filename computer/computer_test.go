package computer

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHealthcheck(t *testing.T) {
	t.Run("returns 200 OK if there's a signal running with message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
		response := httptest.NewRecorder()

		ComputerServer(response, request)

		got := response.Body.String()
		want := "Successfully shutdown the computer."

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
