/*
Command chess-cli provides a chess command line interface.

It requires a utf-8 capable, 256 color terminal emulator
like most modern Linux terminals.

The player types moves as algebraic notation:
https://en.wikipedia.org/wiki/Algebraic_notation_(chess)

E.g.:
	a2a3    a2 to a3
	a3      pawn to a3, if unambiguous
	Ng3     knight to g3, if unambiguous
*/
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
	flagD = flag.Int("d", 4, "depth")
	flagE = flag.String("e", "worf", "opponent: tarr|riker|worf")
	flagV = flag.Bool("v", false, "verbose output")
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

	Render(b, map[Pos]bool{})

	players := [2]Engine{ai(), Stdin("player: ")}
	colors := [2]Color{White, Black}
	names := []string{"white", "black"}
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

		if *flagV {
			fmt.Printf("score:%+.2f, evals:%.3fM, time:%v, speed:%.3fM/s\n",
				score, evals/1e6, duration.Round(time.Millisecond), rate/1e6)
		}

		fmt.Printf("\n%v: %v%v", names[current], src, m)
		if x := b.At(m.Dst); x != 00 {
			fmt.Printf(" x%v", x)
		}
		if x := IsCheck(b.WithMove(m)); x != 00 {
			fmt.Printf("+ [CHECK!]")
		}
		fmt.Println()

		b = b.WithMove(m)
		mark := map[Pos]bool{
			m.Src: true,
			m.Dst: true,
		}
		Render(b, mark)

		current = 1 - current
	}

	fmt.Println(b.Winner(), "wins")
}
