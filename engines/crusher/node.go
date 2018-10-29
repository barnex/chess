package crusher

import "github.com/barnex/chess"

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
