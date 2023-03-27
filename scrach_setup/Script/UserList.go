package main

import (
	"fmt"
	"log"
	"scrach_setup/Config"
	"scrach_setup/Migrations"
	"scrach_setup/Models"

	"github.com/joho/godotenv"
)

func userlist() {
	var user []Models.User
	if err := Config.DB.Find(&user).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	////initialize the database
	Config.InitDB()
	///add this line to main.go to initialize the migration
	Migrations.Migrate()

}
func main() {
	userlist()
}
