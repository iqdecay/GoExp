package main

import (
	"fmt"
)
func main() {
	// Bitwise operation

	/* These operation operate on bits only : 
	^ is the XOR operator
	&^ is the NAND operator :
	z = x &^ y 
	A bit in z is 0 iff the corresponding y bit is 1, otherwise 
	it is the corresponding bit of x
	<< is the left shift
	>> is the right shift
	| is the bitwise OR
	& is the bitwise AND
	*/
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n",x)
	fmt.Printf("%08b\n",y)
	fmt.Printf("x&y : %08b\n",x&y)
	fmt.Printf("x|y : %08b\n",x|y)
	fmt.Printf("x^y : %08b\n",x^y)
	fmt.Printf("x&^y : %08b\n",x&^y)
	for i := uint(0); i<8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i)
		}
	}
	fmt.Printf("%08b\n", x << 1)
	fmt.Printf("%08b\n", x >> 1)

}
