package main

import (
	"fmt"
	"todo/Config"

	"todo/Models"

	"todo/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Creating a connection to the database
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})

	//setup routes
	r := Routes.SetupRouter()

	// running
	r.Run()
}
