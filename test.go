package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"bob": 14,
		"jean": 15,
		"victor": 21,
	}
	age, ok := ages["bob"]
	age1, ok1 := ages["albane"]
	fmt.Println(age, ok, age1, ok1)
	var b bool
	fmt.Println(b)
}
