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
	log.SetFlags(log.LstdFlags | log.Llongfile)

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

		responseMembers, err := ioutil.ReadAll(res.Body)

		err = res.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		responseMemberList := getMembers.MemberList{}

		_ = json.Unmarshal(responseMembers, &responseMemberList)

		saveMembers(responseMemberList)

		memberList := getMemberIdList(responseMemberList)

		saveRelationChannelMember(channel, memberList)
	}
}
