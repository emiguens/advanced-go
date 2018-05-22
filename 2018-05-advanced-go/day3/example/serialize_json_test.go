package serializejson_test

import (
	"encoding/json"
)

// START 1 OMIT
type Person struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Level    int    `json:"-"`
}

func testJson() {
	data := []byte(`
		{
			"name": "carlos",
			"lastname": "petruza",
		}
	`)
	var person Person
	if err := json.Unmarshal(data, &person); err != nil { // HLxxx
		panic(err)
	}
}

// END 1 OMIT
