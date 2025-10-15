package main

import (
	"chess-engine/internal/engine"
	"fmt"
)

func main() {
	b := engine.NewBoard()
	b.Print()

	moves := engine.GenerateMoves(b)
	fmt.Printf("Generated %d moves\n", len(moves))
	for _, m := range moves {
		fmt.Printf("%d -> %d\n", m.From, m.To)
	}
}
