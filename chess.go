package main

import (
	"fmt"
)

type Move struct {
	// This type represents a move of a particular piece via the coordinates variation it entails
	letter_variation int
	number_variation int
}

type ChessPieceType struct {
	name      string
	moves     []Move
	ascii_rep string
}

type ChessPiece struct {
	chesspiecetype ChessPieceType
	color          int // 0 is white 1 is black
	id             int
}

func main() {
	fool_move := Move{1, 1}
	var fool_moves []Move = []Move{fool_move}
	fool := ChessPieceType{"fool", fool_moves, "F"}
	fmt.Println(fool)
	first_white_fool := ChessPiece{fool, 0, 0}
	fmt.Println(first_white_fool)
}
