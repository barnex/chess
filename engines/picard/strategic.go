package picard

import (
	"github.com/barnex/chess"
)

func (e *E) Strategic(b *chess.Board) float64 {

	mobility := 0.0
	protection := 0.0
	threat := 0.0
	//fork := 0.0
	//space := 0.0
	//center := 0.0

	allW := chess.AllPre(b, chess.White)
	allB := chess.AllPre(b, chess.Black)
	all := append(allW, allB...)

	for _, m := range all {
		src := b.At(m.SrcI())
		dst := b.At(m.DstI())
		srcC := src.Color()
		dstC := dst.Color()
		if dst == 00 {
			mobility += float64(srcC)
		}
		if dstC == srcC && abs(ValueOf(dst)) < 10 {
			protection += float64(srcC)
		}
		if dstC == -srcC {
			threat -= float64(ValueOf(dst))
		}
	}

	return mobility*e.Weight[0] + protection*e.Weight[1] + threat*e.Weight[2]
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
