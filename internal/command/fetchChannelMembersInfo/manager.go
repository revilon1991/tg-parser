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

	sql := `
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

	rows, err := conn.Query(sql)

	if err != nil {
		log.Fatal("Consumer channelCatch. saveChannel error: " + err.Error())
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

func saveMembers(memberList getMembers.MemberList) {
	conn := mysql.Open()

	defer mysql.Close(conn)

	now := time.Now()
	nowString := now.Format("2006-01-02 15:04:05")

	var insertValues []string
	var insertArgs []interface{}

	for _, member := range memberList.Members {
		insertValues = append(insertValues, "(?, ?, ?)")
		insertArgs = append(insertArgs, strconv.Itoa(member.UserId), nowString, nowString)
	}

	sql := fmt.Sprintf(
		"insert ignore into Member (userId, createdAt, updatedAt) values %s",
		strings.Join(insertValues, ","),
	)

	_, err := conn.Exec(sql, insertArgs...)

	if err != nil {
		log.Fatal("command fetchChannelMembersInfo. saveMembers error: " + err.Error())
	}
}
