package main

import (
	"fmt"
)

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

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
	b := Location{Loc: "string"}
	fmt.Println(b)

}
