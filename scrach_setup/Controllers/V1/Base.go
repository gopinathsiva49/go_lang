package V1

import (
	"fmt"
	"scrach_setup/Config"
	"scrach_setup/Models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// abort() - Abort prevents pending handlers from being called.
// Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers for this request are not called.
// ref - https://pkg.go.dev/github.com/gin-gonic/gin#Context.Abort

func Auth(c *gin.Context) {
	var session Models.Session

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(400, gin.H{"error": "Missing header 'Authorization'"})
		c.Abort() // not allow next function
		return
	}

	if err := Config.DB.Where("auth_token = ?", tokenString).First(&session).Error; err != nil {
		fmt.Println(err)
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
