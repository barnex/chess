package chess

type Index uint8

func (i Index) Pos() Pos {
	return Pos{int8(i.Row()), int8(i.Col())}
}

func (i Index) Row() uint8 {
	return uint8(i >> 3)
}

func (i Index) Col() uint8 {
	return uint8(i & 0x7)
}
