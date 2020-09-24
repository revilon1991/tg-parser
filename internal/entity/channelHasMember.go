package entity

type ChannelHasMember struct {
	Id        int    `db:"id"`
	channelId int    `db:"channelId"`
	memberId  int    `db:"memberId"`
	CreatedAt string `db:"createdAt"`
	UpdatedAt string `db:"updatedAt"`
}
