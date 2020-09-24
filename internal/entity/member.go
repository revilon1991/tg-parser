package entity

type Member struct {
	Id        int    `db:"id"`
	UserId    int    `db:"userId"`
	CreatedAt string `db:"createdAt"`
	UpdatedAt string `db:"updatedAt"`
}
