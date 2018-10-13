package main

import (
    "fmt"
)

func main() {
    type Point struct {
        X, Y int
        }
    type Circle struct {
        Point
        Radius int
        }
    type Wheel struct {
        Circle
        Spokes int
        }
    var w Wheel
    w.X = 1
    w.Y = 2
    w.Radius = 4
    w.Spokes = 5
    fmt.Printf("%#v",w)
}
    
