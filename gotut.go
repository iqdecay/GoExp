package main

import "fmt"

func main() {
	x := 25
	a := &x
	fmt.Println(a)
	fmt.Println(*a)
	*a = 15
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)

}
