package picard

import "github.com/barnex/chess"

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
