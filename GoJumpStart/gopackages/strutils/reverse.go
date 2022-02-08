// package main
package strutils

// import "fmt"

func Reverse(s string) string {
	runes := []rune(s) // convert s into int32 (alias rune) => [72 101 108 108 111 55]
	// fmt.Println(runes)
	// fmt.Println(len(runes))

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		// i = 0, j = 4 | h e l l o | middle char will not change position for odd value
		// h will suffel with o and e will suffel with l

		// fmt.Println(string(runes[i]), string(runes[j]), string(runes[j]), string(runes[i]))

		runes[i], runes[j] = runes[j], runes[i]

		// fmt.Println(string(runes[i]), string(runes[j]), string(runes[j]), string(runes[i]))
	}
	// fmt.Println(string(runes))
	return string(runes)
}

// func main() {
// 	Reverse("Hello7")
// }
