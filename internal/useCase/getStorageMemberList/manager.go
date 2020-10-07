package getStorageMemberList

import (
	"database/sql"
	"log"
)

func getMemberList(conn *sql.DB, channel int) ([]*ResponseMember, error) {
	query := `
        select
            m.id,
            m.userId,
            m.username,
            m.firstName,
            m.lastName,
            m.phoneNumber,
            m.type,
            m.bio,
            m.updatedAt,
            chm.joinDate
        from ChannelHasMember chm
        inner join Member m on m.id = chm.memberId
        where 1
            and chm.channelId = ?
        order by chm.joinDate desc
    `

	rows, err := conn.Query(query, channel)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	responseMemberList := make([]*ResponseMember, 0)

	for rows.Next() {
		responseMember := &ResponseMember{
			PhotoList: []string{},
		}
		err := rows.Scan(
			&responseMember.Id,
			&responseMember.UserId,
			&responseMember.Username,
			&responseMember.FirstName,
			&responseMember.LastName,
			&responseMember.PhoneNumber,
			&responseMember.Type,
			&responseMember.Bio,
			&responseMember.UpdatedAt,
			&responseMember.JoinDate,
		)

		if err != nil {
			return nil, err
		}

		responseMemberList = append(responseMemberList, responseMember)
	}

	return responseMemberList, nil
}

func getMembersPhotoList(conn *sql.DB, channel int) (map[int]MemberPhoto, error) {
	query := `
        select
            mp.memberId,
            mp.link
        from MemberPhoto mp
        inner join ChannelHasMember chm on chm.memberId = mp.memberId
        where 1
            and chm.channelId = ?
    `

	rows, err := conn.Query(query, channel)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	memberPhotoList := make(map[int]MemberPhoto, 0)

	for rows.Next() {
		responseMember := MemberPhoto{}
		err := rows.Scan(
			&responseMember.MemberId,
			&responseMember.Link,
		)

		if err != nil {
			return nil, err
		}

		memberPhotoList[responseMember.MemberId] = responseMember
	}

	return memberPhotoList, nil
}
