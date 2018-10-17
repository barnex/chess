package chess

type BitBoard struct {
	any Bits // presence of any piece, black or white
}

func (bb *BitBoard) SetTo(b *Board) {
	var any uint64

	for i := 63; i >= 0; i-- {
		var v uint64
		if b[i] != 00 {
			v = 1
		}
		any = (any << 1) | v
	}
	bb.any = Bits(any)
}

func (bb *BitBoard) String() string {
	return bb.any.String()
}
