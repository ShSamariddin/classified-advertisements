go install github.com/pressly/goose/v3/cmd/goose@latest

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="host=localhost port=6432 user=advertisements password=advertisements dbname=advertisements sslmode=disable"

cd migrations
goose up
