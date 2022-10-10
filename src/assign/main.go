package main

import "fmt"

func main() {
	var a, b int

	a, b = 10, 5
	fmt.Println("Assigned values:\ta= ", a, "\tb = ", b)

	a += b // a = 10 + 5
	fmt.Println("Addion and assign:\ta= ", a)

	a -= b // a = 15 - 5
	fmt.Println("Subtract and assign:\ta= ", a)

	a *= b // a = 10 * 5
	fmt.Println("Multiply and assign:\ta= ", a)

	a /= b // a = 50 / 5
	fmt.Println("Divide and assign:\ta= ", a)

	a %= b // a = 10 % 5
	fmt.Println("Remainder and assign:\ta= ", a)
}
