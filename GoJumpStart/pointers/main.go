package main

import "fmt"

func main() {
	// Point allowes to point to a value in the memory/location address thats in a variable
	a := 1

	b := &a // memory address

	println(b)     // print memory address like 0xc00001a098
	fmt.Println(a) // print 1
	fmt.Println(b) // 0xc00001a098
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // *int , "*" this is a pointer sign
	// Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)
	// Change Value With Pointer
	*b = 7
	fmt.Println(a)  // print 7
	fmt.Println(b)  // print like 0xc0000aa058
	fmt.Println(*b) // print 7
}
