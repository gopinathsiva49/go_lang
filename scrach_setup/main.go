package main

import (
	"scrach_setup/Config"
	"scrach_setup/Migrations"

	"scrach_setup/Routes"
)

func init() {
	////initialize the database
	Config.InitDB()
	///add this line to main.go to initialize the migration
	Migrations.Migrate()
}
func main() {

	//setup routes
	r := Routes.SetupRouter()

	// running
	r.Run(":8080")
}
