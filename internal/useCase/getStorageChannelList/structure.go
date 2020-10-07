package getStorageChannelList

type ResponseChannelList []ResponseChannel

type ResponseChannel struct {
	Id          int    `json:"id"`
	OuterId     int    `json:"channel_id"`
	Username    string `json:"username"`
	Description string `json:"description"`
	MemberCount int    `json:"member_count"`
	UpdatedAt   string `json:"updated_at"`
}
