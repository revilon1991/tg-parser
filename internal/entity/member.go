package entity

import "database/sql"

type Member struct {
    Id          int            `db:"id"`
    UserId      int            `db:"userId"`
    Username    sql.NullString `db:"username"`
    FirstName   sql.NullString `db:"first_name"`
    LastName    sql.NullString `db:"last_name"`
    PhoneNumber sql.NullString `db:"phone_number"`
    Bio         sql.NullString `db:"bio"`
    Type        string         `db:"type"`
    CreatedAt   string         `db:"createdAt"`
    UpdatedAt   string         `db:"updatedAt"`
}
