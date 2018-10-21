package picard

import (
	"math"
	"sort"

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

var (
	alphaCutoffs, betaCutoffs int
)

func (e *E) Move(b *chess.Board, c chess.Color) (chess.Move, float64) {
	alphaCutoffs = 0
	betaCutoffs = 0

	root := &Node{
		board: *b,
		value: MaterialValue(b),
	}
	m, s := e.AlphaBeta(root, c, e.depth, -inf, inf)
	//log.Println("alpha cuttoffs:", alphaCutoffs)
	//log.Println("beta cuttoffs:", betaCutoffs)
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
		sort.Sort(descending(children))
		bv := -inf
		bm := chess.Move{}
		for _, c := range children {
			_, v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			if v > bv {
				bv = v
				bm = c.move
			}
			alpha = max(alpha, bv)
			if alpha >= beta {
				alphaCutoffs++
				break
			}

		}
		return bm, bv
	} else {
		sort.Sort(ascending(children))
		bv := inf
		bm := chess.Move{}
		for _, c := range children {
			_, v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			if v < bv {
				bv = v
				bm = c.move
			}
			beta = min(beta, bv)
			if alpha >= beta {
				betaCutoffs++
				break
			}
		}
		return bm, bv
	}
}

type ascending []Node
type descending []Node

func (c ascending) Len() int            { return len(c) }
func (c descending) Len() int           { return len(c) }
func (c ascending) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c descending) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ascending) Less(i, j int) bool  { return c[i].value < c[j].value }
func (c descending) Less(i, j int) bool { return c[i].value > c[j].value }

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
