package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type DBReader interface {
	Read()
}

func main() {

	content, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
