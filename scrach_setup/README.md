Things you may want to cover:

- go mod init scrach_setup
- go get .

PACKAGES

- go get github.com/gin-gonic/gin (gin framwork)
- go get -u gorm.io/gorm (orm)
- go get gorm.io/driver/mysql
- go get -u github.com/dgrijalva/jwt-go
- golang.org/x/crypto/bcrypt

ENV (update the enviroment value in .env)

RUN

- go run .

API

- / (alive)
- /v1/session
- /v1/register
- /v1/login
- /v1/users (GET,POST,PUT,DELETE)
