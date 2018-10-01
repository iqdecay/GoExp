package main

import "fmt"

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

func main() {
	//Creating the complex moves
	var square_list = []int{}
	for i := 1; i < 9; i++ {
		square_list = append(square_list, i*i, -i*i)
	}
	queen_list := append(square_list, 0)

	//Creating the different types of pieces
	bishop := ChessPieceType{"bishop", square_list, "F"}
	knight := ChessPieceType{"knight", []int{-2, 2}, "H"}
	king := ChessPieceType{"king", []int{-1, 0, 1}, "K"}
	queen := ChessPieceType{"queen", queen_list, "Q"}
	tower := ChessPieceType{"tower", []int{0}, "T"}
	pawn := ChessPieceType{"pawn", []int{-1, 0, 1}, "P"}

	//Each pieces appears a certain number of time
	chess_set := make(map[int][]ChessPieceType)
	chess_set[1] = []ChessPieceType{queen, king}
	chess_set[2] = []ChessPieceType{bishop, knight, tower}
	chess_set[8] = []ChessPieceType{pawn}

	//We create the game with the previously defined variables
	chess_game := []ChessPiece{}
	for number, piece_slice := range chess_set {
		for i := 0; i < len(piece_slice); i++ {
			piece_type := piece_slice[i]
			for j := 0; j < number; j++ {
				id_white := len(chess_game) + 1
				id_black := id_white + 1
				//We create each piece for one color
				chess_game = append(chess_game, ChessPiece{piece_type, 0, id_white})
				chess_game = append(chess_game, ChessPiece{piece_type, 1, id_black})
			}

		}
	}
	fmt.Println(len(chess_game))
	for i := 0; i < len(chess_game); i++ {
		fmt.Println(chess_game[i])
	}
}
