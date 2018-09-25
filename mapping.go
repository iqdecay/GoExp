package main

import "fmt"

func main() {
	grades := make(map[string]float32)
	grades["Mathisson"] = 1
	grades["Victor"] = 100
	grades["Timéa"] = 150
	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}
