package main

import (
	"flag"
	"fmt"

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

		m, score := players[current].Move(b, colors[current])

		src := b.At(m.Src)
		fmt.Printf("%v%v %.3f\n", src, m, score)
		b = b.WithMove(m)
		Render(b)

		current = 1 - current
	}

	fmt.Println(b.Winner(), "wins")
}
