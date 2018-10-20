package picard

import "github.com/barnex/chess"

func New(depth int) chess.Engine {
	return &Adaptor{&E{depth: depth}}
}

type E struct {
	depth   int
	buffers [][]chess.Move
	bufferN [][]Node
}

func (e *E) ValueOf(b *chess.Board, nextPlayer chess.Color) int {
	root := &Node{
		board: *b,
		value: MaterialValue(b),
	}

	return e.AlphaBeta(root, -nextPlayer, e.depth)
}

func (e *E) AlphaBeta(n *Node, currPlayer chess.Color, depth int) int {
	if depth == 0 {
		return n.value
	}

	//allMoves := e.AllMoves(n, currPlayer)
	//defer e.Recycle(allMoves)
	//children := e.BufferNodes()
	//defer e.RecycleNodes(children)
	panic("TODO")
}

func (e *E) Buffer() []chess.Move {
	if len(e.buffers) > 0 {
		b := e.buffers[len(e.buffers)-1]
		b = b[:0]
		e.buffers = e.buffers[:len(e.buffers)-1]
		return b
	}
	return make([]chess.Move, 0, 64)
}

func (e *E) Recycle(b []chess.Move) {
	b = b[:0]
	e.buffers = append(e.buffers, b)
}

func (e *E) BufferNodes() []Node {
	if len(e.bufferN) > 0 {
		b := e.bufferN[len(e.bufferN)-1]
		b = b[:0]
		e.bufferN = e.bufferN[:len(e.bufferN)-1]
		return b
	}
	return make([]Node, 0, 64)
}

func (e *E) RecycleNode(b []Node) {
	b = b[:0]
	e.bufferN = append(e.bufferN, b)
}
