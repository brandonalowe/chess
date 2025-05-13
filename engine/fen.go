package engine

import (
	"strings"
)

var pieceMap = map[rune]struct {
	color int
	pt    int
}{
	'P': {White, Pawn}, 'N': {White, Knight}, 'B': {White, Bishop},
	'R': {White, Rook}, 'Q': {White, Queen}, 'K': {White, King},
	'p': {Black, Pawn}, 'n': {Black, Knight}, 'b': {Black, Bishop},
	'r': {Black, Rook}, 'q': {Black, Queen}, 'k': {Black, King},
}

func LoadFEN(fen string) (g GameState) {
	rank := 7
	file := 0

	fields := strings.Fields(fen)
	if len(fields) < 2 {
		panic("Invalid FEN string")
	}

	piecePlacement := fields[0]
	sideToMove := fields[1]

	for _, c := range piecePlacement {
		switch {
		case c == '/':
			rank--
			file = 0
		case c >= '1' && c <= '8':
			file += int(c - '0')
		default:
			if piece, ok := pieceMap[c]; ok {
				sq := rank*8 + file
				g.Bitboards[piece.color][piece.pt] = SetBit(g.Bitboards[piece.color][piece.pt], sq)
				file++
			}
		}
	}

	if sideToMove == "w" {
		g.SideToMove = White
	} else {
		g.SideToMove = Black
	}

	for color := 0; color < 2; color++ {
		for pt := 0; pt < 6; pt++ {
			g.Occupancy[color] |= g.Bitboards[color][pt]
			g.AllPieces |= g.Bitboards[color][pt]
		}
	}

	return g
}
