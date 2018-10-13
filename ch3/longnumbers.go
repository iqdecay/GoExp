package main

import "fmt"

func space (s string) string{
	n := len(s)
	if n <= 3 {
		return s
		}
	return space(s[:n-3]) + " " + s[n-3:]
}

func main() {
	fmt.Println(space("123456789"))
}
