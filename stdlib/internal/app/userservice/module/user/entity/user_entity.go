package entity

type User struct {
	Id        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	CreatedAt int64  `db:"created_at" json:"createdAt"`
	UpdatedAt int64  `db:"updated_at" json:"updatedAt"`
}
