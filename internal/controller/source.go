package controller

import (
	"encoding/json"
	"github.com/revilon1991/tg-parser/internal/client"
	"github.com/revilon1991/tg-parser/internal/config"
	"github.com/revilon1991/tg-parser/internal/useCase/getChannel"
	"github.com/revilon1991/tg-parser/internal/useCase/getChannelInfo"
	"github.com/revilon1991/tg-parser/internal/useCase/getMe"
	"github.com/revilon1991/tg-parser/internal/useCase/getMembers"
	"github.com/revilon1991/tg-parser/internal/useCase/getPhoto"
	"github.com/revilon1991/tg-parser/internal/useCase/getUser"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetMeAction(clientStorage *client.Storage) {
	http.HandleFunc(config.Routing.V1GetMe, func(w http.ResponseWriter, r *http.Request) {
		me := getMe.Handle(clientStorage)

		js, err := json.Marshal(me)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}

func GetMembersAction(clientStorage *client.Storage) {
	http.HandleFunc(config.Routing.V1GetMembers, func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["channel_id"]

		if !ok || len(keys[0]) < 1 {
			http.Error(w, "Missing query param 'channel_id'", http.StatusBadRequest)
			return
		}
		channelId64, _ := strconv.Atoi(keys[0])
		channelId := int32(channelId64)

		memberList := getMembers.Handle(clientStorage, channelId)

		js, err := json.Marshal(memberList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}

func GetUserAction(clientStorage *client.Storage) {
	http.HandleFunc(config.Routing.V1GetUser, func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["user_id"]

		if !ok || len(keys[0]) < 1 {
			http.Error(w, "Missing query param 'user_id'", http.StatusBadRequest)
			return
		}
		userId64, _ := strconv.Atoi(keys[0])
		userId := int32(userId64)

		user := getUser.Handle(clientStorage, userId)

		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}

func GetPhotoAction(clientStorage *client.Storage) {
	http.HandleFunc(config.Routing.V1GetPhoto, func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["photo_id"]

		if !ok || len(keys[0]) < 1 {
			http.Error(w, "Missing query param 'photo_id'", http.StatusBadRequest)
			return
		}
		userId64, _ := strconv.Atoi(keys[0])
		userId := int32(userId64)

		photoPath, err := getPhoto.Handle(clientStorage, userId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		img, err := os.Open(photoPath)

		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err := img.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		w.Header().Set("Content-Type", "image/jpeg")
		_, _ = io.Copy(w, img)
	})
}

func GetChannelInfoAction(clientStorage *client.Storage) {
	http.HandleFunc(config.Routing.V1GetChannelInfo, func(w http.ResponseWriter, r *http.Request) {
		channelKeys, _ := r.URL.Query()["channel_id"]

		channel64, _ := strconv.Atoi(channelKeys[0])
		channelId := int32(channel64)

		result := getChannelInfo.Handle(clientStorage, channelId)

		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}

func GetChannelAction(clientStorage *client.Storage) {
	http.HandleFunc(config.Routing.V1GetChannel, func(w http.ResponseWriter, r *http.Request) {
		channelKeys, _ := r.URL.Query()["channel_id"]

		channel64, _ := strconv.Atoi(channelKeys[0])
		channelId := int32(channel64)

		result := getChannel.Handle(clientStorage, channelId)

		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)
	})
}
