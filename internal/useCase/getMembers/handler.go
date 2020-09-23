package getMembers

import (
	"encoding/json"
	"github.com/revilon1991/tg-parser/internal/client"
)

const (
	SearchTextLength = 2
	LimitMemberCount = 200
)

func Handle(clientStorage *client.Storage, channelId int32) *MemberList {
	var memberList = MemberList{
		Total:   0,
		Members: make(map[int]Member),
	}

	searchTextMapList := getSearchTextMapList()

	getMemberListBySearchText(clientStorage, &memberList, channelId, searchTextMapList)

	memberList.Total = len(memberList.Members)

	return &memberList
}

func getMemberListBySearchText(clientStorage *client.Storage, memberList *MemberList, channelId int32, searchTextMapList SearchTextMapList) {
	for searchText, nestedSearchText := range searchTextMapList {
		responseMemberList := queryMemberList(clientStorage, channelId, searchText)

		for _, responseMember := range responseMemberList.Members {
			var userId = responseMember.UserId

			memberList.Members[userId] = Member{
				UserId:         userId,
				JoinedChatDate: responseMember.JoinedChatDate,
			}
		}

		if len(responseMemberList.Members) < LimitMemberCount {
			delete(searchTextMapList, searchText)
		} else {
			getMemberListBySearchText(clientStorage, memberList, channelId, nestedSearchText)
		}
	}
}

func queryMemberList(clientStorage *client.Storage, channelId int32, searchText string) *ResponseMemberList {
	result, _ := clientStorage.SendAndCatch(client.Request{
		"@type":         "getSupergroupMembers",
		"supergroup_id": channelId,
		"filter": struct {
			Type  string `json:"@type"`
			Query string `json:"query"`
		}{Type: "supergroupMembersFilterSearch", Query: searchText},
		"offset": int32(0),
		"limit":  int32(LimitMemberCount),
	})

	var responseMemberList ResponseMemberList
	_ = json.Unmarshal(result.Raw, &responseMemberList)

	return &responseMemberList
}

func getSearchTextMapList() SearchTextMapList {
	var letters = []string{
		"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z", "1", "2",
		"3", "4", "5", "6", "7", "8", "9",
		"0",
	}
	var length = SearchTextLength

	searchTextMapList := make(SearchTextMapList)

	searchTextGenerator(searchTextMapList, letters, length, "")

	return searchTextMapList
}

func searchTextGenerator(searchTextMapList SearchTextMapList, letters []string, length int, prefix string) {
	var nestedSearchTextMapList = make(SearchTextMapList)

	searchTextMapList[prefix] = nestedSearchTextMapList

	if len(prefix) == length {
		searchTextMapList[prefix] = make(SearchTextMapList)

		return
	}

	for i := 0; i < len(letters); i++ {
		searchTextGenerator(nestedSearchTextMapList, letters, length, prefix+letters[i])
	}
}
