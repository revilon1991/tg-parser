package getStorageMemberList

import (
	"github.com/revilon1991/tg-parser/internal/connection/mysql"
)

func Handle(channelId int) ([]*ResponseMember, error) {
	conn := mysql.Open()

	defer mysql.Close(conn)

	responseMemberList, err := getMemberList(conn, channelId)
	if err != nil {
		return nil, err
	}

	membersPhotoList, err := getMembersPhotoList(conn, channelId)
	if err != nil {
		return nil, err
	}

	for _, responseMember := range responseMemberList {
		if memberPhoto, ok := membersPhotoList[responseMember.Id]; ok {
			responseMember.PhotoList = append(responseMember.PhotoList, memberPhoto.Link)
		}
	}

	return responseMemberList, nil
}
