package crusher

import . "github.com/barnex/chess"

func MaterialValue(b *Board) int {
	v := 0
	for _, p := range b {
		v += ValueOf(p)
	}
	return v
}

func (n *Node) KingTaken() bool {
	return n.value > 100 || n.value < -100
}

func ValueOf(p Piece) int {
	return valueOf[p+6]
}

var valueOf = [13]int{
	WP + 6: 1,
	BP + 6: -1,

	WN + 6: 3,
	BN + 6: -3,

	WB + 6: 3,
	BB + 6: -3,

	WR + 6: 5,
	BR + 6: -5,

	WQ + 6: 9,
	BQ + 6: -9,

	WK + 6: 1000,
	BK + 6: -1000,
}
