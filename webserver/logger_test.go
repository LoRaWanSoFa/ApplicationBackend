package webserver

import (
	"net/http"
	"testing"
)

func TestLogger(t *testing.T) {
	var h http.Handler
	result := Logger(h, "name")
	if result == h {
		t.Error("Something went wrong")
	}
}
