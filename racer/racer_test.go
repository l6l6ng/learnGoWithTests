package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	//slowURL := "http://www.facebook.com"
	//fastURL := "http://www.quii.co.uk"
	//
	//want := fastURL
	//
	//got := Racer(slowURL, fastURL)
	//
	//if got != want {
	//	t.Errorf("got '%s' want '%s'", got, want)
	//}

	t.Run("1", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastSever := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastSever.Close()

		slowURL := slowServer.URL
		fastURL := fastSever.URL

		want := fastURL
		got,_ := Racer(slowURL, fastURL,1*time.Second)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastSever := makeDelayedServer(12 * time.Second)

		defer slowServer.Close()
		defer fastSever.Close()

		_, err := Racer(slowServer.URL, fastSever.URL,1*time.Second)

		if err == nil {
			t.Errorf("expectecd an error but din't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
