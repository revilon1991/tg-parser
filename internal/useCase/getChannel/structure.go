package getChannel

type ResponseSupergroup struct {
	Id                int32  `json:"id"`
	Type              string `json:"@type"`
	Username          string `json:"username"`
	Date              int32  `json:"date"`
	MemberCount       int32  `json:"member_count"`
	HasLinkedChat     bool   `json:"has_linked_chat"`
	HasLocation       bool   `json:"has_location"`
	SignMessages      bool   `json:"sign_messages"`
	IsSlowModeEnabled bool   `json:"is_slow_mode_enabled"`
	IsChannel         bool   `json:"is_channel"`
	IsVerified        bool   `json:"is_verified"`
	RestrictionReason string `json:"restriction_reason"`
	IsScam            bool   `json:"is_scam"`
	Status            struct {
		Type               string `json:"@type"`
		CanChangeInfo      bool   `json:"can_change_info"`
		CanEditMessages    bool   `json:"can_edit_messages"`
		CanDeleteMessages  bool   `json:"can_delete_messages"`
		CustomTitle        string `json:"custom_title"`
		CanInviteUsers     bool   `json:"can_invite_users"`
		CanRestrictMembers bool   `json:"can_restrict_members"`
		CanPinMessages     bool   `json:"can_pin_messages"`
		CanPromoteMembers  bool   `json:"can_promote_members"`
		CanBeEdited        bool   `json:"can_be_edited"`
		CanPostMessages    bool   `json:"can_post_messages"`
	} `json:"status"`
}
