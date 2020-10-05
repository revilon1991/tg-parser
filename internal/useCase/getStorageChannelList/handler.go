package getStorageChannelList

import (
	"github.com/revilon1991/tg-parser/internal/connection/mysql"
	"log"
)

func Handle() ChannelList {
	conn := mysql.Open()

	defer mysql.Close(conn)

	channelList, err := getChannelList(conn)

	if err != nil {
		log.Fatal(err)
	}

	var responseChannelList ChannelList

	for _, channel := range channelList {
		responseChannelList = append(responseChannelList, Channel{
			Id:          channel.Id,
			ChannelId:   channel.ChannelId,
			Username:    channel.Username.String,
			Description: channel.Description.String,
			MemberCount: channel.MemberCount,
			CreatedAt:   channel.CreatedAt,
			UpdatedAt:   channel.UpdatedAt,
		})
	}

	return responseChannelList
}
