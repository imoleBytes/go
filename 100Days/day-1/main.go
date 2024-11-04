package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Person: a custom type, Person is defined
	type Person struct {
		Name    string
		Age     int
		Gender  string
		Hobbies []string
	}
	// an object, p of type Person is defined.      -- object here is lack of a better word! no objects in go!
	p := Person{
		Name:    "Ayo Balogun",
		Age:     34,
		Gender:  "Male",
		Hobbies: []string{"singing", "peppering osapkolo"},
	}

	json_str, _ := json.Marshal(p)
	fmt.Println(string(json_str))
	// Output:
	// {"Name":"Ayo Balogun","Age":34,"Gender":"Male","Hobbies":["singing","peppering osapkolo"]}

	// lets now assume this json string is what we got from another program or a text file
	str := `{"Name":"Ayo Balogun","Age":34,"Gender":"Male","Hobbies":["singing","peppering osapkolo"]}`
	var p2 Person
	json.Unmarshal([]byte(str), &p2)

	// Now we can use the data in memory
	fmt.Println(p2.Name)
	fmt.Println(p2.Gender)
	// Output:
	// Ayo Balogun
	// Male
}
