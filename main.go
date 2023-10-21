package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgres://postgres:@localhost:5432/postgres?sslmode=disable")))
	result := map[string]interface{}{}
	db := bun.NewDB(sqldb, pgdialect.New())

	err := db.NewRaw("SELECT VERSION()").Scan(context.Background(), &result)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("success", result)
}
