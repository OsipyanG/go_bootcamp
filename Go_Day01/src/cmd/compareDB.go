package main

import (
	dbreader "db-reader/pkg/db-reader"
	"errors"
	"flag"
	"fmt"
	"log"
)

type DBReader interface {
	Read(filename string) (dbreader.DataBase, error)
}

var ErrNoArguments = errors.New("required arguments are not entered, use --help to view the arguments")
var ErrXmlReading = errors.New("error reading the xml database")
var ErrJsonReading = errors.New("error reading the json database")

type Flags struct {
	old string
	new string
}

func main() {
	var oldRead DBReader = &dbreader.XmlReader{}
	var newRead DBReader = &dbreader.JsonReader{}

	flags := Flags{}
	flags.initFlags()
	if flags.new == "" || flags.old == "" {
		log.Fatalln(ErrNoArguments.Error())
	}

	oldDB, err := oldRead.Read(flags.old)
	if err != nil {
		log.Fatalln(ErrXmlReading.Error())
	}
	newDB, err := newRead.Read(flags.new)
	if err != nil {
		log.Fatalln(ErrJsonReading.Error())
	}
	OldMap := CakesToMap(oldDB.Cakes)
	NewMap := CakesToMap(newDB.Cakes)
	CompareDataBases(OldMap, NewMap)
}

func CakesToMap(Cakes []dbreader.Cake) map[string]dbreader.Cake {
	CakesMap := make(map[string]dbreader.Cake, len(Cakes))
	for _, cake := range Cakes {
		CakesMap[cake.Name] = cake
	}
	return CakesMap
}
func IngridientsToMap(Ingredients []dbreader.Ingredient) map[string]dbreader.Ingredient {
	IngredientsMap := make(map[string]dbreader.Ingredient, len(Ingredients))
	for _, ingredient := range Ingredients {
		IngredientsMap[ingredient.IngredientName] = ingredient
	}
	return IngredientsMap
}

func CompareDataBases(OldMap map[string]dbreader.Cake, NewMap map[string]dbreader.Cake) {
	// Проверка на новый торт
	for _, cake := range NewMap {
		_, ok := OldMap[cake.Name]
		if !ok {
			fmt.Printf(`ADDED cake "%s"%s`, cake.Name, "\n")
		}
	}
	// Проверка на уделнные торты
	for _, cake := range OldMap {
		_, ok := NewMap[cake.Name]
		if !ok {
			fmt.Printf(`REMOVED cake "%s"%s`, cake.Name, "\n")
		}
	}
	// Проверка на изменение времени и ингридиенты
	for _, cake := range OldMap {
		newCake, ok := NewMap[cake.Name]
		if ok {
			time := cake.Time != newCake.Time
			if time {
				fmt.Printf(`CHANGED cooking time for cake "%s" - "%s" instead of "%s"%s`, cake.Name, newCake.Time, cake.Time,
					"\n")
			}
			cakeIngredients := IngridientsToMap(cake.Ingredients)
			newCakeIngredients := IngridientsToMap(newCake.Ingredients)
			for _, ingredient := range newCakeIngredients {
				_, ok := cakeIngredients[ingredient.IngredientName]
				if !ok {
					fmt.Printf(`ADDED ingredient "%s" for cake "%s"%s`, ingredient.IngredientName, cake.Name, "\n")
				}
			}
			for _, ingredient := range cakeIngredients {
				_, ok := newCakeIngredients[ingredient.IngredientName]
				if !ok {
					fmt.Printf(`REMOVED ingredient "%s" for cake "%s"%s`, ingredient.IngredientName, cake.Name, "\n")
				}
			}
			for _, ingredient := range cakeIngredients {
				newIngredient, ok := newCakeIngredients[ingredient.IngredientName]
				if ok {
					if ingredient.IngredientUnit != newIngredient.IngredientUnit && ingredient.IngredientUnit != "" && newIngredient.IngredientUnit != "" {
						fmt.Printf(`CHANGED unit for ingredient "%s" for cake "%s" - "%s" instead of "%s"%s`, ingredient.IngredientName, cake.Name, newIngredient.IngredientUnit, ingredient.IngredientUnit, "\n")
					}
					if ingredient.IngredientCount != newIngredient.IngredientCount {
						fmt.Printf(`CHANGED count for ingredient "%s" for cake "%s" - "%s" instead of "%s"%s`, ingredient.IngredientName, cake.Name, newIngredient.IngredientCount, ingredient.IngredientCount, "\n")
					}
					if ingredient.IngredientUnit != "" && newIngredient.IngredientUnit == "" {
						fmt.Printf(`REMOVED unit "%s" for ingredient "%s" for cake %s%s"`, ingredient.IngredientUnit, ingredient.IngredientName, cake.Name, "\n")
					}
				}
			}
		}
	}
}

func (f *Flags) initFlags() {
	flag.StringVar(&f.old, "old", "", "the file with the old recipe")
	flag.StringVar(&f.new, "new", "", "the file with the new recipe")
	flag.Parse()
}
