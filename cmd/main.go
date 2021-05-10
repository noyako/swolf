package main

import (
	"fmt"

	"github.com/aoyako/swolf"
)

func main() {
	d := swolf.Setup(swolf.Config{
		Driver: "postgres",
		Connection: "host=localhost user=postgres " +
			"password=admin sslmode=disable",
		MasterDatabase: "MasterDatabase",
		MasterTable:    "main_table",
		MasterData:     swolf.MasterFieldResolver("tenant_id", "tenant_db"),
	})

	db, err := d.Create("poggers1")
	fmt.Println(db, err)
}
