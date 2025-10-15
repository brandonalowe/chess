package engine

import "math/bits"

type Bitboard uint64

// Edge Masks
const (
	FileA Bitboard = 0x0101010101010101
	FileB Bitboard = 0x0202020202020202
	FileG Bitboard = 0x4040404040404040
	FileH Bitboard = 0x8080808080808080

	Rank1 Bitboard = 0x00000000000000FF
	Rank2 Bitboard = 0x000000000000FF00
	Rank3 Bitboard = 0x0000000000FF0000
	Rank4 Bitboard = 0x00000000FF000000
	Rank5 Bitboard = 0x000000FF00000000
	Rank6 Bitboard = 0x0000FF0000000000
	Rank7 Bitboard = 0x00FF000000000000
	Rank8 Bitboard = 0xFF00000000000000

	EmptyBB Bitboard = 0
	FullBB  Bitboard = ^EmptyBB
)

// Utility functions
func Bit(sq int) Bitboard {
	return 1 << sq
}

func (bb Bitboard) LSB() int {
	if bb == 0 {
		return -1
	}
	return bits.TrailingZeros64(uint64(bb))
}

func (bb Bitboard) ForEachSquare(fn func(int)) {
	tmp := bb
	for tmp != 0 {
		sq := tmp.LSB()
		fn(sq)
		tmp &= tmp - 1
	}
}

func (bb Bitboard) Count() int {
	return bits.OnesCount64(uint64(bb))
}

// Shifts
func (bb Bitboard) North() Bitboard {
	return bb << 8
}

func (bb Bitboard) South() Bitboard {
	return bb >> 8
}

func (bb Bitboard) East() Bitboard {
	return (bb & ^FileH) << 1
}

func (bb Bitboard) West() Bitboard {
	return (bb & ^FileA) >> 1
}

func (bb Bitboard) NorthEast() Bitboard {
	return (bb & ^FileH) << 9
}

func (bb Bitboard) NorthWest() Bitboard {
	return (bb & ^FileA) << 7
}

func (bb Bitboard) SouthEast() Bitboard {
	return (bb & ^FileH) >> 7
}

func (bb Bitboard) SouthWest() Bitboard {
	return (bb & ^FileA) >> 9
}
