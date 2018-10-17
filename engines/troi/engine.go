package troi

import (
	"fmt"
	"sort"

	. "github.com/barnex/chess"
)

func New(depth int) Engine {
	return &troi{depth}
}

type troi struct {
	depth int
}

func (e *troi) Move(b *Board, c Color) (Move, float64) {
	bestScore := Inf(-1)
	var allScores []ms
	for _, m := range AllMoves(b, c) {
		n := &Node{*b, Heuristic2(b, m)}
		s := float64(e.negamax(n, e.depth, c, m))
		allScores = append(allScores, ms{m, s})
		if s > bestScore {
			bestScore = s
		}
	}

	//printMS("initial", allScores)
	allScores = cutoff(allScores)
	//printMS("cutoff", allScores)

	for i, ms := range allScores {
		allScores[i].score = strategic(b.WithMove(ms.move), c)
	}
	sortMS(allScores)
	//printMS("strategical", allScores)

	best := allScores[0]
	return best.move, bestScore + best.score/100
}

func strategic(b *Board, c Color) float64 {
	//return 0
	//return mobility(b, c) - mobility(b, -c)
	//return threat(b, c)
	return 3*protection(b, c) + threat(b, c) + mobility(b, c) - mobility(b, -c)
}

func mobility(b *Board, c Color) float64 {
	return float64(len(AllMoves(b, c)))
}

func protection(b *Board, c Color) float64 {
	s := 0.0
	for _, m := range AllPre(b, c) {
		if b.At(m.Dst).Color() == c {
			s++
		}
	}
	return s
}

func threat(b *Board, c Color) float64 {
	s := 0.0
	for _, m := range AllPre(b, c) {
		if b.At(m.Dst).Color() == -c {
			s++
		}
	}
	return s
}

func cutoff(allScores []ms) []ms {
	sortMS(allScores)
	cuti := 0
	cutoff := allScores[cuti].score
	for i, ms := range allScores {
		if ms.score < cutoff {
			cuti = i
			break
		} else {
			cuti = i + 1
		}
	}
	allScores = allScores[:cuti]
	return allScores
}

func sortMS(allScores []ms) {
	sort.Slice(allScores, func(i, j int) bool {
		return allScores[i].score > allScores[j].score
	})
}

func printMS(msg string, allScores []ms) {
	fmt.Print(msg, ":")
	for _, ms := range allScores {
		fmt.Printf("(%v %v) ", ms.move, ms.score)
	}
	fmt.Println()
	fmt.Println()
}

type ms struct {
	move  Move
	score float64
}

type Node struct {
	board Board
	value int
}

func (n *Node) WithMove(m Move) *Node {
	return &Node{
		*n.board.WithMove(m),
		Heuristic2(&n.board, m),
	}
}

func (e *troi) negamax(n *Node, depth int, c Color, m Move) int {
	if dst := n.board.At(m.Dst); dst == WK || dst == BK {
		return inf(-c * dst.Color())
	}

	if depth == 0 {
		return int(c) * Heuristic3(n, m)
	}

	value := inf(1)

	n2 := n.WithMove(m)
	n = nil
	for _, m := range AllMoves(&n2.board, -c) {
		v := e.negamax(n2, depth-1, -c, m) * -1
		value = min(value, v)
	}
	return value
}

func Heuristic3(n *Node, m Move) int {
	NumEvals++
	delta := -valueOf[n.board.At(m.Dst)+6]
	fast := (n.value + delta)
	return fast
}

func Heuristic2(b *Board, m Move) int {
	NumEvals++
	b = b.WithMove(m)
	h := 0
	for _, p := range b {
		h += valueOf[p+6]
	}
	return h
}

var valueOf = [13]int{
	WP + 6: 1,
	WN + 6: 6,
	WB + 6: 5,
	WR + 6: 10,
	WQ + 6: 20,
	BP + 6: -1,
	BN + 6: -6,
	BB + 6: -5,
	BR + 6: -10,
	BQ + 6: -20,
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func inf(c Color) int {
	return int(c) * 99999
}
