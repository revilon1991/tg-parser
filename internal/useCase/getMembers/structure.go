package getMembers

type MemberList struct {
	Total   int            `json:"total"`
	Members map[int]Member `json:"members"`
}

type Member struct {
	JoinedChatDate int `json:"joined_chat_date"`
	UserId         int `json:"user_id"`
}

type SearchTextMapList map[string]SearchTextMapList

type ResponseMemberList struct {
	Type    string `json:"@type"`
	Members []struct {
		Type           string `json:"@type"`
		JoinedChatDate int    `json:"joined_chat_date"`
		UserId         int    `json:"user_id"`
	} `json:"members"`
}
