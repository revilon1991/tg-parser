package client

import (
	"sync"
	"unsafe"
)

type Storage struct {
	Client             unsafe.Pointer
	AuthorizationState string
	Updates            chan Update
	BotToken           string
	waiters            sync.Map
	Version            string
	CommitID           string
}

type Update struct {
	UpdateData map[string]interface{}
	Raw        []byte
}
type Request map[string]interface{}

type Response struct {
	ResponseData map[string]interface{}
	Raw          []byte
}
