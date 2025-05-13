package main

import (
	"brandonalowe/chess/engine"
	"fmt"
)

const startingPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const randomPositionOne = "r4rk1/ppp3pp/6q1/2bP4/3Np3/2P4b/P1B2PP1/R2QR1K1 w - - 0 21"
const randomPositionTwo = "2k5/pp6/1n2p2p/2p1P1p1/2P5/PP2Q1B1/5P1P/q6K"
const randomPositionThree = "r1bqk2r/3n1ppp/2p2n2/1pb1p3/p3P3/P1NP1N2/BPP3PP/R1BQK2R"
const queenPosition = "8/3Rp3/8/3P4/4Q3/8/8/8"
const knightPosition = "8/8/8/8/4N3/8/8/8"
const bishopPosition = "8/8/8/8/4B3/8/8/8"
const rookPosition = "8/8/8/8/2R5/8/8/8"
const kingPosition = "8/8/8/8/4K3/8/8/8"
const pawnPosition = "8/8/8/8/8/8/4P3/8"

func main() {
	game := engine.LoadFEN(randomPositionOne)

	game.PrintBoard()
	fmt.Println(engine.SquareToStr(1))
	fmt.Println(engine.SquareToStr(63))

	// gameGui := &gui.ChessGame{}
	// err := game.LoadPieceImages()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// gameGui.LoadBitBoard(game)
	// ebiten.SetWindowSize(60*8, 60*8)
	// ebiten.SetWindowTitle("Chess")

	// if err := ebiten.RunGame(gameGui); err != nil {
	// 	log.Fatal(err)
	// }

}
