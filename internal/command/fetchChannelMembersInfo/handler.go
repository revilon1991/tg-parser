package fetchChannelMembersInfo

import (
	"encoding/json"
	"fmt"
	"github.com/revilon1991/tg-parser/internal/config"
	"github.com/revilon1991/tg-parser/internal/useCase/getMembers"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
)

func Handle(c *cli.Context) {
	channelList := getChannelList()

	for _, channel := range channelList {
		url := fmt.Sprintf(
			"http://localhost:8080%s?channel_id=%d",
			config.Routing.V1GetMembers,
			channel.ChannelId,
		)

		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		members, err := ioutil.ReadAll(res.Body)

		err = res.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		memberList := getMembers.MemberList{}

		_ = json.Unmarshal(members, &memberList)

		saveMembers(memberList)
	}
}
