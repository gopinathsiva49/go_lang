package Controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Alive struct {
	CurrentTime time.Time `json:"current_time"`
}

func GetAlive(c *gin.Context) {
	a := Alive{CurrentTime: time.Now()}
	c.JSON(http.StatusOK, a)
	// c.Abort() // api end
}
