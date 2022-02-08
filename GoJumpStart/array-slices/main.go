package main

import "fmt"

func main() {
	// Arrays

	var fruitArr [2]string

	// Assaign value
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	// Declare and assaign same time

	var fruitArr2 = []string{"Orange", "Pineapple"}
	fruitArr3 := []string{"Orange", "Pineapple"}
	fruitArr4 := [2]string{"Orange", "Pineapple"}

	fmt.Println(fruitArr2)
	fmt.Println(fruitArr2[0])
	fmt.Println(fruitArr3[0])
	fmt.Println(fruitArr4[0])

	// Slices

	var fruitArr5 = []string{"Grape", "Cherry"}
	fruitArr6 := []string{"Grape", "Cherry", "Mango", "Peach"}
	fruitArr7 := [2]string{"Grape", "Cherry"}

	fmt.Println(fruitArr5)
	fmt.Println(fruitArr5[0:1]) // return sliced array, not string
	fmt.Println(fruitArr6[2:3])
	fmt.Println(fruitArr7[0:]) // no need to put len(fruitArr7)
}
