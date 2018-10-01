package main

//import "fmt"

type ChessPieceType struct {
	name string
	// if the vertical variation of a move is i, and the horizontal one is j,
	// then if i*j is not in "moves", then it is not a legal move for the piece
	moves     []int
	ascii_rep string // used for board representation in CLI
	original_x []int // the original x positions of a ChessPiece
	original_y []int // the original y positions of a ChessPiece
}

type ChessPiece struct {
	chesspiecetype ChessPieceType
	color          int // 0 is white 1 is black
	id             int // each piece is identified by its unique id
	number	       int // over the number of piece of a certain type
}

func main() {
	// We create the board that will hold the position of the pieces
	//board := [8][8]int{}
	// We create its string representation
	board_rep := [8][8]string{}
	for i := 0; i <8; i++ {
		for j := 0; j <8; j++ {
			board_rep[i][j] = "[]"
			if i==j {
				board_rep[i][j] = "x"
			}
		}
	}

	//Creating the complex moves
	var square_list = []int{}
	for i := 1; i < 9; i++ {
		square_list = append(square_list, i*i, -i*i)
	}
	queen_list := append(square_list, 0)

	//Creating the different types of pieces
	bishop := ChessPieceType{"bishop", square_list, "F", []int{2,5}, []int{0,7}}
	knight := ChessPieceType{"knight", []int{-2, 2}, "H",[]int{1,6}, []int{0,7}}
	king := ChessPieceType{"king", []int{-1, 0, 1}, "K", []int{4}, []int{0,7}}
	queen := ChessPieceType{"queen", queen_list, "Q", []int{3}, []int{0,7}}
	tower := ChessPieceType{"tower", []int{0}, "T", []int{0,7}, []int{0,7}}
	pawn := ChessPieceType{"pawn", []int{-1, 0, 1}, "P", []int{0,1,2,3,4,5,6,7}, []int{1,6}}

	//Each pieces appears a certain number of time
	chess_set := make(map[int][]ChessPieceType)
	chess_set[1] = []ChessPieceType{queen, king}
	chess_set[2] = []ChessPieceType{bishop, knight, tower}
	chess_set[8] = []ChessPieceType{pawn}

	//We create the game with the previously defined variables
	chess_game := make(map[int]ChessPiece)
	for number, piece_slice := range chess_set {
		for i := 0; i < len(piece_slice); i++ {
			piece_type := piece_slice[i]
			for j := 0; j < number; j++ {
				id_white := len(chess_game) + 1
				id_black := id_white + 1
				//We create each piece for one color
				chess_game[id_white] = ChessPiece{piece_type, 0, id_white, j}
				chess_game[id_black] = ChessPiece{piece_type, 1, id_black, j}
			}
		}
	}
       /* for _, piece := range chess_game {*/
		//if piece.piece_type =="pawn"{
			//if piece.color == 0{
				//board[1][piece.number] = piece.id
				//board_rep[1][piece.number] = piece.rep
			//}
			//if piece.color == 1{
				//board[6][piece.number] = piece.id
				//board_rep[6][piece.number] = piece.rep
			//}
			//}
		//if piece.piece_type =="bishop"{


		/*}*/
}
