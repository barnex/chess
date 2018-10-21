package picard

import (
	"math"

	"github.com/barnex/chess"
)

func New(depth int) chess.Engine {
	return &E{depth: depth}
}

type E struct {
	depth   int
	buffers [][]chess.Move
	bufferN [][]Node
}

func (e *E) Move(b *chess.Board, c chess.Color) (chess.Move, float64) {
	root := &Node{
		board: *b,
		value: MaterialValue(b),
	}
	m, s := e.AlphaBeta(root, c, e.depth, -inf, inf)
	return m, float64(s)
}

// TODO: also return move
func (e *E) AlphaBeta(n *Node, currPlayer chess.Color, depth int, alpha, beta int) (chess.Move, int) {
	if depth == 0 {
		return chess.Move{}, n.value
	}

	if n.KingTaken() {
		return chess.Move{}, n.value
	}

	allMoves := e.AllMoves(&n.board, currPlayer)
	defer e.Recycle(allMoves)
	//log.Println("allmoves", currPlayer, allMoves)

	children := e.BufferNodes()[:len(allMoves)]
	defer e.RecycleNodes(children)
	for i, m := range allMoves {
		n.WithMove(&children[i], m)
		chess.NumEvals++
	}

	if currPlayer == chess.White {
		bv := -inf
		bm := chess.Move{}
		for i, c := range children {
			_, v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			if v > bv {
				bv = v
				bm = allMoves[i]
			}
			alpha = max(alpha, bv)
			if alpha >= beta {
				break
			}

		}
		return bm, bv
	} else {
		bv := inf
		bm := chess.Move{}
		for i, c := range children {
			_, v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			if v < bv {
				bv = v
				bm = allMoves[i]
			}
			beta = min(beta, bv)
			if alpha >= beta {
				break
			}
		}
		return bm, bv
	}
}

const inf = math.MaxInt64

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (e *E) AllMoves(b *chess.Board, c chess.Color) []chess.Move {
	moves := e.Buffer()
	for i := range b {
		if b.At(chess.Index(i)).Color() == c {
			chess.Moves(b, chess.Index(i), &moves)
		}
	}
	return moves
}
