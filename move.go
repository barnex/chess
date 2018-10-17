package chess

type Move struct {
	src, dst Index
}

func MoveP(src, dst Pos) Move {
	return Move{src.Index(), dst.Index()}
}

func (m Move) Src() Pos    { return m.src.Pos() }
func (m Move) Dst() Pos    { return m.dst.Pos() }
func (m Move) SrcI() Index { return m.src }
func (m Move) DstI() Index { return m.dst }

func (m Move) String() string {
	return m.Src().String() + m.Dst().String()
}
