package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	dist "github.com/LoRaWanSoFa/LoRaWanSoFa/Core/distributor"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}

func MessageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var messageId int
	var err error
	if messageId, err = strconv.Atoi(vars["messageId"]); err != nil {
		panic(err)
	}
	message := RepoFindMessage(int64(messageId))
	if message.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(message); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"deveui":"AFC147", "payload":"DESG6184FHAS"}' http://localhost:8080/messages
*/
func MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message components.MessageDownLink
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &message); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateMessage(message)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	dist.New().InputDownlink(message)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
