package getStorageChannelList

import (
	"github.com/revilon1991/tg-parser/internal/connection/mysql"
	"log"
)

func Handle() ResponseChannelList {
	conn := mysql.Open()

	defer mysql.Close(conn)

	channelList, err := getChannelList(conn)

	if err != nil {
		log.Fatal(err)
	}

	var responseChannelList ResponseChannelList

	for _, channel := range channelList {
		responseChannelList = append(responseChannelList, ResponseChannel{
			Id:          channel.Id,
			OuterId:     channel.ChannelId,
			Username:    channel.Username.String,
			Description: channel.Description.String,
			MemberCount: channel.MemberCount,
			UpdatedAt:   channel.UpdatedAt,
		})
	}

	return responseChannelList
}
