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

