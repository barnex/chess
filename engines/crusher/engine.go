package crusher

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"

	"github.com/barnex/chess"
)

func New(depth1, depth2 int) *E {
	return &E{
		depth1: depth1,
		depth2: depth2,
		//EnableRandom: true,
		//	Weight: [2]float64{0.001, 0.002},
	}
}

type E struct {
	depth1, depth2  int
	EnableAlphaBeta bool
	EnableSort      bool
	EnableRandom    bool
	Weight          [2]float64
	CapturePenalty  float64
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
		if e.EnableRandom {
			v += rand.Float64() / (1e9)
		}
		mv[i] = node{m, v}
	}

	if c == chess.White {
		sort.Sort(desc(mv))
	} else {
		sort.Sort(asc(mv))
	}
	log.Print(c, el.evals, el.alphaCutoffs, el.betaCutoffs)

	//log.Print("refining...")

	//N := runtime.NumCPU()
	//if N > len(mv) {
	//	N = len(mv)
	//}

	//var wg sync.WaitGroup
	//for i := 0; i < N; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		e1 := e
	//		e := New()
	//		e.Weight = e1.Weight
	//		e.depth = e1.depth
	//		b := b.WithMove(mv[i].Move)
	//		root := &Node{board: *b, value: MaterialValue(b)}
	//		_, v := e.AlphaBeta(root, -c, 5, -inf, inf)
	//		mv[i].Value = v
	//	}(i)
	//}

	//wg.Wait()

	//for i := range mv {
	//	if isCapture(b, mv[i].Move) {
	//		mv[i].Value -= e.CapturePenalty * float64(c)
	//	}
	//}

	//if c == chess.White {
	//	sort.Sort(desc(mv))
	//} else {
	//	sort.Sort(asc(mv))
	//}

	//log.Print(mv[:N])

	return mv[0].Move, mv[0].Value
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
