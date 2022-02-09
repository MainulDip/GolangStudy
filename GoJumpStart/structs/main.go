package main

import (
	"fmt"
	"strconv"
)

// Define Person Struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// two kinds of method, value recever (calcularion, not change anything)
	// and pointer recever for changin something.
	// method defines outside of the struct
}

// struct method
// Greeting method ( value reciever )
func (p Person) greet() string {
	return "Hello from the struct's value reciever method and firstname is " + p.firstName + " And last name is " + p.lastName + " and age is " + strconv.Itoa((p.age))
}

// Changing function (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Albert", lastName: "Einistine", city: "Berlin", gender: "Male", age: 77}
	fmt.Println(person1)
	person2 := Person{"Nicolo", "Tesla", "Smiljan", "Male", 93}
	fmt.Println(person2)

	person3 := Person{"Nile", "Smith", "Miami", "Female", 77}
	fmt.Println(person3)

	fmt.Println("First Name", person2.firstName)
	person2.city = "New York"
	fmt.Println("Last city", person2.city)

	// method calling
	fmt.Println("greet() => ", person2.greet())
	person2.hasBirthday()
	fmt.Println("hasBirthDay() => ", person2.age)

	person3.getMarried("Fox")
	fmt.Println("Person3's new lastname => ", person3.lastName)
}
