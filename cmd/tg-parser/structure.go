package main

import (
	"sync"
	"unsafe"
)

type ClientStorage struct {
	Client   unsafe.Pointer
	Updates  chan Update
	botToken string
	waiters  sync.Map
	Version  string
	CommitID string
}

type Update = map[string]interface{}
