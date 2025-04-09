package main

import (
	"net/http"
	"os"
	"testing"
)

const (
	TEST_USERNAME = "test"
	TEST_PASSWORD = "test"
)

func TestMain(m *testing.M) {
	ConfigServerUsername = TEST_USERNAME
	ConfigServerPassword = TEST_PASSWORD

	s := mkServer()
	go func() {
		s.ListenAndServe()
	}()

	code := m.Run()

	s.Close()
	os.Exit(code)
}

func TestServerReturnsUnauthorizedWhenNoCredentialsAreProvided(t *testing.T) {
	res, _ := http.Get("http://localhost:3000/api/v1/dev")
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected response status code of unauthenticated request to be 401")
	}
}

func TestServerReturnsOkWhenCredentialsAreProvided(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/dev", nil)
	req.SetBasicAuth(TEST_USERNAME, TEST_PASSWORD)

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected response status code of authenticated request to be 200")
	}
}
