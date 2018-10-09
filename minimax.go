package chess

import (
	"math/rand"
)

func Minimax(depth int, h Heuristic) Engine {
	return &minimax{newRand(), depth, h}
}

type minimax struct {
	rnd   *rand.Rand
	depth int
	h     Heuristic
}

func (e *minimax) Move(b *Board, c Color) Move {
	moves := allMoves(b, c)

	var (
		bestMove  = moves[e.rnd.Intn(len(moves))]
		bestScore = MinusInf
	)
	for _, m := range moves {
		b2 := b.WithMove(m)
		s := e.negamax(b2, e.depth, c, true)
		//fmt.Println("score ", i, s)
		if s.GT(bestScore) {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove
}

//function minimax(node, depth, maximizingPlayer) is
// if depth = 0 or node is a terminal node then
//     return the heuristic value of node
// if maximizingPlayer then
//     value := −∞
//     for each child of node do
//         value := max(value, minimax(child, depth − 1, FALSE))
//     return value
// else (* minimizing player *)
//     value := +∞
//     for each child of node do
//         value := min(value, minimax(child, depth − 1, TRUE))
//     return value
func (e *minimax) negamax(b *Board, depth int, c Color, max bool) Value {
	if depth == 0 {
		return e.heuristicValue(b).Mul(c)
	}

	counterMoves := allMoves(b, -c)

	if max {
		value := MinusInf.Mul(-1)
		for _, m := range counterMoves {
			b2 := b.WithMove(m)
			v := e.negamax(b2, depth-1, -c, !max).Mul(-1)
			value = Min(value, v)
		}
		return value
	} else {
		value := MinusInf
		for _, m := range counterMoves {
			b2 := b.WithMove(m)
			v := e.negamax(b2, depth-1, -c, !max).Mul(-1)
			value = Max(value, v)
		}
		return value
	}
}

func (e *minimax) heuristicValue(b *Board) Value {
	return Value{Win: b.Winner(), Heuristic: e.h(b)}
}

//func (e *minimax) heuristicMove(b *Board, m Move, c Color) Value {
//	b2 := b.Copy()
//	b2.Move(m.Src, m.Dst)
//	return e.heuristicBoard(b2, c)
//}
