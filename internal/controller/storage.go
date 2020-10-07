package controller

import (
	"encoding/json"
	"github.com/revilon1991/tg-parser/internal/config"
	"github.com/revilon1991/tg-parser/internal/useCase/getStorageChannelList"
	"github.com/revilon1991/tg-parser/internal/useCase/getStorageMemberList"
	"net/http"
	"strconv"
)

func GetStorageChannelList() {
	http.HandleFunc(config.Routing.V1StorageGetChannelList, func(w http.ResponseWriter, r *http.Request) {
		responseChannelList := getStorageChannelList.Handle()

		js, err := json.Marshal(responseChannelList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}

func GetStorageMemberList() {
	http.HandleFunc(config.Routing.V1StorageGetMemberList, func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["channel_id"]

		if !ok || len(keys[0]) < 1 {
			http.Error(w, "Missing query param 'channel_id'", http.StatusBadRequest)
			return
		}
		channelId, _ := strconv.Atoi(keys[0])

		memberList, err := getStorageMemberList.Handle(channelId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(memberList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}
