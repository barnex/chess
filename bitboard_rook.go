package chess

import "fmt"

var luRook [8][256]Bits8

func (bb *BitBoard) SuperR(src Index, dst *[]Pos) {
	r, c := src.Row(), src.Col()
	pattern := bb.any.Row(r)
	movesPattern := luRook[c][pattern]

	for d := 0; d < 8; d++ {
		if movesPattern.Get(d) == 1 {
			*dst = append(*dst, RC(int(r), d))
		}
	}
}

func init() {
	var buf []Pos
	for rookCol := 0; rookCol < 8; rookCol++ {
		for pattern := 0; pattern < 256; pattern++ {

			var b Board
			for j := uint8(0); j < 8; j++ {
				if Bits(pattern).At(j) == 1 {
					b[j] = BP // any black piece would do
				}
			}
			b[rookCol] = WR

			buf = buf[:0]
			Moves(&b, RC(rookCol, 0).Index(), &buf)
			for _, m := range buf {
				luRook[rookCol][pattern].Set(m.Col())
			}
		}
	}
}

func copyIndex(buf []Pos) []Index {
	dsts := make([]Index, len(buf))
	for i := range buf {
		dsts[i] = buf[i].Index()
	}
	return dsts
}

type Bits8 uint8

func (b *Bits8) Set(i int) {
	if i < 0 || i > 7 {
		panic(fmt.Sprintf("bits8 index out of bounds: %v", i))
	}
	*b = *b | (0x1 << uint8(i))
}

func (b Bits8) Get(i int) uint8 {
	if i < 0 || i > 7 {
		panic(fmt.Sprintf("bits8 index out of bounds: %v", i))
	}
	return (uint8(b) >> uint8(i)) & 0x1
}
