## GoLang Language Tour:
personalized overview and instruction for everyday use golang projects and language structure.

### Workspace Setup:
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

// fibonacci is a function that returns a function, and that function returns an int.
func fibonacci() func() int {

	prev, curr, i, x := 0, 0, 0, 0
	f := 0
	return func() int {
	
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

### Struct: like class in other programming languages, no class declaration in go.
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

### interfaces:
> Interfaces are named collections of method signatures. Or method signature for structs
```go
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}

// Results
// {3 4}
// 12
// 14
// {5}
// 78.53981633974483
// 31.41592653589793
```