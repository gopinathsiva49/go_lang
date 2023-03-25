go mod init scrach_setup

PACKAGES
go get github.com/gin-gonic/gin (gin framwork)
go get -u gorm.io/gorm (orm)
go get gorm.io/driver/mysql

ENV
export HOST=localhost
export USER=root
export DBNAME=todo
export PASSWORD=
