package getStorageChannelList

import (
	"database/sql"
	"github.com/revilon1991/tg-parser/internal/entity"
	"log"
)

func getChannelList(conn *sql.DB) ([]*entity.Channel, error) {
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
		return nil, err
	}

	defer func() {
		err = rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	var storedChannelList []*entity.Channel

	for rows.Next() {
		storedChannel := new(entity.Channel)
		err := rows.Scan(
			&storedChannel.Id,
			&storedChannel.ChannelId,
			&storedChannel.Username,
			&storedChannel.Description,
			&storedChannel.MemberCount,
			&storedChannel.CreatedAt,
			&storedChannel.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		storedChannelList = append(storedChannelList, storedChannel)
	}

	return storedChannelList, nil
}
