package Migrations

import (
	"scrach_setup/Config"
	"scrach_setup/Models"
)

func UserMigrate() {
	Config.DB.AutoMigrate(Models.User{})
	// Config.DB.Migrator().DropColumn(Models.User{}, "MiddleName") // need to study
}
