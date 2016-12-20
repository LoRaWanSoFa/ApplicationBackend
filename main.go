// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	"github.com/LoRaWanSoFa/ApplicationBackend/webserver"
)

func main() {
	router := webserver.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
