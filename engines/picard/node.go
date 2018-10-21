package picard

import "github.com/barnex/chess"

type Node struct {
	board chess.Board
	value int
}

func (n *Node) WithMove(dst *Node, m chess.Move) {
	dst.value = n.value - ValueOf(n.board[m.DstI()])

	dst.board = n.board
	dst.board[m.SrcI()] = 00
	dst.board[m.DstI()] = n.board[m.SrcI()]
}

func (n *Node) KingTaken() bool {
	return n.value > 500 || n.value < -500
}
