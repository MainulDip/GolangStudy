## GoLang Jumpstart:
personalized overview and instruction for everyday use golang projects and language structure.

### Workspact Setup:
Usually go projects are all seat into one structure.
```sh
mkdir GoProjects
mkdir src && mkdir bin
cd src && mkdir <github.com>
cd <github.com> && mkdir <projectDirectory> && cd <projectDirectory>
git clone <repo> .
```

### Commands Initials:
```sh
go run <file.go>
# go install will build the binary inside user home go directory
go mod init directory/filename
```

### Closure
> Closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

> For example, the adder function returns a closure. Each closure is bound to its own sum variable. 

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	prev, curr, i, x := 0, 0, 0, 0
	f := 0
	return func() int {
	// if f == 0 , g := f, f++ return g 
	// if f == 1 , curr := f, f++ return g
	// if f > 1 ,  prev = curr curr = prev + curr return prev + current
	
	if f <= 1 {
		curr = f
		x = curr
		f++
		println("f<=1", curr, prev)
		
	} else if f > 1 {
		x = curr
		i = prev
		println("f>1", curr, prev)
		
		curr = curr + prev
		prev = x
		
	}
	return x + i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

### Struct: like class but there is no class in golang
```go
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// two kinds of method, value recever (calcularion, not change anything)
// and pointer recever for changin something.
// method defines outside of the struct


// struct method
// Greeting method ( value reciever )
func (p Person) greet() string {
	return "Hello from the struct's value reciever method and firstname is " + p.firstName + " And last name is " + p.lastName + " and age is " + strconv.Itoa((p.age))
}

// Changing function (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// initialization
person3 := Person{"Nile", "Smith", "Miami", "Female", 77}
fmt.Println(person3)
person3.getMarried("Fox")
fmt.Println("Person3's new lastname => ", person3.lastName)
```

### interfaces: method signature for structs