package entity

import "time"

type UserSchema struct {
	Id        string    `bson:"_id," json:"id"`
	Name      string    `bson:"name,omitempty" json:"name"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}
