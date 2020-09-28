package getUser

import (
	"encoding/json"
	"fmt"
	"github.com/revilon1991/tg-parser/internal/client"
	"github.com/revilon1991/tg-parser/internal/config"
)

func Handle(clientStorage *client.Storage, userId int32) User {
	var responseUser ResponseUser
	var responsePhotos ResponsePhotos

	resUser, _ := clientStorage.SendAndCatch(client.Request{
		"@type":   "getUser",
		"user_id": userId,
	})

	_ = json.Unmarshal(resUser.Raw, &responseUser)

	resUserFull, _ := clientStorage.SendAndCatch(client.Request{
		"@type":   "getUserFullInfo",
		"user_id": userId,
	})

	user := User{
		Id:          responseUser.Id,
		FirstName:   responseUser.FirstName,
		LastName:    responseUser.LastName,
		Username:    responseUser.Username,
		PhoneNumber: responseUser.PhoneNumber,
		Bio:         resUserFull.ResponseData["bio"].(string),
		Type:        responseUser.Type.Type,
	}

	resPhoto, _ := clientStorage.SendAndCatch(client.Request{
		"@type":   "getUserProfilePhotos",
		"user_id": userId,
		"offset":  int32(0),
		"limit":   int32(100),
	})

	_ = json.Unmarshal(resPhoto.Raw, &responsePhotos)

	for _, photo := range responsePhotos.PhotoList {
		photoSizes := photo.Sizes
		photoId := photoSizes[len(photoSizes)-1].Photo.Id

		photoUrl := fmt.Sprintf(config.Routing.V1GetPhoto+"?photo_id=%d", photoId)

		user.PhotoList = append(user.PhotoList, photoUrl)
	}

	return user
}
