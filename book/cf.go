package main

import (
	"fmt"
	"os"
	"strconv"
	"tutorial/book/tempconv"
)

fun main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
