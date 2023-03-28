Things you may want to cover:

- go mod init scrach_setup
- go get .

PACKAGES

- (mac os) brew install golang-migrate
- (Linux) curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz

- go get github.com/gin-gonic/gin (gin framwork)
- go get -u gorm.io/gorm (orm)
- go get gorm.io/driver/mysql
- go get -u github.com/dgrijalva/jwt-go
- golang.org/x/crypto/bcrypt

ENV (update the enviroment value in .env)

CMD

- make db-create
- make db-drop
- make db-generate-migration file_name=[filename]
- make db-migrate-all
- make db-rollback-all
- make db-migrate-one-by-one
- make db-rollback-one-by-one
- make run

API

- / (alive)
- /v1/session
- /v1/register
- /v1/login
- /v1/users (GET,POST,PUT,DELETE)
