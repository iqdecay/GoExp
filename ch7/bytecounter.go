package main
import (
    "fmt"
    "bufio"
    "io"
    )

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error){
  count, _, err := bufio.ScanWords(p, true)
  fmt.Println("count :",count, "p :", p)
  if err != nil {
    return 0, err
  }
  *c += WordCounter(count)
  return count, nil
}

type ByteWrittenCounter struct {
  w io.Writer
  written int64
}

func (c *ByteWrittenCounter) Write(p []byte) (int, error) {
  n, err := c.w.Write(p)
  c.written += int64(n)
  return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64){
  c := &ByteWrittenCounter{w, 0}
  return c, &c.written
}


func main() {
    var c ByteCounter
    c.Write([]byte("hello"))
    fmt.Println(c)
    c = 0
    name := "Dolly"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c)
    var cprime WordCounter
    listOfWords := "Dolly the cloned sheep"
    fmt.Fprintf(&cprime, "hello, %s", listOfWords)
    fmt.Println(cprime)
    fmt.Fprintf(&cprime, "goodbye, %s", "my brother")
    fmt.Println(cprime)



}
