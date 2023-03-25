package Migrations

// https://gorm.io/docs/migration.html
func Migrate() {
	UserMigrate()
	SessionMigrate()
}
