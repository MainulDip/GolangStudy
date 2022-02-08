package main

import (
	"fmt"
	"gojumpstart/gopackages/strutils"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))           // print 2
	fmt.Println(math.Ceil(2.7))            // print 3
	fmt.Println(math.Sqrt(16))             // print 4
	fmt.Println(strutils.Reverse("olleH")) // print olleH
	fmt.Println(strutils.Reverse("olleH")) // print olleH
}
