package getChannelInfo

import (
    "encoding/json"
    "github.com/revilon1991/tg-parser/internal/client"
)

func Handle(clientStorage *client.Storage, channelId int32) *ResponseSupergroupFullInfo {
    var responseSupergroupFullInfo ResponseSupergroupFullInfo

    res, _ := clientStorage.SendAndCatch(client.Request{
        "@type":         "getSupergroupFullInfo",
        "supergroup_id": channelId,
    })

    _ = json.Unmarshal(res.Raw, &responseSupergroupFullInfo)

    return &responseSupergroupFullInfo
}
