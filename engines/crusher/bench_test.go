package crusher

import (
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

func BenchmarkNoAlphaBeta(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const d = 4
		const N = 20
		{
			ea := New(d, 0)
			eb := New(d, 0)

			game(ea, eb, N)
			b.SetBytes(int64(ea.evals + eb.evals))
		}
	}
}

func BenchmarkAlphaBeta(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const d = 4
		const N = 20
		{
			ea := New(d, 0)
			ea.EnableAlphaBeta = true
			eb := New(d, 0)
			eb.EnableAlphaBeta = true

			game(ea, eb, N)
			b.SetBytes(int64(ea.evals + eb.evals))
		}
	}
}

func BenchmarkAlphaBetaSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const d = 4
		const N = 20
		{
			ea := New(d, 0)
			ea.EnableAlphaBeta = true
			ea.EnableSort = true
			eb := New(d, 0)
			eb.EnableAlphaBeta = true
			eb.EnableSort = true

			game(ea, eb, N)
			b.SetBytes(int64(ea.evals + eb.evals))
		}
	}
}
