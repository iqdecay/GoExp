package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibo(1000))
}

func fibo(n int) (int){
	var x, y int	
	x, y = 0, 1 
	for i := 0; i<n; i++{
		x, y = y, x+y
	}
	return y
}

