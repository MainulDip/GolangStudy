package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		fmt.Println("Value of i is => " + strconv.Itoa(i))
		fmt.Println(fmt.Sprint("Hello ", i))
		// https://go.dev/doc/effective_go#printing
		i++
	}
	// Short Method
	for x := 0; x <= 10; x++ {
		fmt.Println("Counter x is => " + strconv.Itoa(x))
		fmt.Println("Counter x is => ", x)
	}

	// FixxBuzz

	for y := 0; y <= 100; y++ {
		if y%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if y%3 == 0 {
			fmt.Println("Fizz")
		} else if y%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(y)
		}
	}
}
