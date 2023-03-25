package Migrations

import (
	"scrach_setup/Config"
	"scrach_setup/Models"
)

func SessionMigrate() {
	Config.DB.AutoMigrate(Models.Session{})
}
