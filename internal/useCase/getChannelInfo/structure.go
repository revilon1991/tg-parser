package getChannelInfo

type ResponseSupergroupFullInfo struct {
	Description        string `json:"description"`
	MemberCount        int32  `json:"member_count"`
	AdministratorCount int32  `json:"administrator_count"`
	RestrictedCount    int32  `json:"restricted_count"`
	BannedCount        int32  `json:"banned_count"`
	LinkedChatId       int64  `json:"linked_chat_id"`
	CanGetMembers      bool   `json:"can_get_members"`
	CanSetUsername     bool   `json:"can_set_username"`
	CanSetStickerSet   bool   `json:"can_set_sticker_set"`
	CanSetLocation     bool   `json:"can_set_location"`
	CanViewStatistics  bool   `json:"can_view_statistics"`
	StickerSetId       int64  `json:"sticker_set_id"`
}
