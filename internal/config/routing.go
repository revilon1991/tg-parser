package config

var Routing = RoutingStruct{
	V1GetMe:          "/v1/getMe",
	V1GetMembers:     "/v1/getMembers",
	V1GetUser:        "/v1/getUser",
	V1GetPhoto:       "/v1/getPhoto",
	V1GetChannelInfo: "/v1/getChannelInfo",
	V1GetChannel:     "/v1/getChannel",

	V1StorageGetChannelList: "/v1/storage/getChannelList",
	V1StorageGetMemberList:  "/v1/storage/getMemberList",
}

type RoutingStruct struct {
	V1GetMe          string
	V1GetMembers     string
	V1GetUser        string
	V1GetPhoto       string
	V1GetChannelInfo string
	V1GetChannel     string

	V1StorageGetChannelList string
	V1StorageGetMemberList  string
}
