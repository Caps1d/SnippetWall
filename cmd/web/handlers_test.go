package main

import (
	"net/http"
	"testing"

	"github.com/Caps1d/Lets-Go/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.Get(t, "/ping")

	// Check that the status code written by the ping handler was 200.
	assert.Equal(t, code, http.StatusOK)

	assert.Equal(t, string(body), "OK")
}
