package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
       /* reader := bufio.NewReader(os.Stdin)*/
	//fmt.Println("Enter your name :")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	//lettersToInt := make(map[byte]int)
	//letters := [8]byte{'a','b','c','d','e','f','g','h'}
	//for index, letter := range letters {
		//lettersToInt[letter] = index
	//}
	//a := "abcde"

	//fmt.Println(lettersToInt[a[2]])
	a := `learning go is pretty hard 
when you	think
	about 
	it, isn't \x000 it ?
	`
	fmt.Println(a)
	

}
