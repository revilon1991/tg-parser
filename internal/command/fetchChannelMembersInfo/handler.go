package fetchChannelMembersInfo

import (
    "encoding/json"
    "fmt"
    "github.com/revilon1991/tg-parser/internal/config"
    "github.com/revilon1991/tg-parser/internal/useCase/getMembers"
    "github.com/revilon1991/tg-parser/internal/useCase/getUser"
    "github.com/urfave/cli"
    "io/ioutil"
    "log"
    "net/http"
    "time"
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

        memberInfoList := MemberInfoList{
            MemberList:        make(map[int]Member, 0),
            ChannelExternalId: channel.ChannelId,
            ChannelId:         channel.Id,
        }

        for userExternalId, responseMember := range responseMemberList.Members {
            joinChannelDate := time.Unix(int64(responseMember.JoinedChatDate), 0).Format("2006-01-02 15:04:05")

            user := getUserInfo(responseMember.UserId)

            memberInfoList.MemberList[userExternalId] = Member{
                ChannelId:       channel.Id,
                JoinChannelDate: joinChannelDate,
                UserExternalId:  responseMember.UserId,
                Username:        user.Username,
                FirstName:       user.FirstName,
                LastName:        user.LastName,
                PhoneNumber:     user.PhoneNumber,
                Type:            user.Type,
                Bio:             user.Bio,
            }
        }

        saveMembers(memberInfoList)
        fetchMemberIdList(memberInfoList)
        saveRelationChannelMember(memberInfoList)
    }
}

func getUserInfo(userId int) getUser.User {
    url := fmt.Sprintf(
        "http://localhost:8080%s?user_id=%d",
        config.Routing.V1GetUser,
        userId,
    )

    res, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    }

    responseMember, err := ioutil.ReadAll(res.Body)

    err = res.Body.Close()

    if err != nil {
        log.Fatal(err)
    }

    user := getUser.User{}

    _ = json.Unmarshal(responseMember, &user)

    return user
}
