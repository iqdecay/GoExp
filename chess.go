package main

import (
	"fmt"
)

type Move struct {
	// This type represents a move of a particular piece via the coordinates variation it entails
	letter_variation int
	number_variation int
}

type ChessPiece struct {
	name  string
	id    int
	moves []Move
	color int // 0 is white, 1 is black
}
