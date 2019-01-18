golang versus rust -- JSON deserialisation into simple structs


### Go

```
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
```

### Rust

```
#[macro_use]
extern crate serde_derive;
extern crate serde;
extern crate serde_json;

use std::io;

fn main() -> Result<(), io::Error> {
    let data = r#"{"name": "John Doe"}"#;
    let person: Person = serde_json::from_str(data)?;
    println!("{:?}", person);

    // Output: Error: Custom { kind: InvalidData, error: Error("missing field `age`", line: 1, column: 20) }

    Ok(())
}

#[derive(Debug, Deserialize)]
struct Person {
    name: String,
    age: i32,
}

```
