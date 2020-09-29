package fetchChannelMembersInfo

type MemberInfoList struct {
    MemberList        map[int]Member
    ChannelExternalId int
    ChannelId         int
}

type Member struct {
    UserId          int
    ChannelId       int
    UserExternalId  int
    JoinChannelDate string
    Username        string
    FirstName       string
    LastName        string
    PhoneNumber     string
    Type            string
    Bio             string
}
