package getPhoto

import (
	"errors"
	"github.com/revilon1991/tg-parser/internal/client"
)

func Handle(clientStorage *client.Storage, photoId int32) (string, error) {
	photo, _ := clientStorage.SendAndCatch(client.Request{
		"@type":       "downloadFile",
		"file_id":     photoId,
		"priority":    int32(32),
		"offset":      int32(0),
		"limit":       int32(0),
		"synchronous": true,
	})

	if photo.ResponseData["@type"].(string) == "error" {
		return "", errors.New(photo.ResponseData["message"].(string))
	}

	photoPath := photo.ResponseData["local"].(map[string]interface{})["path"].(string)

	return photoPath, nil
}
