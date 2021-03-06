package crusher

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"sync"

	"github.com/barnex/chess"
)

func New(depth1, depth2 int) *E {
	return &E{
		depth1:          depth1,
		depth2:          depth2,
		EnableAlphaBeta: true,
		EnableSort:      true,
		EnableStrategy:  true,
		CapturePenalty:  0.5,
		WMobility:       1,
		WProtection:     1,
		EnableRandom:    true,
	}
}

type E struct {
	depth1, depth2                   int
	EnableAlphaBeta                  bool
	EnableSort                       bool
	EnableRandom                     bool
	EnableStrategy                   bool
	CapturePenalty                   float64
	WMobility                        float64
	WProtection                      float64
	evals, alphaCutoffs, betaCutoffs int
}

func (e *E) Move(b *chess.Board, c chess.Color) (chess.Move, float64) {
	moves := chess.AllMoves(b, c)
	mv := make([]node, len(moves))

	el := &Enginelet{e: e}

	bv := -inf
	if c == chess.Black {
		bv = inf
	}
	for i, m := range moves {
		alpha := -inf
		beta := inf
		b := b.WithMove(m)
		root := &Node{board: *b, value: MaterialValue(b)}
		v := el.AlphaBeta(root, -c, e.depth1-1, alpha, beta)
		if c == chess.White {
			bv = max(bv, v)
			alpha = max(alpha, bv)
		} else {
			bv = min(bv, v)
			beta = min(beta, bv)
		}
		//if e.EnableStrategy {
		//	v += e.Strategic(b) / (1000)
		//}
		mv[i] = node{m, v}
	}

	e.evals += el.evals
	e.alphaCutoffs += el.alphaCutoffs
	e.betaCutoffs += el.betaCutoffs

	if c == chess.White {
		sort.Sort(desc(mv))
	} else {
		sort.Sort(asc(mv))
	}
	log.Println(c, "evals:", el.evals, "alpha cutoffs:", el.alphaCutoffs, "beta cutoffs:", el.betaCutoffs)

	if e.depth2 > 0 {
		log.Print("refining...")

		N := runtime.NumCPU()
		if N > len(mv) {
			N = len(mv)
		}

		var wg sync.WaitGroup
		for i := 0; i < N; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				el := &Enginelet{e: e}
				b := b.WithMove(mv[i].Move)
				root := &Node{board: *b, value: MaterialValue(b)}
				v := el.AlphaBeta(root, -c, e.depth2-1, -inf, inf)
				if el.e.EnableStrategy {
					v += e.Strategic(b) / (1000)
				}
				mv[i].Value = v
				log.Println(c, "evals:", el.evals, "alpha cutoffs:", el.alphaCutoffs, "beta cutoffs:", el.betaCutoffs)
			}(i)
		}

		wg.Wait()

		for i := range mv {
			if isCapture(b, mv[i].Move) {
				mv[i].Value -= e.CapturePenalty * float64(c)
			}
		}

		if c == chess.White {
			sort.Sort(desc(mv))
		} else {
			sort.Sort(asc(mv))
		}
		log.Print(mv[:N])
	}

	sel := 0
	if e.EnableRandom && rand.Intn(3) == 0 && math.Abs(mv[0].Value-mv[1].Value) < 0.5 {
		log.Println("randomly selecting", mv[1], "over", mv[0])
		sel = 1
	}
	return mv[sel].Move, mv[sel].Value
}

func isCapture(b *chess.Board, m chess.Move) bool {
	return b.At(m.DstI()) != 00
}

type desc []node

func (c desc) Len() int           { return len(c) }
func (c desc) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c desc) Less(i, j int) bool { return c[i].Value > c[j].Value }

type asc []node

func (c asc) Len() int           { return len(c) }
func (c asc) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c asc) Less(i, j int) bool { return c[i].Value < c[j].Value }

type node struct {
	chess.Move
	Value float64
}

func (n node) String() string {
	return fmt.Sprintf("%v=%.5f", n.Move, n.Value)
}

var inf = math.Inf(1)

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func cmax(c chess.Color, a, b float64) float64 {
	if c == chess.White {
		return max(a, b)
	} else {
		return min(a, b)
	}
}
