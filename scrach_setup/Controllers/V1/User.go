package V1

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"scrach_setup/Config"
	"scrach_setup/Models"
)

func ListUsers(c *gin.Context) {
	var user []Models.User
	if err := Config.DB.Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func CreateUsers(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	if err := Config.DB.Create(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func UpdateUsers(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")

	if err := Config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		fmt.Println(err)
	} else {
		c.BindJSON(&user)
		Config.DB.Save(&user)
		c.JSON(200, user)
	}
}

func DeleteUsers(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")

	if err := Config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"message": "deleted successfully"})
	}
}
