package main

import (
	"fmt"
)

var codingOn string = "Golang"

func main() {
	// fmt.Println("Hello World, Golang Starting")
	// Main Types
	// String
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr (unsigned integer, no negative numbers)
	// byte ~ alias for uint8
	// rune ~ alias for int32
	// float32, float64
	// complex64, complex128

	// var
	var name string = "first name"
	var age int = 7777
	const isCool bool = true // cannot reasign const
	shortHand := ":"         // not possible outside of the function
	floatNum := 77.77
	var floatNum32 float32 = 1.2

	var val1, val2 = "Value 1", 2
	var val3, val4 string = "Value 3", "Value 4"
	var (
		val5 int    = 5555555555
		val6 string = "Value 6"
	)
	val7, val8 := 7, "value 8"
	const val9, val10, val11 = "value 9", "value 10", 77

	fmt.Println(name, age, shortHand, isCool, "Coding on", codingOn)
	fmt.Println(val1, val2, val3, val4, val5, val6, val7, val8, val9, val10, val11)
	fmt.Printf("Name is %s and age is %d \n", name, age)
	fmt.Printf("Name is %v and age is %v \n", name, age)
	fmt.Printf("%T\n", age)
	fmt.Printf("%v\n", age)
	fmt.Printf("%T\n", floatNum)
	fmt.Printf("%T\n", floatNum32)
	// https://pkg.go.dev/fmt

	// go vet <filename.go> will print formatting errors

}
