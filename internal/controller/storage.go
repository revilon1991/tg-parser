package controller

import (
	"encoding/json"
	"github.com/revilon1991/tg-parser/internal/config"
	"github.com/revilon1991/tg-parser/internal/useCase/getStorageChannelList"
	"net/http"
)

func GetStorageChannelList() {
	http.HandleFunc(config.Routing.V1StorageGetChannelList, func(w http.ResponseWriter, r *http.Request) {
		channelList := getStorageChannelList.Handle()

		js, err := json.Marshal(channelList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}
