package main

import (
	"fmt"

	"github.com/aoyako/swolf"
)

func main() {
	// fmt.Println("Hello, world!")
	d := swolf.Setup(swolf.Config{
		Driver: "postgres",
		Connection: "host=localhost user=postgres " +
			"password=admin sslmode=disable",
		MasterDatabase: "MasterDatabase",
		MasterTable:    "main_table",
		MasterData:     swolf.MasterFieldResolver("tenant_id", "tenant_db"),
	})

	value, err := d.Create("poggers")
	fmt.Println(value, err)
}
