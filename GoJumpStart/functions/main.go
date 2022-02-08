package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func getSum2(num1, num2 int) int { // if parameters are same type, mention once at the last
	return num1 + num2
}

func main() {
	fmt.Println("Working on functions")
	fmt.Println(greeting("Golang"))
	fmt.Println(getSum(4, 3))
	fmt.Println(getSum2(4, 3))
}
