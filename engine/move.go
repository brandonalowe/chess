package engine

import "strconv"

type Move struct {
	From int
	To   int
}

func (m Move) UCI() string {
	return SquareToStr(m.From) + SquareToStr(m.To)
}

func SquareToStr(sq int) string {
	file := sq % 8
	rank := sq / 8
	return string(rune('a'+file)) + strconv.Itoa(rank+1)
}
