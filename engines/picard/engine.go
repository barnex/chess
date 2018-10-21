package picard

import (
	"math"

	"github.com/barnex/chess"
)

func New(depth int) chess.Engine {
	return &Adaptor{&E{depth: depth}}
}

type E struct {
	depth   int
	buffers [][]chess.Move
	bufferN [][]Node
}

// TODO: also return move
func (e *E) ValueOf(b *chess.Board, nextPlayer chess.Color) int {
	root := &Node{
		board: *b,
		value: MaterialValue(b),
	}
	return e.AlphaBeta(root, nextPlayer, e.depth)
}

// TODO: also return move
func (e *E) AlphaBeta(n *Node, currPlayer chess.Color, depth int) int {
	if depth == 0 {
		return n.value
	}

	if n.KingTaken() {
		return n.value
	}

	allMoves := e.AllMoves(&n.board, currPlayer)
	defer e.Recycle(allMoves)
	//log.Println("allmoves", currPlayer, allMoves)

	children := e.BufferNodes()[:len(allMoves)]
	defer e.RecycleNodes(children)
	for i, m := range allMoves {
		n.WithMove(&children[i], m)
	}

	if currPlayer == chess.White {
		v := -inf
		for _, c := range children {
			v = max(v, e.AlphaBeta(&c, -currPlayer, depth-1))
		}
		return v
	} else {
		v := inf
		for _, c := range children {
			v = min(v, e.AlphaBeta(&c, -currPlayer, depth-1))
		}
		return v
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
