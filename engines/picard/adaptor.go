package picard

import (
	"math"

	"github.com/barnex/chess"
)

type Adaptor struct {
	Valuator
}

type Valuator interface {
	ValueOf(b *chess.Board, nextMove chess.Color) int
}

func Adapt(v Valuator) chess.Engine {
	return &Adaptor{
		Valuator: v,
	}
}

func (a *Adaptor) Move(b *chess.Board, c chess.Color) (chess.Move, float64) {
	var (
		bestMove  = chess.Move{}
		bestScore = math.Inf(-1)
	)
	for _, m := range chess.AllMoves(b, c) {
		s := float64(a.ValueOf(b.WithMove(m), -c)) * float64(c) //  + rand.Float64()/1024
		//log.Println(m, s)
		if s > bestScore {
			bestScore = s
			bestMove = m
		}
	}
	return bestMove, bestScore
}
