package main

import (
	"flag"
	"fmt"
	"time"

	. "github.com/barnex/chess"
	"github.com/barnex/chess/engines/riker"
)

var (
	flagD = flag.Int("d", 3, "depth")
)

func main() {
	flag.Parse()

	fmt.Println()

	b := NewBoard()

	Render(b)

	players := [2]Engine{Stdin("player: "), riker.New(*flagD)}
	colors := [2]Color{White, Black}
	current := 0
	for b.Winner() == Nobody {

		NumEvals = 0
		start := time.Now()

		player := players[current]
		color := colors[current]
		m, score := player.Move(b, color)

		duration := time.Since(start)
		evals := float64(NumEvals)
		rate := evals / duration.Seconds()
		src := b.At(m.Src)
		fmt.Printf("%v%v [%+.2f] %.3fM evals in %v = %.3fM/s \n", src, m, score, evals/1e6, duration.Round(time.Millisecond), rate/1e6)

		b = b.WithMove(m)
		Render(b)

		current = 1 - current
	}

	fmt.Println(b.Winner(), "wins")
}
