package entity

import "database/sql"

type Channel struct {
	Id          int            `db:"id"`
	ChannelId   int            `db:"channelId"`
	Username    sql.NullString `db:"username"`
	Description sql.NullString `db:"description"`
	MemberCount int            `db:"memberCount"`
	CreatedAt   string         `db:"createdAt"`
	UpdatedAt   string         `db:"updatedAt"`
}
