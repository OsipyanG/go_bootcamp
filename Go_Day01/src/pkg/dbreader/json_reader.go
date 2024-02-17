package dbreader

type JsonReader struct {
	Cakes []Cake `json:"cake"`
}
type Cake struct {
	Name        string
	Time        string
	Ingredients []Ingredient
}
type Ingredient struct {
	IngredientName  string
	IngredientCount string
	IngredientUnit  string
}

func (jr *JsonReader) UnmrshalJson() {}
func (jr *JsonReader) MarshalXml()   {}
