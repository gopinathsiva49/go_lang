package Models

import "time"

type User struct {
	ID        int       `json:"id,primary_key"` //this will be the primary key field
	Email     string    `json:"email" gorm:"size:255;not null;unique"`
	Password  string    `json:"password" gorm:"size:255;not null;"`
	FirstName string    `json:"first_name" gorm:"index"`
	LastName  string    `json:"last_name" gorm:"index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
