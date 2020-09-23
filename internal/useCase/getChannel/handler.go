package getChannel

import (
	"encoding/json"
	"github.com/revilon1991/tg-parser/internal/client"
)

func Handle(clientStorage *client.Storage, channelId int32) *ResponseSupergroup {
	var responseSupergroup ResponseSupergroup

	res, _ := clientStorage.SendAndCatch(client.Request{
		"@type":         "getSupergroup",
		"supergroup_id": channelId,
	})

	_ = json.Unmarshal(res.Raw, &responseSupergroup)

	return &responseSupergroup
}
