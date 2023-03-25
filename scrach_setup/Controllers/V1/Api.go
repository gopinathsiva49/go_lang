package V1

import (
	"scrach_setup/Models"

	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) {
	session, err := Models.NewSession()
	if err != nil {
		c.JSON(422, err)
	} else {
		c.JSON(200, gin.H{"auth_token": session.AuthToken})
	}
}
