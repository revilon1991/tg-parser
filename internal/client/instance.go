package client

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
	"github.com/davecgh/go-spew/spew"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

const (
	Port = "8080"
)

var (
	Version = "0"
)

func NewClient() *Storage {
	// Seed rand with time
	rand.Seed(time.Now().UnixNano())

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	setLogConfig()

	clientStorage := Storage{
		Client:   C.td_json_client_create(),
		Updates:  make(chan Update, 100),
		BotToken: botToken,
		Version:  Version,
	}

	handleCtrlC(&clientStorage)

	go func() {
		for {
			if clientStorage.AuthorizationState == "authorizationStateClosing" {
				break
			}

			// get response
			result := C.td_json_client_receive(clientStorage.Client, C.double(10))
			var response Response
			var receiveUpdate Update
			var responseData map[string]interface{}
			var raw = []byte(C.GoString(result))

			_ = json.Unmarshal(raw, &responseData)
			spew.Dump(responseData)
			// does new response has @extra field?
			if extra, hasExtra := responseData["@extra"].(string); hasExtra {
				// trying to load response with this salt
				if waiter, found := clientStorage.waiters.Load(extra); found {
					// found? send it to waiter channel
					response.ResponseData = responseData
					response.Raw = raw

					waiter.(chan Response) <- response

					// trying to prevent memory leak
					close(waiter.(chan Response))
				}
			} else {
				// does new updates has @type field?
				if _, hasType := responseData["@type"]; hasType {
					// if yes, send it in main channel
					receiveUpdate.UpdateData = responseData
					receiveUpdate.Raw = raw

					clientStorage.Updates <- receiveUpdate
				}
			}
		}
	}()

	return &clientStorage
}

func (clientStorage *Storage) SendAndCatch(jsonQuery interface{}) (Response, error) {
	request := make(Request)

	switch jsonQuery.(type) {
	case string:
		// unmarshal JSON into map, we don't have @extra field, if user don't set it
		_ = json.Unmarshal([]byte(jsonQuery.(string)), &request)
	case Request:
		request = jsonQuery.(Request)
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
	request["@extra"] = randomString

	// create waiter chan and save it in Waiters
	waiter := make(chan Response, 1)
	clientStorage.waiters.Store(randomString, waiter)

	// send it through already implemented method
	var query *C.char

	jsonBytes, _ := json.Marshal(request)
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
		return Response{}, errors.New("timeout")
	}
}

func (clientStorage *Storage) Close() {
	C.td_json_client_destroy(clientStorage.Client)
	os.Exit(0)
}

func handleCtrlC(clientStorage *Storage) {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		_, _ = clientStorage.SendAndCatch(Request{
			"@type": "close",
		})

		time.Sleep(2 * time.Second)
		fmt.Println("\nWait timeout 5 second for close tdlib client...")
		time.Sleep(5 * time.Second)
		fmt.Println("Force close tdlib client!")
		os.Exit(1)
	}()
}

func setLogConfig() {
	queryLogVerbosityLevel, _ := json.Marshal(struct {
		Type              string `json:"@type"`
		NewVerbosityLevel int    `json:"new_verbosity_level"`
	}{Type: "setLogVerbosityLevel", NewVerbosityLevel: 1})

	queryLogVerbosityLevelC := C.CString(string(queryLogVerbosityLevel))
	C.td_json_client_execute(nil, queryLogVerbosityLevelC)
	C.free(unsafe.Pointer(queryLogVerbosityLevelC))

	queryLogStreamFile, _ := json.Marshal(struct {
		Type      string      `json:"@type"`
		LogStream interface{} `json:"log_stream"`
	}{Type: "setLogStream", LogStream: struct {
		Type        string `json:"@type"`
		Path        string `json:"path"`
		MaxFileSize int64  `json:"max_file_size"`
	}{Type: "logStreamFile", Path: "./var/tdlib/errors.txt", MaxFileSize: 10485760}})

	queryLogStreamFileC := C.CString(string(queryLogStreamFile))
	C.td_json_client_execute(nil, queryLogStreamFileC)
	C.free(unsafe.Pointer(queryLogStreamFileC))
}
