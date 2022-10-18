package main

import "fmt"

func main() {
	// Byte = 8 bits, each with possible values 0 and 1. Combined all those bits can represent numbers between 0 and 255.
	// nibble = 4 bits or half byte

	// If the bits are written from right-to-left we have on the Least significant bit (LSB) on the very right and the
	// Most significant bit (MSB) on the far left.

	// To represent number 50 we need to add all the contributing bits (aka bits with value 1) in the binary representation:
	// bit number 	8 	7 	6 	5 	4 	3 	2 	1
	// decimal		128 64 	32 	16  8 	4 	2 	1
	// binary 		0 	0 	1 	1 	0 	0 	1 	0
	// Aka: 32 + 16 + 2 makes 50.

	// Bitwise operators
	// &	AND 			0011 & 0101 = 0001  - only if both corresponding bits are 1 we get a 1 as result
	// |	OR 				0011 | 0101 = 0111  - if either of the corresponding bits are 1 we get a 1 as result
	// ^	XOR 			0011 ^ 0101 = 0110  - only if the corresponding bits are different we get a 1 as result
	// ^	NOT 			^0101 = 1010  	    - unary operator - reverse all the bits
	// &^	AND NOT 		0011 &^ 0101 = 0010 -
	// <<	Shift left 		0010 << 2 = 1000
	// >>	Shift right 	1000 >> 2 = 0010

	var flags byte = 10 // 1010 (1x8 0 0x4 1x2 0x1)

	fmt.Println("Flag #1:%v", (flags&1) > 0)
	fmt.Println("Flag #2:%v", (flags&2) > 0)
	fmt.Println("Flag #3:%v", (flags&4) > 0)
	fmt.Println("Flag #4:%v", (flags&8) > 0)

	// %b is format specifier used to output the binary numerical value in each bit
	// Println doesn't work with %08b

	fmt.Printf("Flags value: %08b \t%v\n", flags, flags)

	flags = flags >> 1
	fmt.Printf("New value: %08b \t%v\n", flags, flags)
}
