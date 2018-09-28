package main

import (
	"fmt"
)

//type Move struct {
//// This type represents a move of a particular piece via the coordinates variation it entails
//letter_variation int
//number_variation int
//}

type ChessPieceType struct {
	name      string
	moves     []int
	ascii_rep string
}

type ChessPiece struct {
	chesspiecetype ChessPieceType
	color          int // 0 is white 1 is black
	id             int
}

/*type PieceSet struct {*/
//// the number of pieces of each type that has to be created
//piece_type ChessPieceType
//number     int
/*}*/

func main() {
	var square_list = []int{}
	for i := 1; i < 9; i++ {
		square_list = append(square_list, i*i, -i*i)
	}
	fmt.Println(square_list)
	bishop := ChessPieceType{"bishop", square_list, "F"}
	knight := ChessPieceType{"knight", []int{-2, 2}, "H"}
	king := ChessPieceType{"king", []int{-1, 0, 1}, "K"}
	queen := ChessPieceType{"queen", []int{}, "Q"}

	first_white_bishop := ChessPiece{bishop, 0, 0}
	first_white_knight := ChessPiece{knight, 0, 1}
	white_king := ChessPiece{king, 0, 2}
	white_queen := ChessPiece{queen, 0, 3}
	fmt.Println(first_white_bishop)
	chess_set := make(map[string]int)
	// chess_set is the number of pieces of each type for a given color
	// the following lines can be rationnalized if necessary
	chess_set["bishop"] = 2
	chess_set["pawn"] = 8
	chess_set["king"] = 1
	chess_set["queen"] = 1
	chess_set["knight"] = 2
	chess_set["tower"] = 2
	fmt.Println(first_white_bishop, first_white_knight, white_king, white_queen)

	fmt.Println(chess_set)
}
