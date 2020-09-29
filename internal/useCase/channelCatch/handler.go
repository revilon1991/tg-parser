package channelCatch

import (
    "encoding/json"
    "github.com/revilon1991/tg-parser/internal/client"
)

func Handle(clientStorage *client.Storage, update client.Update) {
    var responseUpdateSupergroup ResponseUpdateSupergroup
    var responseSupergroupFullInfo ResponseSupergroupFullInfo

    _ = json.Unmarshal(update.Raw, &responseUpdateSupergroup)

    if responseUpdateSupergroup.Supergroup.Status.Type != "chatMemberStatusAdministrator" {
        return
    }

    res, _ := clientStorage.SendAndCatch(client.Request{
        "@type":         "getSupergroupFullInfo",
        "supergroup_id": responseUpdateSupergroup.Supergroup.Id,
    })

    _ = json.Unmarshal(res.Raw, &responseSupergroupFullInfo)

    channelId := responseUpdateSupergroup.Supergroup.Id
    username := responseUpdateSupergroup.Supergroup.Username
    memberCount := responseSupergroupFullInfo.MemberCount
    description := responseSupergroupFullInfo.Description

    saveChannel(channelId, username, memberCount, description)
}
