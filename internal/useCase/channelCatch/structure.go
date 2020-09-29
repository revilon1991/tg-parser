package channelCatch

type ResponseUpdateSupergroup struct {
    Type       string `json:"@type"`
    Supergroup struct {
        Id                int32  `json:"id"`
        Type              string `json:"@type"`
        Username          string `json:"username"`
        Date              int32  `json:"date"`
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
    } `json:"supergroup"`
}

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
type Channel struct {
    Id          int32
    Description string
    Username    string
}
