package schema

import "time"

type User struct {
	Id        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name,omitempty" json:"name,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
