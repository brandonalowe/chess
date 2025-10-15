package engine

import (
	"fmt"
	"strings"
)

const (
	White = iota
	Black
)

const (
	Pawn = iota
	Knight
	Bishop
	Rook
	Queen
	King
	NumPieceTypes
)

type Board struct {
	Pieces          [2][NumPieceTypes]Bitboard // [colour][pieceType]
	SideToMove      int
	CastlingRights  uint8
	EnPassantSquare int
	HalfmoveClock   int
	FullmoveNumber  int
}

func Square(file, rank int) int {
	return rank*8 + file
}

// TODO: replace this constructor of new board with FEN notation constructor
func NewBoard() *Board {
	b := &Board{
		EnPassantSquare: -1,
		SideToMove:      White,
		CastlingRights:  0b1111, // KGkq
		FullmoveNumber:  1,
	}
	// White pieces
	b.Pieces[White][Pawn] = 0x000000000000FF00
	b.Pieces[White][Rook] = 0x0000000000000081
	b.Pieces[White][Knight] = 0x0000000000000042
	b.Pieces[White][Bishop] = 0x0000000000000024
	b.Pieces[White][Queen] = 0x0000000000000008
	b.Pieces[White][King] = 0x0000000000000010

	// Black pieces
	b.Pieces[Black][Pawn] = 0x00FF000000000000
	b.Pieces[Black][Rook] = 0x8100000000000000
	b.Pieces[Black][Knight] = 0x4200000000000000
	b.Pieces[Black][Bishop] = 0x2400000000000000
	b.Pieces[Black][Queen] = 0x0800000000000000
	b.Pieces[Black][King] = 0x1000000000000000

	return b
}

func (b *Board) All(colour int) Bitboard {
	var bb Bitboard
	for i := range NumPieceTypes {
		bb |= b.Pieces[colour][i]
	}

	return bb
}

func (b *Board) Occupied() Bitboard {
	return b.All(White) | b.All(Black)
}

func (b *Board) Print() {
	fmt.Println()
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d ", rank+1)
		for file := range 8 {
			sq := Square(file, rank)
			char := "."
			for colour := range 2 {
				for piece := range NumPieceTypes {
					if b.Pieces[colour][piece]&Bit(sq) != 0 {
						char = pieceChar(piece, colour)
					}
				}
			}
			fmt.Printf("%s ", char)
		}
		fmt.Println()
	}
	fmt.Println("  a b c d e f g h")
}

func pieceChar(piece, colour int) string {
	chars := []string{"P", "N", "B", "R", "Q", "K"}
	ch := chars[piece]
	if colour == Black {
		ch = strings.ToLower(ch)
	}
	return ch
}
