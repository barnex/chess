package picard

import (
	"math"
	"math/rand"
	"sort"

	"github.com/barnex/chess"
)

func New(depth int) chess.Engine {
	return &E{depth: depth, EnableRandom: true, Weight: [3]float64{0.001, 0.002, 0.001}}
}

func NewOpts(depth int, enableRandom bool) *E {
	return &E{depth: depth, EnableRandom: enableRandom}
}

type E struct {
	depth        int
	buffers      [][]chess.Move
	bufferN      [][]Node
	EnableRandom bool
	Weight       [3]float64
}

var (
	alphaCutoffs, betaCutoffs int
)

func (e *E) Move(b *chess.Board, c chess.Color) (chess.Move, float64) {
	alphaCutoffs = 0
	betaCutoffs = 0

	moves := chess.AllMoves(b, c)

	bm := chess.Move{}
	bv := -inf
	for _, m := range moves {
		b := b.WithMove(m)
		root := &Node{
			board: *b,
			value: MaterialValue(b),
		}
		_, v := e.AlphaBeta(root, -c, e.depth-1, bv, inf)
		v += rand.Float64() / 1e9
		v += e.Strategic(b)
		if v*float64(c) > bv {
			bm = m
			bv = v * float64(c)
		}
	}
	return bm, bv
}

func (e *E) AlphaBeta(n *Node, currPlayer chess.Color, depth int, alpha, beta float64) (chess.Move, float64) {
	if depth == 0 {
		return chess.Move{}, float64(n.value)
	}

	if n.KingTaken() {
		return chess.Move{}, float64(n.value)
	}

	allMoves := e.AllMoves(&n.board, currPlayer)
	defer e.Recycle(allMoves)

	children := e.BufferNodes()[:len(allMoves)]
	defer e.RecycleNodes(children)
	for i, m := range allMoves {
		n.WithMove(&children[i], m)
		chess.NumEvals++
	}

	bv := math.NaN()
	bm := chess.Move{}
	if currPlayer == chess.White {
		if depth > 1 {
			sort.Sort(ascending(children))
		}
		bv = -inf
		bm = chess.Move{}
		for _, c := range children {
			_, v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			//v += rand.Float64() / 1024
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
	} else {
		if depth > 1 {
			sort.Sort(descending(children))
		}
		bv = inf
		bm = chess.Move{}
		for _, c := range children {
			_, v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			//v += rand.Float64() / 1024
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
	}
	return bm, bv
}

type ascending []Node
type descending []Node

func (c ascending) Len() int            { return len(c) }
func (c descending) Len() int           { return len(c) }
func (c ascending) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c descending) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ascending) Less(i, j int) bool  { return c[i].value < c[j].value }
func (c descending) Less(i, j int) bool { return c[i].value > c[j].value }

var inf = math.Inf(1)

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
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
