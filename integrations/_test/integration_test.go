package integration_test

import (
	"net/http/httptest"
	"testing"

	//	. "github.com/cloudnativego/websockets/server"
	"github.com/codegangsta/negroni"
)

var (
	server   *negroni.Negroni
	recorder *httptest.ResponseRecorder
)

func TestIntegration(t *testing.T) {

	// Create server
	/*	appEnv, err := cfenv.Current()
		if err != nil {
			t.Fatalf("Environment not set: %s", err)
		}
		server = NewServer(appEnv)

		// Unauthenticated user can access home page
		getHomePageRequest, _ := http.NewRequest("GET", "/", nil)
		recorder = httptest.NewRecorder()
		server.ServeHTTP(recorder, getHomePageRequest)
		if recorder.Code != http.StatusOK {
			t.Errorf("Expected response code to be %d, received: %d", http.StatusOK, recorder.Code)
		}

		// Unauthenticated user cannot access user page
		getUserPageRequest, _ := http.NewRequest("GET", "/user", nil)
		recorder = httptest.NewRecorder()
		server.ServeHTTP(recorder, getUserPageRequest)
		if recorder.Code != http.StatusMovedPermanently {
			t.Errorf("Expected response code to be %d, received: %d", http.StatusMovedPermanently, recorder.Code)
		} */

	// HACK go back and re-enable this test.
}
