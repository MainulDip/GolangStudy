package main

import "fmt"

func main() {
	x := 6
	y := 7

	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else if y < x {
		fmt.Printf("%d is less than %d \n", y, x)
	} else {
		fmt.Printf("%d is equal to %d \n", y, x)
	}

	// <=, >= possible

	// Switch

	color := "red"

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Not Red Nor Blue")
	}

}
