package crusher

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"sort"
	"sync"

	"github.com/barnex/chess"
)

func New() *E {
	return &E{depth: 4, EnableRandom: true, Weight: [2]float64{0.001, 0.002}}
}

type E struct {
	depth        int
	buffers      [][]chess.Move
	bufferN      [][]Node
	EnableRandom bool
	Weight       [2]float64
}

var (
	alphaCutoffs, betaCutoffs int
)

func (e *E) Move(b *chess.Board, c chess.Color) (chess.Move, float64) {
	moves := chess.AllMoves(b, c)
	mv := make([]node, len(moves))

	for i, m := range moves {
		b := b.WithMove(m)
		root := &Node{board: *b, value: MaterialValue(b)}
		_, v := e.AlphaBeta(root, -c, 3, -inf, inf)
		mv[i] = node{m, v}
	}

	if c == chess.White {
		sort.Sort(desc(mv))
	} else {
		sort.Sort(asc(mv))
	}

	log.Print(mv)
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
			e1 := e
			e := New()
			e.Weight = e1.Weight
			e.depth = e1.depth
			b := b.WithMove(mv[i].Move)
			root := &Node{board: *b, value: MaterialValue(b)}
			_, v := e.AlphaBeta(root, -c, 5, -inf, inf)
			mv[i].Value = v
		}(i)
	}

	wg.Wait()

	if c == chess.White {
		sort.Sort(desc(mv))
	} else {
		sort.Sort(asc(mv))
	}

	log.Print(mv[:N])

	return mv[0].Move, mv[0].Value
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

func (e *E) AlphaBeta(n *Node, currPlayer chess.Color, depth int, alpha, beta float64) (chess.Move, float64) {

	// reached end of recursion
	if depth == 0 {
		return chess.Move{}, float64(n.value)
	}

	// If the king has been captured, the rest of the tree must not
	// be evaluated further. Otherwise, we might capture the opposing
	// king back and consider this a zero-value move.
	if n.KingTaken() {
		return chess.Move{}, float64(n.value)
	}

	// Construct all possible moves
	allMoves := e.AllMoves(&n.board, currPlayer)
	defer e.Recycle(allMoves)
	children := e.BufferNodes()[:len(allMoves)]
	defer e.RecycleNodes(children)
	for i, m := range allMoves {
		n.WithMove(&children[i], m)
	}

	// Sort the most promising moves first,
	// to get more alpha-beta cut-offs.
	// But only do so near the top of the tree.
	// Benchmarks show the sorting cost is not
	// payed back near the bottom of the tree.
	if depth > 1 {
		if currPlayer == chess.White {
			sort.Sort(ascending(children))
		} else {
			sort.Sort(descending(children))
		}
	}

	// Negamax with alpha-beta pruning
	bv := math.NaN()
	bm := chess.Move{}

	c := float64(currPlayer)
	bv = -inf * c
	bm = chess.Move{}
	for _, ch := range children {
		_, v := e.AlphaBeta(&ch, -currPlayer, depth-1, alpha, beta)

		if depth == e.depth-1 {
			v += e.Strategic(&ch.board)
		}

		if v*c > bv*c {
			bv = v
			bm = ch.move
		}

		alpha = c * max(c*alpha, c*bv)
		if alpha >= beta {
			alphaCutoffs++
			break
		}
	}

	return bm, bv
}

type ascending []Node
type descending []Node

func (c ascending) Len() int            { return len(c) }
func (c descending) Len() int           { return len(c) }
func (c ascending) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c descending) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ascending) Less(i, j int) bool  { return c[i].value < c[j].value }
func (c descending) Less(i, j int) bool { return c[i].value > c[j].value }

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

func (e *E) AllMoves(b *chess.Board, c chess.Color) []chess.Move {
	moves := e.Buffer()
	for i := range b {
		if b.At(chess.Index(i)).Color() == c {
			chess.Moves(b, chess.Index(i), &moves)
		}
	}
	return moves
}
