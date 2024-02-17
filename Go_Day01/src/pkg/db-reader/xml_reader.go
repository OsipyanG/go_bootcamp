package db_reader

import (
	"encoding/xml"
	"os"
)

type XmlReader struct {
	DB DataBase
}

func (x *XmlReader) Read(filename string) (DataBase, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return DataBase{}, err
	}
	err = xml.Unmarshal(data, &x.DB)
	if err != nil {
		return DataBase{}, err
	}
	return x.DB, nil
}
