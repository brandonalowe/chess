package engine

func GenerateMoves(b *Board) []Move {
	var moves []Move

	if b.SideToMove == White {
		moves = append(moves, generatePawnMoves(b, White)...)
	}

	return moves
}

func generatePawnMoves(b *Board, colour int) []Move {
	var moves []Move
	pawns := b.Pieces[colour][Pawn]
	empty := ^b.Occupied()

	var singlePushes Bitboard
	var doublePushes Bitboard

	if colour == White {
		singlePushes = pawns.North() & empty
		doublePushes = (singlePushes.North() & empty) & Rank4
	} else {
		singlePushes = pawns.South() & empty
		doublePushes = (singlePushes.North() & empty) & Rank5
	}

	singlePushes.ForEachSquare(func(to int) {
		from := to - 8
		if colour == Black {
			from = to + 8
		}
		moves = append(moves, Move{From: from, To: to})
	})

	doublePushes.ForEachSquare(func(to int) {
		from := to - 16
		if colour == Black {
			from = to + 16
		}
		moves = append(moves, Move{From: from, To: to})
	})

	return moves
}
