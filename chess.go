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
	name string
	// if the vertical variation of a move is i, and the horizontal one is j,
	// then if i*j is not in "moves", then it is not a legal move for the piece
	moves     []int
	ascii_rep string
}

type ChessPiece struct {
	chesspiecetype ChessPieceType
	color          int // 0 is white 1 is black
	id             int // each piece is identified by its unique id
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
	queen_list := append(square_list, 0)
	bishop := ChessPieceType{"bishop", square_list, "F"}
	knight := ChessPieceType{"knight", []int{-2, 2}, "H"}
	king := ChessPieceType{"king", []int{-1, 0, 1}, "K"}
	queen := ChessPieceType{"queen", queen_list, "Q"}
	tower := ChessPieceType{"tower", []int{0}, "T"}
	pawn := ChessPieceType{"pawn", []int{-1, 0, 1}, "P"}
	/* chess_set := make(map[ChessPieceType]int)*/

	//// chess_set is the number of pieces of each type for a given color
	//// the following lines can be rationnalized if necessary
	//chess_set[bishop] = 2
	//chess_
	//chess_set[pawn] = 8
	//chess_set[king] = 1
	//chess_set[queen] = 1
	//chess_set[knight] = 2
	//chess_set[tower] = 2
	chess_set := make(map[int][]ChessPieceType)
	chess_set[1] = []ChessPieceType{queen, king}
	chess_set[2] = []ChessPieceType{bishop, knight, tower}
	chess_set[8] = []ChessPieceType{pawn}

	chess_game := []ChessPiece{}
	for number, piece_slice := range chess_set {
		for i := 0; i < len(piece_slice); i++ {
			piece_type := piece_slice[i]
			for j := 0; j < number; j++ {
				id_white := len(chess_game) + 1
				id_black := id_white + 1
				chess_game = append(chess_game, ChessPiece{piece_type, 0, id_white})
				chess_game = append(chess_game, ChessPiece{piece_type, 1, id_black})
			}

		}
	}
	fmt.Println(len(chess_game))
	for i := 0; i < len(chess_game); i++ {
		fmt.Println(chess_game[i])
	}

	//first_white_bishop := ChessPiece{bishop, 0, 1}
	//first_white_knight := ChessPiece{knight, 0, 2}
	//white_king := ChessPiece{king, 0, 3}
	//white_queen := ChessPiece{queen, 0, 4}
	//first_white_pawn := ChessPiece{pawn, 0, 5}
	/*first_white_tower := ChessPiece{tower, 0, 6}*/
}
