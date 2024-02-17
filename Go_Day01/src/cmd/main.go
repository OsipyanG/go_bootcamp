package main

import (
	dbreader "db-reader/pkg/db-reader"
	"fmt"
	"log"
)

type DBReader interface {
	Read(filename string) (dbreader.DataBase, error)
}

func main() {
	var dbReader DBReader = &dbreader.JsonReader{}
	db, err := dbReader.Read("../stolen_database.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(db)
}
