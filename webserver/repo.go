package webserver

import "fmt"
import components "github.com/LoRaWanSoFa/ApplicationBackend/Components"

var currentID int64

var messages components.Messages

// Give us some seed data
func init() {
	RepoCreateMessage(components.MessageDownLink{ID: 1})
	RepoCreateMessage(components.MessageDownLink{ID: 2})
	RepoCreateMessage(components.MessageDownLink{ID: 3})
}

func RepoFindMessage(id int64) components.MessageDownLink {
	for _, t := range messages {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return components.MessageDownLink{}
}

func RepoCreateMessage(t components.MessageDownLink) components.MessageDownLink {
	currentID += 1
	t.ID = currentID
	messages = append(messages, t)
	return t
}

func RepoDestroyMessage(id int64) error {
	for i, t := range messages {
		if t.ID == id {
			messages = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Message with id of %d to delete", id)
}
