package crusher

import (
	"log"
	"testing"

	"github.com/barnex/chess"
	"github.com/barnex/chess/engines/random"
)

func init() {
	log.SetFlags(0)
	//log.SetOutput(ioutil.Discard)
}

func BenchmarkDepthGame4(b *testing.B) {
	e := New(4, 0)
	r := random.New()
	for i := 0; i < b.N; i++ {
		b := chess.NewBoard()
		for i := 0; i < 30; i++ {
			e.Move(b, chess.White)
			r.Move(b, chess.Black)
		}
	}
}
