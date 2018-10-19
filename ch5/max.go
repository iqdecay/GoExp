package main

import (
  "fmt"
)
func max(a int, vals ...int) (int){
  m := a
  for _, val := range vals {
    if m < val {
      m = val
    }
  }
  return m
}

func main(){
  fmt.Println(max(7,54,1,450))
}
