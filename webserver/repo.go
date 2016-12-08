package webserver

import "fmt"
import components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"

var currentId int64

var messages components.Messages

// Give us some seed data
func init() {
	RepoCreateMessage(components.MessageDownLink{Id: 1})
	RepoCreateMessage(components.MessageDownLink{Id: 2})
	RepoCreateMessage(components.MessageDownLink{Id: 3})
}

func RepoFindMessage(id int64) components.MessageDownLink {
	for _, t := range messages {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return components.MessageDownLink{}
}

func RepoCreateMessage(t components.MessageDownLink) components.MessageDownLink {
	currentId += 1
	t.Id = currentId
	messages = append(messages, t)
	return t
}

func RepoDestroyMessage(id int64) error {
	for i, t := range messages {
		if t.Id == id {
			messages = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Message with id of %d to delete", id)
}
