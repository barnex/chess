package chess

type Move struct {
	Src, Dst Pos
}

func (m Move) String() string {
	return m.Src.String() + m.Dst.String()
}
