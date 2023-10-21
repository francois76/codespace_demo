package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type utilisateur struct {
	bun.BaseModel `bun:"table:utilisateur,alias:u"`
	Age           int    `bun:"age"`
	Id            int    `bun:"id,pk,autoincrement"`
	Nom           string `bun:"nom"`
	Prenom        string `bun:"prenom"`
}

func main() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgres://postgres:@localhost:5432/postgres?sslmode=disable")))
	result := []utilisateur{}
	db := bun.NewDB(sqldb, pgdialect.New())

	err := db.NewRaw("SELECT * from utilisateurs").Scan(context.Background(), &result)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("success", result)
}
