package main
import (
    "fmt"
    "bufio"
    )

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error){
  count, p, _ := bufio.ScanWords(p, true)
  *c += WordCounter(p)
  return p, nil
}

func main() {
    // var c ByteCounter
    // c.Write([]byte("hello"))
    // fmt.Println(c)
    // c = 0
    // name := "Dolly"
    // fmt.Fprintf(&c, "hello, %s", name)
    // fmt.Println(c)
    var cprime WordCounter
    listOfWords := "Dolly the cloned sheep"
    fmt.Fprintf(&cprime, "hello, %s", listOfWords)
    fmt.Println(cprime)
    fmt.Fprintf(&cprime, "goodbye, %s", "my brother")
    fmt.Println(cprime)



}
