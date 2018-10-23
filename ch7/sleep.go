package main

import ("fmt"
    "flag"
    "time"
  )


var period = flag.Duration("sleepPeriod", 1*time.Second, "Duration during which the program sleeps")

func main() {
  flag.Parse()
  fmt.Printf("Sleeping for %v ...", *period)
  time.Sleep(*period)
  fmt.Println()
}
