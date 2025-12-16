package main

import (
	"fmt"
)

// ---------- Struct ----------
type Address struct {
    City    string
    Country string
}

type Person struct {
    Name string
    Age  int
    Address  // embedding
}

// ---------- Constructor ----------
func NewPerson(name string, age int, city, country string) Person {
    return Person{
        Name: name,
        Age:  age,
        Address: Address{
            City:    city,
            Country: country,
        },
    }
}

// ---------- Factory ----------
func PersonFactory(name string, age int) Person {
    if age < 18 {
        fmt.Println("Creating a minor")
    } else {
        fmt.Println("Creating an adult")
    }

    return Person{Name: name, Age: age}
}

// ---------- Stringer ----------
func (p Person) String() string {
    return fmt.Sprintf("%s (%d) from %s, %s",
        p.Name, p.Age, p.City, p.Country)
}

func main() {
    // using constructor
    p1 := NewPerson("Nanda", 30, "Thimphu", "Bhutan")
    fmt.Println(p1)

    // using factory
    p2 := PersonFactory("Karma", 16)
    fmt.Println(p2)
}
