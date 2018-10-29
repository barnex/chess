package crusher

import (
	"math"
	"sort"

	"github.com/barnex/chess"
)

// Enginelet is an alpha-beta worker with its own storage pool.
// One enginelet per thread.
type Enginelet struct {
	e                                *E
	evals, alphaCutoffs, betaCutoffs int
	bufferMoves                      []chess.Move
	bufferN                          [][]Node
}

func (e *Enginelet) AlphaBeta(n *Node, currPlayer chess.Color, depth int, alpha, beta float64) float64 {

	// reached end of recursion
	if depth == 0 {
		e.evals++
		return float64(n.value)
	}

	// If the king has been captured, the rest of the tree must not
	// be evaluated further. Otherwise, we might capture the opposing
	// king back and consider this a zero-value move.
	if n.KingTaken() {
		return float64(n.value)
	}

	children := e.AllChildren(n, currPlayer)
	defer e.Recycle(children)

	// Sort the most promising moves first,
	// to get more alpha-beta cut-offs.
	// But only do so near the top of the tree.
	// Benchmarks show the sorting cost is not
	// payed back near the bottom of the tree.
	if depth > 1 && e.e.EnableSort {
		if currPlayer == chess.White {
			sort.Sort(descending(children))
		} else {
			sort.Sort(ascending(children))
		}
	}

	bv := math.NaN()
	if currPlayer == chess.White {
		bv = -inf
		for _, c := range children {
			v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			bv = max(bv, v)
			alpha = max(alpha, bv)
			if alpha >= beta && e.e.EnableAlphaBeta {
				e.betaCutoffs++
				break
			}

		}
	} else {
		bv = inf
		for _, c := range children {
			v := e.AlphaBeta(&c, -currPlayer, depth-1, alpha, beta)
			bv = min(bv, v)
			beta = min(beta, bv)
			if alpha >= beta && e.e.EnableAlphaBeta {
				e.alphaCutoffs++
				break
			}
		}
	}
	return bv
}

// Construct all possible moves
// TODO: place capturing moves first
// TODO: could construct not all at once, benefit more from cut-offs.
func (e *Enginelet) AllChildren(n *Node, c chess.Color) []Node {

	if e.bufferMoves == nil {
		e.bufferMoves = make([]chess.Move, 128)
	}
	allMoves := e.bufferMoves[:0]

	b := &n.board
	for i := range b {
		if b.At(chess.Index(i)).Color() == c {
			chess.Moves(b, chess.Index(i), &allMoves)
		}
	}

	children := e.BufferNodes()[:len(allMoves)]
	for i, m := range allMoves {
		n.WithMove(&children[i], m)
	}
	return children
}

type Node struct {
	board chess.Board
	value int
	move  chess.Move
}

func (n *Node) WithMove(dst *Node, m chess.Move) {
	dst.value = n.value - ValueOf(n.board[m.DstI()])

	dst.board = n.board
	dst.board[m.SrcI()] = 00
	dst.board[m.DstI()] = n.board[m.SrcI()]
	dst.move = m
}

type ascending []Node
type descending []Node

func (c ascending) Len() int            { return len(c) }
func (c descending) Len() int           { return len(c) }
func (c ascending) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c descending) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ascending) Less(i, j int) bool  { return c[i].value < c[j].value }
func (c descending) Less(i, j int) bool { return c[i].value > c[j].value }

func (e *Enginelet) BufferNodes() []Node {
	if len(e.bufferN) > 0 {
		b := e.bufferN[len(e.bufferN)-1]
		b = b[:0]
		e.bufferN = e.bufferN[:len(e.bufferN)-1]
		return b
	}
	return make([]Node, 0, 128)
}

func (e *Enginelet) Recycle(b []Node) {
	b = b[:0]
	e.bufferN = append(e.bufferN, b)
}
