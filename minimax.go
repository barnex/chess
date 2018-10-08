package chess

import (
	"math/rand"
)

// Minimax thinks 1 move ahead.
func Minimax(h Heuristic) Engine {
	return &minimax{newRand(), h}
}

type minimax struct {
	rnd *rand.Rand
	h   Heuristic
}

func (e *minimax) Move(b *Board, c Color) Move {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = MinusInf
	)
	for _, m := range moves {
		b2 := b.WithMove(m)
		s := e.negamax(b2, 0, c)
		//fmt.Println("score ", i, s)
		if s.GT(bestScore) {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove
}

//function negamax(node, depth, color) is
//    if depth = 0 or node is a terminal node then
//        return color × the heuristic value of node
//    value := −∞
//    for each child of node do
//        value := max(value, −negamax(child, depth − 1, −color))
//    return value
func (e *minimax) negamax(b *Board, depth int, c Color) Value {
	if depth == 0 {
		return e.heuristicValue(b).Mul(c)
	}
	panic("TODO")
	//counterMoves := allMoves(b, -c)
	//value :=
}

func (e *minimax) heuristicValue(b *Board) Value {
	return Value{Win: b.Winner(), Heuristic: e.h(b)}
}

//func (e *minimax) heuristicMove(b *Board, m Move, c Color) Value {
//	b2 := b.Copy()
//	b2.Move(m.Src, m.Dst)
//	return e.heuristicBoard(b2, c)
//}
