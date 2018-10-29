package crusher

import (
	"log"

	"github.com/barnex/chess"
)

func (e *E) Strategic(b *chess.Board) float64 {

	mobility := 0.0
	protection := 0.0

	allW := chess.AllPre(b, chess.White)
	allB := chess.AllPre(b, chess.Black)
	all := append(allW, allB...)

	for _, m := range all {

		srci := m.SrcI()
		dsti := m.DstI()
		src := b.At(srci)
		dst := b.At(dsti)
		srcC := src.Color()
		dstC := dst.Color()

		if dst == 00 {
			mobility += float64(srcC) * posVal(srcC)[dsti]
		}

		if dstC == srcC && abs(ValueOf(dst)) < 10 {
			protection += float64(srcC) * posVal(srcC)[dsti]
		}
	}

	log.Println("mobility:", mobility, "protection:", protection)
	return mobility + protection*20
}

func posVal(c chess.Color) *[64]float64 {
	if c == chess.White {
		return &posValW
	} else {
		return &posValB
	}
}

var posValW = [64]float64{
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 6, 6, 2, 2, 2,
	3, 3, 3, 6, 6, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
}

var posValB = [64]float64{
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 6, 6, 3, 3, 3,
	2, 2, 2, 6, 6, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
