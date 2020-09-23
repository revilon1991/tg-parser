package channelCatch

import (
	"github.com/revilon1991/tg-parser/internal/connection/mysql"
	"log"
	"time"

	// Register mysql
	_ "github.com/go-sql-driver/mysql"
)

func saveChannel(channelId int32, username string, memberCount int32, description string) {
	conn := mysql.Open()

	defer mysql.Close(conn)

	sql := `
       insert into Channel (channelId, username, memberCount, description, createdAt, updatedAt) values (?, ?, ?, ?, ?, ?)
       on duplicate key update
           username=if(values(username) = '', username, values(username)),
           memberCount=if(values(memberCount) > 0, memberCount, values(memberCount)),
           description=if(values(description) = '', description, values(description)),
           updatedAt=values(updatedAt)
    `

	now := time.Now()
	nowString := now.Format("2006-01-02 15:04:05")

	args := []interface{}{channelId, username, memberCount, description, nowString, nowString}
	_, err := conn.Exec(sql, args...)

	if err != nil {
		log.Fatal("Consumer channelCatch. saveChannel error: " + err.Error())
	}
}
