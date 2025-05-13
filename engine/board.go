package engine

import "fmt"

type Bitboard uint64

const (
	White = 0
	Black = 1
)

const (
	Pawn = iota
	Knight
	Bishop
	Rook
	Queen
	King
	PieceNone = -1
)

// const (
// 	fileChars = "abcdefgh"
// 	rankChars = "12345678"
// )

var unicodePieces = [2][6]string{
	{ // White
		"♙", "♘", "♗", "♖", "♕", "♔",
	},
	{ // Black
		"♟︎", "♞", "♝", "♜", "♛", "♚",
	},
}

type GameState struct {
	Bitboards  [2][6]Bitboard
	Occupancy  [2]Bitboard
	AllPieces  Bitboard
	SideToMove int
}

func SetBit(bb Bitboard, sq int) Bitboard {
	return bb | (1 << sq)
}

func GetBit(bb Bitboard, sq int) bool {
	return (bb & (1 << sq)) != 0
}

func (g *GameState) PrintBoard() {
	fmt.Println("  +-----------------+")
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d |", rank+1)
		for file := 0; file < 8; file++ {
			sq := rank*8 + file
			piece := "."
			for color := 0; color < 2; color++ {
				for pt := 0; pt < 6; pt++ {
					if GetBit(g.Bitboards[color][pt], sq) {
						piece = unicodePieces[color][pt]
					}
				}
			}
			fmt.Printf(" %s", piece)
		}
		fmt.Println(" |")
	}
	fmt.Println("  +-----------------+")
	fmt.Println("    a b c d e f g h")
}
