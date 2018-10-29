package picard

import (
	"math"
	"math/rand"
	"sort"

	"github.com/barnex/chess"
)

func New(depth int) chess.Engine {
	return &E{depth: depth, EnableRandom: true, Weight: [4]float64{0.001, 0.005, 0.0, 0.0005}}
}

func NewOpts(depth int) *E {
	return &E{depth: depth, EnableRandom: true}
}

type E struct {
	depth        int
	buffers      [][]chess.Move
	bufferN      [][]Node
	EnableRandom bool
	Weight       [4]float64
}

var (
	alphaCutoffs, betaCutoffs int
)

func (e *E) Move(b *chess.Board, c chess.Color) (chess.Move, float64) {
	alphaCutoffs = 0
	betaCutoffs = 0

	moves := chess.AllMoves(b, c)

	bm := chess.Move{}
	bv := -inf
	for _, m := range moves {
		b := b.WithMove(m)
		root := &Node{
			board: *b,
			value: MaterialValue(b),
		}
		_, v := e.AlphaBeta(root, -c, e.depth-1, bv, inf)
		v += rand.Float64() / (1024 * 1024 * 1024)
		if v*float64(c) > bv {
			bm = m
			bv = v * float64(c)
		}
	}
	return bm, bv
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
		chess.NumEvals++
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
	//nonForced := 0
	for _, ch := range children {
		_, v := e.AlphaBeta(&ch, -currPlayer, depth-1, alpha, beta)

		if depth == e.depth-1 && e.Weight != ([4]float64{}) {
			v += e.Strategic(&ch.board)
		}

		if v*c > bv*c {
			bv = v
			bm = ch.move
		}

		//if depth == 1 {
		//	if v*c >= c*float64(n.value) {
		//		nonForced++
		//	}
		//}

		alpha = c * max(c*alpha, c*bv)
		if alpha >= beta {
			alphaCutoffs++
			break
		}
	}

	//if nonForced != 0 {
	//	//fmt.Println(nonForced)
	//	bv += float64(currPlayer) * float64(nonForced) * e.Weight[3]
	//	//bv += float64(nonForced) / 1024
	//}

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
