package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 20; i += 4 {
		fmt.Println(i)
	}
	j := 0
	for j < 20 {
		fmt.Println(j)
		j += 5
	}
	a := 3
	for x := 3; a < 25; x += 3 {
		fmt.Println("Do something !", x)
		a += 4
	}

}
