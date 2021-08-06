package racer

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("common", func(t *testing.T){
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
	    fastURL := fastServer.URL

	    want := fastURL
	    got, _ := Racer(slowURL, fastURL)

	    if got != want {
	    	t.Errorf("got '%s', want '%s'", got, want)
	    }

	    // slowServer.Close()
	    // fastServer.Close()
	})

	t.Run("return an error if a server dosen't respond with 10s", func(t *testing.T){
		server := makeDelayServer(25 * time.Millisecond)

		defer server.Close()
	 
	    _, err := ConfigurableRacer(server.URL, server.URL, 20 * time.Millisecond)

	    if err == nil {
	    	t.Error("Expected an error but didn't get one")
	    }
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}