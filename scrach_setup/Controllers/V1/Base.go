package V1

import (
	"scrach_setup/Config"
	"scrach_setup/Models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// abort() - Abort prevents pending handlers from being called.
// Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers for this request are not called.
// ref - https://pkg.go.dev/github.com/gin-gonic/gin#Context.Abort
var session Models.Session

func Auth(c *gin.Context) {
	session = Models.Session{} // refresh
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(400, gin.H{"error": "Missing header 'Authorization'"})
		c.Abort() // not allow next function
		return
	}

	if err := Config.DB.Where("auth_token = ?", tokenString).First(&session).Error; err != nil {
		c.JSON(404, gin.H{"error": "No active session found"})
		c.Abort()
		return
	}

	claims := jwt.MapClaims{}
	_, terr := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(session.Salt), nil
	})

	if terr != nil {
		c.JSON(403, gin.H{"error": "Token Verification Error"})
		c.Abort()
		return

	}

	for key, val := range claims {
		if key == "session_key" {
			if val != session.SessionKey {
				c.JSON(403, gin.H{"error": "Token mismatch"})
				c.Abort()
				return
			}
		}
	}
	c.Next() // allow next function
}

type CompareTime struct {
	CurrentTime time.Time
}

func AccountAuth(c *gin.Context) {
	if session.UserId == 0 {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	if session.ExpiresAt.Unix() < time.Now().Unix() {
		c.JSON(403, gin.H{"error": "Token expired"})
		c.Abort()
		return
	}
}
