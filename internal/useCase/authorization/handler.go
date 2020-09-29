package authorization

import (
    "errors"
    "fmt"
    "github.com/revilon1991/tg-parser/internal/client"
    "os"
    "runtime"
    "strconv"
)

func Handle(clientStorage *client.Storage, authorizationState string) (client.Response, error) {
    fmt.Println(authorizationState)

    switch authorizationState {
    case "authorizationStateWaitTdlibParameters":
        appIdString := os.Getenv("TELEGRAM_APP_ID")
        appHash := os.Getenv("TELEGRAM_APP_HASH")
        appId64, _ := strconv.Atoi(appIdString)
        appId := int32(appId64)

        res, err := clientStorage.SendAndCatch(client.Request{
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
                DatabaseDirectory:      "./var/tdlib_" + runtime.GOOS,
                FilesDirectory:         "./var/tdlib_" + runtime.GOOS,
                UseFileDatabase:        true,
                UseChatInfoDatabase:    true,
                UseMessageDatabase:     true,
                UseSecretChats:         false,
                ApiId:                  appId,
                ApiHash:                appHash,
                SystemLanguageCode:     "ru",
                DeviceModel:            runtime.GOOS,
                SystemVersion:          runtime.Version(),
                ApplicationVersion:     clientStorage.Version,
                EnableStorageOptimizer: true,
                IgnoreFileNames:        true,
            },
        })

        if err != nil {
            return res, err
        }
        return res, nil
    case "authorizationStateWaitEncryptionKey":
        res, err := clientStorage.SendAndCatch(client.Request{
            "@type": "checkDatabaseEncryptionKey",
        })
        if err != nil {
            return res, err
        }
        return res, nil
    case "authorizationStateWaitPhoneNumber":
        if clientStorage.BotToken != "" {
            res, err := clientStorage.SendAndCatch(client.Request{
                "@type": "checkAuthenticationBotToken",
                "token": clientStorage.BotToken,
            })

            return res, err
        }

        fmt.Print("Enter phone: ")
        var number string
        _, _ = fmt.Scanln(&number)

        res, err := clientStorage.SendAndCatch(client.Request{
            "@type":        "setAuthenticationPhoneNumber",
            "phone_number": number,
        })
        if err != nil {
            return res, err
        }
        return res, nil
    case "authorizationStateWaitCode":
        fmt.Print("Enter code: ")
        var code string
        _, _ = fmt.Scanln(&code)

        res, err := clientStorage.SendAndCatch(client.Request{
            "@type": "checkAuthenticationCode",
            "code":  code,
        })
        if err != nil {
            return res, err
        }
        return res, nil
    case "authorizationStateReady":
        fmt.Println("Authorized!")
        return client.Response{}, nil
    case "authorizationStateClosing":
        return client.Response{}, nil
    case "authorizationStateClosed":
        clientStorage.Close()
        return client.Response{}, nil
    default:
        return client.Response{}, errors.New(fmt.Sprintf("unexpected authorization state: %s", authorizationState))
    }
}
