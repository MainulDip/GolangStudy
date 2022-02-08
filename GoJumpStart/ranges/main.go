package main

import "fmt"

func main() {
	ids := []int{11, 20, 32, 47, 44, 26, 27}
	// Loop through ids
	fmt.Println(ids)
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// if don't want to use index value, use "_" instade of i
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all the ids togather
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Doing Ranges adding all the ids and the Sum is : ", sum)

	// Range with map
	address := map[string]string{
		"Bob":    "Murfessboro TN",
		"Sharon": "Atlanta GA",
		"Bill":   "Norfork VA",
	}

	address["Nabil"] = "Miami FL"

	for k, v := range address {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range address {
		fmt.Printf("Key: %s\n", k)
	}
	for _, v := range address {
		fmt.Printf("Value: %s\n", v)
	}
}
