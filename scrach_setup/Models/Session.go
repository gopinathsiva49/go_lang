package Models

import (
	"crypto/rand"
	"encoding/base64"
	"scrach_setup/Config"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

type Session struct {
	ID         int       `json:"id,primary_key"`
	UUID       string    `json:"uuid"`
	UserId     int       `json:"user_id" gorm:"index"`
	Salt       string    `json:"salt" gorm:"index"`
	SessionKey string    `json:"session_key" gorm:"index"`
	AuthToken  string    `json:"auth_token" gorm:"size:65536"` // INDEX IS PENDING
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewSession() (Error Session, err error) {
	se := Session{}
	se.Salt = SecureRandom(40)
	se.SessionKey = SecureRandom(60)
	claims := jwt.MapClaims{
		"session_key": se.SessionKey,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(se.Salt))
	if err != nil {
		return se, err
	}
	se.AuthToken = tokenString
	se.ExpiresAt = time.Now().Add(24 * time.Hour)

	if err = Config.DB.Save(&se).Error; err != nil {
		return se, err
	}
	return se, err
}

func SecureRandom(val int) string {
	// Generate 40 random bytes
	randomBytes := make([]byte, val)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Encode the random bytes as base64
	base64String := base64.StdEncoding.EncodeToString(randomBytes)

	return base64String

}
