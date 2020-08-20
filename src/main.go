package main

//#cgo linux CFLAGS: -I/usr/local/include
//#cgo darwin CFLAGS: -I/usr/local/include
//#cgo windows CFLAGS: -IC:/src/td -IC:/src/td/build
//#cgo linux,!tdjson_static LDFLAGS: -L/lol/usr/local/lib -ltdjson
//#cgo linux,tdjson_static LDFLAGS: -L/lol/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm
//#cgo darwin LDFLAGS: -L/usr/local/lib -L/usr/local/opt/openssl/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm
//#cgo windows LDFLAGS: -LC:/src/td/build/Debug -ltdjson
//#include <stdlib.h>
//#include <td/telegram/td_json_client.h>
//#include <td/telegram/td_log.h>
import "C"

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"net/http"
	"time"
	"unsafe"
)

func main() {
	_ = godotenv.Load()

	setLogConfig()
	clientStorage := newClient()
	handleCtrlC(clientStorage)

	go func() {
		//Update loop
		for update := range clientStorage.Updates {
			//Print all updates
			fmt.Println(update)

			//Authorization block
			if update["@type"].(string) == "updateAuthorizationState" {
				if authorizationState, ok := update["authorization_state"].(map[string]interface{})["@type"].(string); ok {
					res, err := clientStorage.Auth(authorizationState)

					if err != nil {
						log.Println(err)
					}

					log.Println(res)
				}
			}
		}
	}()

	runWebServer(clientStorage)
}

func runWebServer(clientStorage *ClientStorage) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting server...")
	getMe(clientStorage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Sends request to the TDLib client and catches the result in updates channel.
// You can provide string or Update.
func (clientStorage *ClientStorage) SendAndCatch(jsonQuery interface{}) (Update, error) {
	update := make(Update)

	switch jsonQuery.(type) {
	case string:
		// unmarshal JSON into map, we don't have @extra field, if user don't set it
		json.Unmarshal([]byte(jsonQuery.(string)), &update)
	case Update:
		update = jsonQuery.(Update)
	}

	// letters for generating random string
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// generate random string for @extra field
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	randomString := string(b)

	// set @extra field
	update["@extra"] = randomString

	// create waiter chan and save it in Waiters
	waiter := make(chan Update, 1)
	clientStorage.waiters.Store(randomString, waiter)

	// send it through already implemented method
	var query *C.char

	jsonBytes, _ := json.Marshal(update)
	query = C.CString(string(jsonBytes))

	defer C.free(unsafe.Pointer(query))
	C.td_json_client_send(clientStorage.Client, query)

	select {
	// wait response from main loop in NewClient()
	case response := <-waiter:
		return response, nil
		// or timeout
	case <-time.After(10 * time.Second):
		clientStorage.waiters.Delete(randomString)
		return Update{}, errors.New("timeout")
	}
}
