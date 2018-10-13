package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	. "github.com/barnex/chess"
	"github.com/barnex/chess/engines/riker"
	"github.com/barnex/chess/engines/tarr"
	"github.com/barnex/chess/engines/worf"
)

var (
	flagD = flag.Int("d", 3, "depth")
	flagE = flag.String("e", "worf", "opponent: tarr|riker|worf")
)

var engines = map[string]func() Engine{
	"tarr":  func() Engine { return tarr.New(tarr.Heuristic2) },
	"riker": func() Engine { return riker.New(*flagD) },
	"worf":  func() Engine { return worf.New(*flagD) },
}

func main() {
	flag.Parse()

	ai, ok := engines[*flagE]
	if !ok {
		log.Fatalf("unknown engine: %v", *flagE)
	}

	fmt.Println()

	b := NewBoard()

	Render(b)

	players := [2]Engine{Stdin("player: "), ai()}
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
