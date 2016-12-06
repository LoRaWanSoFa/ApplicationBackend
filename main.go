// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	"github.com/LoRaWanSoFa/LoRaWanSoFa/webserver"
)

func main() {

	router := webserver.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	//test := MessageConverter.New()

	//bytes := []byte{0x3F, 0xB0, 0xFC, 0x00, 0xAA, 0xA0, 0x84, 0x41}
	//fmt.Println(test.ConvertSingleValue(bytes, 2))
}
