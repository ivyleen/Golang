package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, in := range ints {
		if (in % 2) == 0 {
			fmt.Println("Number ", in, "is even.")
		} else {
			fmt.Println("Number ", in, "is odd.")
		}
	}
}
