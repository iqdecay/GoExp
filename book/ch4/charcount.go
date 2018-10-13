package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    freq := make(map[string]float32)
    files := os.Args[1:]
    if len(files) == 0 {
        fmt.Println("Call the program with filenames in input !")
    } else {
        for _, arg := range files {
            numberOfWords := 0
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr,"%v", err)
                continue
            }
            numberOfWords = frequency(f, freq)
            f.Close()

        for word, count := range freq {
            if count > 0 {
                fmt.Printf("%s :\t %f\n", word, count/float32(numberOfWords))
            }
        }
      }
   }
}

func frequency(f *os.File, freq map[string]float32) (int){
    input := bufio.NewScanner(f)
    input.Split(bufio.ScanWords)
    i := 0
    for input.Scan() {
        freq[input.Text()]++
        i ++
        }
    return i
        // Note : ignoring potential errors from input.Err()
}
