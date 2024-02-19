package main

import (
	dbreader "db-reader/pkg/db-reader"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

var ErrNoArguments = errors.New("there are no input arguments")

type DBReader interface {
	Read(filename string) (dbreader.DataBase, error)
}

type CliFlags struct {
	file bool
}

func main() {
	cliFlags := CliFlags{file: false}
	err := cliFlags.parseFlags()
	var dbReader DBReader
	if err != nil {
		log.Fatalln(err.Error())
	}

	if cliFlags.file {
		filename := flag.Arg(0)
		if filepath.Ext(filename) == ".json" {
			dbReader = &dbreader.JsonReader{}
			db, err := dbReader.Read("../stolen_database.json")
			if err != nil {
				log.Fatalln(err.Error())
			}
			xmlDB, err := xml.MarshalIndent(db, "", "    ")
			if err != nil {
				log.Fatalln("Marshaling fail")
			}
			fmt.Println(string(xmlDB))
		} else if filepath.Ext(filename) == ".xml" {
			dbReader = &dbreader.XmlReader{}
			db, err := dbReader.Read("../original_database.xml")
			if err != nil {
				log.Fatalln(err.Error())
			}
			jsonDB, err := json.MarshalIndent(db, "", "    ")
			if err != nil {
				log.Fatalln("Marshaling fail")
			}
			fmt.Println(string(jsonDB))
		} else {
			log.Fatalln("File format error")
		}
	} else {
		log.Fatalln(ErrNoArguments.Error())
	}
}

func (f *CliFlags) parseFlags() error {
	flag.BoolVar(&f.file, "f", false, "Read the file and output it in a different format")
	flag.Parse()
	if len(flag.Args()) < 1 {
		return ErrNoArguments
	}
	return nil
}
