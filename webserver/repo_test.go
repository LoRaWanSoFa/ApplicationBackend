package webserver

import (
	"testing"

	components "github.com/LoRaWanSoFa/ApplicationBackend/Components"
)

func TestRepoFindMessage(t *testing.T) {
	expected := RepoCreateMessage(components.MessageDownLink{ID: 1})
	result := RepoFindMessage(2)
	if result == expected {
		t.Error("Should not be equal")
	}
}

func TestRepoDestroyMessage(t *testing.T) {
	testMessage := RepoCreateMessage(components.MessageDownLink{ID: 5})
	RepoDestroyMessage(5)
	if RepoFindMessage(5) == testMessage {
		t.Error("This message should be deleted")
	}
	RepoDestroyMessage(5)
}
