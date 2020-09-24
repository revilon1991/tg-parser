package fetchChannelMembersInfo

import (
	"fmt"
	"github.com/revilon1991/tg-parser/internal/connection/mysql"
	"github.com/revilon1991/tg-parser/internal/entity"
	"github.com/revilon1991/tg-parser/internal/useCase/getMembers"
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

func saveMembers(responseMemberList getMembers.MemberList) {
	conn := mysql.Open()

	defer mysql.Close(conn)

	now := time.Now()
	nowString := now.Format("2006-01-02 15:04:05")

	var insertValues []string
	var insertArgs []interface{}

	for _, member := range responseMemberList.Members {
		insertValues = append(insertValues, "(?, ?, ?)")
		insertArgs = append(insertArgs, strconv.Itoa(member.UserId), nowString, nowString)
	}

	query := fmt.Sprintf(
		"insert ignore into Member (userId, createdAt, updatedAt) values %s",
		strings.Join(insertValues, ","),
	)

	_, err := conn.Exec(query, insertArgs...)

	if err != nil {
		log.Fatal(err)
	}
}

func getMemberIdList(responseMemberList getMembers.MemberList) []*entity.Member {
	conn := mysql.Open()

	defer mysql.Close(conn)

	var sqlPlaceholders []string
	var sqlArguments []interface{}

	for _, member := range responseMemberList.Members {
		sqlPlaceholders = append(sqlPlaceholders, "?")
		sqlArguments = append(sqlArguments, strconv.Itoa(member.UserId))
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
		log.Fatal(err)
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	memberList := make([]*entity.Member, 0)

	for rows.Next() {
		member := new(entity.Member)
		err := rows.Scan(
			&member.Id,
			&member.UserId,
		)

		if err != nil {
			log.Fatal(err)
		}

		memberList = append(memberList, member)
	}

	return memberList
}

func saveRelationChannelMember(channel *entity.Channel, memberList []*entity.Member) {
	conn := mysql.Open()

	defer mysql.Close(conn)

	_, err := conn.Exec(`begin`)

	if err != nil {
		log.Fatal(err)
	}

	queryDelete := `
		delete from ChannelHasMember where channelId = ?
	`

	_, err = conn.Exec(queryDelete, channel.Id)

	if err != nil {
		_, errRollback := conn.Exec(`rollback`)

		if errRollback != nil {
			log.Fatal(errRollback, err)
		}

		log.Fatal(err)
	}

	var sqlInsertPlaceholders []string
	var sqlInsertArguments []interface{}

	now := time.Now()
	nowString := now.Format("2006-01-02 15:04:05")

	for _, member := range memberList {
		sqlInsertPlaceholders = append(sqlInsertPlaceholders, "(?, ?, ?, ?)")
		sqlInsertArguments = append(
			sqlInsertArguments,
			strconv.Itoa(channel.Id),
			strconv.Itoa(member.Id),
			nowString,
			nowString,
		)
	}

	queryInsert := fmt.Sprintf(
		"insert into ChannelHasMember (channelId, memberId, createdAt, updatedAt) values %s",
		strings.Join(sqlInsertPlaceholders, ","),
	)

	_, err = conn.Exec(queryInsert, sqlInsertArguments...)

	if err != nil {
		_, errRollback := conn.Exec(`rollback`)

		if errRollback != nil {
			log.Fatal(errRollback, err)
		}

		log.Fatal(err)
	}

	_, err = conn.Exec(`commit`)

	if err != nil {
		_, errRollback := conn.Exec(`rollback`)

		if errRollback != nil {
			log.Fatal(errRollback, err)
		}

		log.Fatal(err)
	}
}
