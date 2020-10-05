package fetchChannelMembersInfo

import (
	"database/sql"
	"fmt"
	"github.com/revilon1991/tg-parser/internal/connection/mysql"
	"github.com/revilon1991/tg-parser/internal/entity"
	"log"
	"strconv"
	"strings"
	"time"

	// Register mysql
	_ "github.com/go-sql-driver/mysql"
)

func getChannelList() []*entity.Channel {
	conn := mysql.Open()

	defer mysql.Close(conn)

	query := `
        select
            c.id,
            c.channelId,
            c.username,
            c.description,
            c.memberCount,
            c.createdAt,
            c.updatedAt
        from Channel c
    `

	rows, err := conn.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	channelList := make([]*entity.Channel, 0)

	for rows.Next() {
		channel := new(entity.Channel)
		err := rows.Scan(
			&channel.Id,
			&channel.ChannelId,
			&channel.Username,
			&channel.Description,
			&channel.MemberCount,
			&channel.CreatedAt,
			&channel.UpdatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}

		channelList = append(channelList, channel)
	}

	return channelList
}

func saveMemberPhotoList(memberInfoList *MemberInfoList, conn *sql.DB) error {
	var err error
	var queryDeleteValues []string
	var queryDeleteArgs []interface{}
	var queryInsertValues []string
	var queryInsertArgs []interface{}

	now := time.Now()
	nowString := now.Format("2006-01-02 15:04:05")

	for _, member := range memberInfoList.MemberList {
		queryDeleteValues = append(queryDeleteValues, "?")
		queryDeleteArgs = append(queryDeleteArgs, strconv.Itoa(member.UserId))

		for _, photo := range member.PhotoList {
			queryInsertValues = append(queryInsertValues, "(?, ?, ?, ?)")
			queryInsertArgs = append(queryInsertArgs, strconv.Itoa(member.UserId), photo, nowString, nowString)
		}
	}

	if len(queryInsertValues) == 0 {
		return err
	}

	queryDeletePattern := `
        delete from MemberPhoto where memberId in (%s)
    `

	queryDelete := fmt.Sprintf(queryDeletePattern, strings.Join(queryDeleteValues, ","))

	_, err = conn.Exec(queryDelete, queryDeleteArgs...)

	if err != nil {
		return err
	}

	queryInsertPattern := `
        insert into MemberPhoto (memberId, link, createdAt, updatedAt) values %s 
    `

	queryInsert := fmt.Sprintf(queryInsertPattern, strings.Join(queryInsertValues, ","))

	_, err = conn.Exec(queryInsert, queryInsertArgs...)

	return err
}

func saveMembers(memberInfoList *MemberInfoList, conn *sql.DB) error {
	var err error

	now := time.Now()
	nowString := now.Format("2006-01-02 15:04:05")

	var insertValues []string
	var insertArgs []interface{}

	for _, member := range memberInfoList.MemberList {
		insertValues = append(insertValues, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		insertArgs = append(
			insertArgs,
			strconv.Itoa(member.UserExternalId),
			member.Username,
			member.FirstName,
			member.LastName,
			member.PhoneNumber,
			member.Type,
			member.Bio,
			nowString,
			nowString,
		)
	}

	query := fmt.Sprintf(`
		insert into Member (userId, username, firstName, lastName, phoneNumber, type, bio, createdAt, updatedAt) values %s
		on duplicate key update
			username=values(username),
			firstName=values(firstName),
			lastName=values(lastName),
			phoneNumber=values(phoneNumber),
			type=values(type),
			bio=values(bio)
		`,
		strings.Join(insertValues, ","),
	)

	_, err = conn.Exec(query, insertArgs...)

	return err
}

func fetchMemberIdList(memberInfoList *MemberInfoList, conn *sql.DB) error {
	var err error
	var sqlPlaceholders []string
	var sqlArguments []interface{}

	for _, member := range memberInfoList.MemberList {
		sqlPlaceholders = append(sqlPlaceholders, "?")
		sqlArguments = append(sqlArguments, strconv.Itoa(member.UserExternalId))
	}

	queryPattern := `
		select
			m.id,
			m.userId
		from Member m
		where 1
			and userId in (%s)
	`
	query := fmt.Sprintf(queryPattern, strings.Join(sqlPlaceholders, ","))

	rows, err := conn.Query(query, sqlArguments...)

	if err != nil {
		return err
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		storedMember := new(entity.Member)
		err := rows.Scan(
			&storedMember.Id,
			&storedMember.UserId,
		)

		if err != nil {
			return err
		}

		member := memberInfoList.MemberList[storedMember.UserId]
		member.UserId = storedMember.Id

		memberInfoList.MemberList[storedMember.UserId] = member
	}

	return err
}

func saveRelationChannelMember(memberInfoList *MemberInfoList, conn *sql.DB) error {
	var err error

	queryDelete := `
		delete from ChannelHasMember where channelId = ?
	`

	_, err = conn.Exec(queryDelete, memberInfoList.ChannelId)

	if err != nil {
		return err
	}

	var sqlInsertPlaceholders []string
	var sqlInsertArguments []interface{}

	now := time.Now()
	nowString := now.Format("2006-01-02 15:04:05")

	for _, member := range memberInfoList.MemberList {
		sqlInsertPlaceholders = append(sqlInsertPlaceholders, "(?, ?, ?, ?, ?)")
		sqlInsertArguments = append(
			sqlInsertArguments,
			strconv.Itoa(memberInfoList.ChannelId),
			strconv.Itoa(member.UserId),
			member.JoinChannelDate,
			nowString,
			nowString,
		)
	}

	queryInsert := fmt.Sprintf(
		"insert into ChannelHasMember (channelId, memberId, joinDate, createdAt, updatedAt) values %s",
		strings.Join(sqlInsertPlaceholders, ","),
	)

	_, err = conn.Exec(queryInsert, sqlInsertArguments...)

	return err
}
