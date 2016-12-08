package webserver

import "testing"

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	if router != router {
		t.Error("was hier los")
	}
}
