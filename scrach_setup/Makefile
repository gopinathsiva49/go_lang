include .env


MYSQL_DSN ?= $(DBUSER):$(DBPASSWORD)@tcp($(DBHOST):3306)/$(DBNAME)

db-create:
	mysql -h $(DBHOST) -u $(DBUSER) -p$(DBPASSWORD) -e "CREATE DATABASE $(DBNAME)"

db-drop:
	mysql -h $(DBHOST) -u $(DBUSER) -p$(DBPASSWORD) -e "DROP DATABASE $(DBNAME)"

db-generate-migration:
	migrate create -ext sql -dir db/migrations -seq $(file_name) 

db-migrate-all:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -verbose -path db/migrations up

db-rollback-all: 
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -verbose  -path db/migrations down

db-migrate-one-by-one:
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -verbose -path db/migrations up 1

db-rollback-one-by-one: 
	migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -verbose  -path db/migrations down 1

run:
	go run .

	



