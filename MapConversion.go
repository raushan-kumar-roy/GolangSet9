package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	ingredients := map[string]string{
		"Thinly Sliced Peeled Apples": "6 Cups",
		"sugar":                       "3/4 cup",
		"flour":                       "2 tablespoons",
		"cinamon":                     "1/4 teaspoon",
		"nutmeg":                      "1/8 tablespoon",
		"lemon juice":                 "1 tablespoon",
	}
	for key, value := range ingredients {
		fmt.Println(key, value)
	}
	jsonData, err := json.MarshalIndent(ingredients, "", " ")
	if err != nil {
		fmt.Println("Error marshalling ingredients to JSON:", err)
		return
	}

	err = ioutil.WriteFile("ingredients.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

}
