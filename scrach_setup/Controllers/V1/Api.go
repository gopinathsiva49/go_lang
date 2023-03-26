package V1

import (
	"html"
	"net/http"
	"regexp"
	"scrach_setup/Config"
	"scrach_setup/Models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetSession(c *gin.Context) {
	session, err := Models.NewSession()
	if err != nil {
		c.JSON(422, err)
	} else {
		c.JSON(200, gin.H{"auth_token": session.AuthToken})
	}
}

func Register(c *gin.Context) {
	var input RegisterInput
	emailre := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	c.ShouldBindJSON(&input)
	if input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is missing"})
		return
	}
	if !emailre.MatchString(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password parameter is missing"})
		return
	}

	if input.ConfirmPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "confirm_password parameter is missing"})
		return
	}

	if input.ConfirmPassword != input.Password {
		c.JSON(403, gin.H{"error": "password is mismatch"})
		return
	}

	//turn password into hash
	hashedPassword, berr := bcrypt.GenerateFromPassword([]byte(input.ConfirmPassword), bcrypt.DefaultCost)
	if berr != nil {
		c.JSON(403, gin.H{"error": "password issue"})
		return
	}

	u := Models.User{}
	u.Email = html.EscapeString(strings.TrimSpace(input.Email))
	u.Password = string(hashedPassword)

	if err := Config.DB.Save(&u).Error; err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func Login(c *gin.Context) {
	var input LoginInput
	c.ShouldBindJSON(&input)
	emailre := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	c.ShouldBindJSON(&input)
	if input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is missing"})
		return
	}
	if !emailre.MatchString(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password parameter is missing"})
		return
	}
	u := Models.User{}
	err := Config.DB.Model(&u).Where("email = ?", input.Email).Take(&u).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "email is not found"})
		return
	}

	berr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))

	if berr != nil && berr == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(403, gin.H{"error": "invalid password"})
		return
	}

	session.UserId = u.ID
	session.ExpiresAt = time.Now().Add(24 * time.Hour)
	Config.DB.Save(&session)
	c.JSON(403, gin.H{"message": "login successfully"})

}
