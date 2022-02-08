package main

import "fmt"

func main() {
	//  Map is a key value pair
	//  Define map

	emails := make(map[string]string)
	// The make built-in function allocates and initializes an object of type slice, map, or chan

	// Assign key value
	emails["Bob"] = "bob@email.com"
	emails["Sharon"] = "sharon@email.com"
	emails["Bill"] = "bill@email.com"

	fmt.Println(emails, "\n", len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map

	delete(emails, "Bill")
	fmt.Println("Email maps after deleting: ", emails)

	// Defining and adding value at the same time
	address := map[string]string{
		"Bob":    "Murfessboro TN",
		"Sharon": "Atlanta GA",
		"Bill":   "Norfork VA",
	}

	address["Nabil"] = "Miami FL"
	// Note map will sort by alphabetic order when populuted

	fmt.Println(address)

}
