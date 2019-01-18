package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

func main() {
	var person Person
	data := []byte(`{"name":"John Doe"}`)
	if err := json.Unmarshal(data, &person); err != nil {
		panic(err)
	}
	fmt.Printf("person = %s\n", person)

	// output: person = Person{Name: John Doe, Age: 0}
}
