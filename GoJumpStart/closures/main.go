package main

import "fmt"

/*
A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables;
in this sense the function is "bound" to the variables.

For example, the adder function returns a closure.
Each closure is bound to its own sum variable. and
keep track of the previous returned value
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum, otherSum := adder(), adder() // closure is not directly accessible
	counter := 0
	for i := 0; i < 10; i++ {
		// fmt.Println(counter, "+", i, "=", sum(i))
		fmt.Println("sum() => ", sum(i), "|| otherSum() => ", otherSum(i+1))
		counter += i
	}
	fmt.Println("Closures Testing")
	fmt.Println("Good Going")
	// Closures explanation
	fmt.Println("Golden Good")

	f := 7
	g := f
	f = 12

	println("G value => ", g)
	println("f value => ", f)

}
