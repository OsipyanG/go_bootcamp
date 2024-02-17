package db_reader

import (
	"encoding/json"
	"os"
)

type JsonReader struct {
	DB DataBase
}

func (j *JsonReader) Read(filename string) (DataBase, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return DataBase{}, err
	}
	err = json.Unmarshal(data, &j.DB)
	if err != nil {
		return DataBase{}, err
	}
	return j.DB, nil
}
