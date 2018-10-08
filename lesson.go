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


	Unicode encoding : 
	Unicode code point is the standard number associated with 
	a letter, an accent, any character in any language
	In go, a unicode code point is referred to as a rune
	Since there are so many code points, unicode runes are encoded on 
	int32 types (that is, an int encoded on 32 bits)
	UTF-8 uses 1 byte for ASCII chars, and only 2 or 3 for most
	common characters in use
	The high order bits of the first byte of the encoding for a rune 
	indicates how many byte follow
	0 means 7-bit ASCII (hence 0 full bytes follow)
	110 means that the rune takes two bytes 
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
