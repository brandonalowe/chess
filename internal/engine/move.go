package engine

type Move struct {
	From      int
	To        int
	Promotion int
	Capture   bool
}
