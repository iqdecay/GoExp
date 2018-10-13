package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow, extend the slice
		z = x[:zlen]
	}else {
		// There is unsufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zlen = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // See page 88
	}
	z[len(x)] = y
	return z
}

func main(){
	fmt.Println([]int{1,2,4} , 5)
}
