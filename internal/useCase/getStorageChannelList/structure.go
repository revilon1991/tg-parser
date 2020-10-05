package getStorageChannelList

type ChannelList []Channel

type Channel struct {
	Id          int    `json:"id"`
	ChannelId   int    `json:"channel_id"`
	Username    string `json:"username"`
	Description string `json:"description"`
	MemberCount int    `json:"member_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
