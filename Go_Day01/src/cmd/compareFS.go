package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type cliFlags struct {
	old string
	new string
}

var ErrNoArguments = errors.New("required arguments are not entered, use --help to view the arguments")

func main() {
	flags := cliFlags{}
	flags.initFlags()
	if flags.new == "" || flags.old == "" {
		log.Fatalln(ErrNoArguments.Error())
	}
	oldMap := getMapOfFile(flags.old)
	newMap := getMapOfFile(flags.new)
	CompareFiles(oldMap, newMap)
}

func CompareFiles(oldMap, newMap map[string]struct{}) {
	for file, _ := range newMap {
		_, ok := oldMap[file]
		if !ok {
			fmt.Printf("ADDED %s \n", file)
		}
	}
	for file, _ := range oldMap {
		_, ok := newMap[file]
		if !ok {
			fmt.Printf("REMOVED %s \n", file)
		}
	}
}

func getMapOfFile(filename string) map[string]struct{} {
	fileMap := make(map[string]struct{})
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileMap[scanner.Text()] = struct{}{}
	}
	return fileMap
}

func (f *cliFlags) initFlags() {
	flag.StringVar(&f.old, "old", "", "the file with the old recipe")
	flag.StringVar(&f.new, "new", "", "the file with the new recipe")
	flag.Parse()
}
