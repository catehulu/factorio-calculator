package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]string
	Data      map[string]interface{}
	Recipe    map[string]Recipe
}
