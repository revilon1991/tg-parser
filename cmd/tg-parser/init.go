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
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

func handleCtrlC(clientStorage *ClientStorage) {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		C.td_json_client_destroy(clientStorage.Client)
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

func newClient() *ClientStorage {
	// Seed rand with time
	rand.Seed(time.Now().UnixNano())

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	clientStorage := ClientStorage{
		Client:   C.td_json_client_create(),
		Updates:  make(chan Update, 100),
		botToken: botToken,
		Version:  Version,
		CommitID: CommitID,
	}

	go func() {
		for {
			// get update
			result := C.td_json_client_receive(clientStorage.Client, C.double(10))
			var update map[string]interface{}
			_ = json.Unmarshal([]byte(C.GoString(result)), &update)
			// does new update has @extra field?
			if extra, hasExtra := update["@extra"].(string); hasExtra {
				// trying to load update with this salt
				if waiter, found := clientStorage.waiters.Load(extra); found {
					// found? send it to waiter channel
					waiter.(chan Update) <- update

					// trying to prevent memory leak
					close(waiter.(chan Update))
				}
			} else {
				// does new updates has @type field?
				if _, hasType := update["@type"]; hasType {
					// if yes, send it in main channel
					clientStorage.Updates <- update
				}
			}
		}
	}()

	return &clientStorage
}

//Method for interactive authorizations process, just provide it authorization state from updates and api credentials.
func (clientStorage *ClientStorage) Auth(authorizationState string) (Update, error) {
	fmt.Println(authorizationState)

	switch authorizationState {
	case "authorizationStateWaitTdlibParameters":
		appIdString := os.Getenv("TELEGRAM_APP_ID")
		appHash := os.Getenv("TELEGRAM_APP_HASH")
		appId64, _ := strconv.Atoi(appIdString)
		appId := int32(appId64)

		res, err := clientStorage.SendAndCatch(Update{
			"@type": "setTdlibParameters",
			"parameters": struct {
				Type                   string `json:"@type"`
				UseTestDc              bool   `json:"use_test_dc"`
				DatabaseDirectory      string `json:"database_directory"`
				FilesDirectory         string `json:"files_directory"`
				UseFileDatabase        bool   `json:"use_file_database"`
				UseChatInfoDatabase    bool   `json:"use_chat_info_database"`
				UseMessageDatabase     bool   `json:"use_message_database"`
				UseSecretChats         bool   `json:"use_secret_chats"`
				ApiId                  int32  `json:"api_id"`
				ApiHash                string `json:"api_hash"`
				SystemLanguageCode     string `json:"system_language_code"`
				DeviceModel            string `json:"device_model"`
				SystemVersion          string `json:"system_version"`
				ApplicationVersion     string `json:"application_version"`
				EnableStorageOptimizer bool   `json:"enable_storage_optimizer"`
				IgnoreFileNames        bool   `json:"ignore_file_names"`
			}{
				Type:                   "tdlibParameters",
				UseTestDc:              false,
				DatabaseDirectory:      "./var/tdlib",
				FilesDirectory:         "./var/tdlib",
				UseFileDatabase:        true,
				UseChatInfoDatabase:    true,
				UseMessageDatabase:     true,
				UseSecretChats:         false,
				ApiId:                  appId,
				ApiHash:                appHash,
				SystemLanguageCode:     "ru",
				DeviceModel:            runtime.GOOS,
				SystemVersion:          runtime.Version(),
				ApplicationVersion:     Version,
				EnableStorageOptimizer: true,
				IgnoreFileNames:        true,
			},
		})

		if err != nil {
			return nil, err
		}
		return res, nil
	case "authorizationStateWaitEncryptionKey":
		res, err := clientStorage.SendAndCatch(Update{
			"@type": "checkDatabaseEncryptionKey",
		})
		if err != nil {
			return nil, err
		}
		return res, nil
	case "authorizationStateWaitPhoneNumber":
		if clientStorage.botToken != "" {
			res, err := clientStorage.SendAndCatch(Update{
				"@type": "checkAuthenticationBotToken",
				"token": clientStorage.botToken,
			})

			return res, err
		}

		fmt.Print("Enter phone: ")
		var number string
		_, _ = fmt.Scanln(&number)

		res, err := clientStorage.SendAndCatch(Update{
			"@type":        "setAuthenticationPhoneNumber",
			"phone_number": number,
		})
		if err != nil {
			return nil, err
		}
		return res, nil
	case "authorizationStateWaitCode":
		fmt.Print("Enter code: ")
		var code string
		_, _ = fmt.Scanln(&code)

		res, err := clientStorage.SendAndCatch(Update{
			"@type": "checkAuthenticationCode",
			"code":  code,
		})
		if err != nil {
			return nil, err
		}
		return res, nil
	case "authorizationStateReady":
		fmt.Println("Authorized!")
		return nil, nil
	default:
		return nil, errors.New(fmt.Sprintf("unexpected authorization state: %s", authorizationState))
	}
}
