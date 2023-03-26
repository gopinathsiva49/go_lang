package Routes

import (
	"scrach_setup/Controllers"
	"scrach_setup/Controllers/V1"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", Controllers.GetAlive)
	version1 := r.Group("/v1")
	// .Use(Controllers.GetAlive)
	{
		version1.GET("/session", V1.GetSession)
		version1.POST("/register", V1.Auth, V1.Register)
		version1.POST("/login", V1.Auth, V1.Login)
		version1.GET("/users", V1.Auth, V1.AccountAuth, V1.ListUsers)
		version1.POST("/users", V1.Auth, V1.AccountAuth, V1.CreateUsers)
		version1.PUT("/users/:id", V1.Auth, V1.AccountAuth, V1.UpdateUsers)
		version1.DELETE("/users/:id", V1.Auth, V1.AccountAuth, V1.DeleteUsers)

	}
	return r
}
