## A RSS feed Aggrigator


Initalize the goproject
```bash
go mod init
```

install SQLC and GOOSE
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

go get  github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
Goose is help us in db migratition

SQLC enable us to generate codes related to our sql db queries

Goose Migration command 
```bash
goose postgres db_connection_string up
# example :=  goose postgres postgres://postgres:root@locahost:5432 up
```

