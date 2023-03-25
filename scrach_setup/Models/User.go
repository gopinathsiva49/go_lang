package Models

import "time"

type User struct {
	ID        int       `json:"id,primary_key"` //this will be the primary key field
	FirstName string    `json:"first_name" gorm:"index"`
	LastName  string    `json:"last_name" gorm:"index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
