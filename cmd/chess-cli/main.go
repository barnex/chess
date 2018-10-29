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
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	. "github.com/barnex/chess"
	"github.com/barnex/chess/engines/crusher"
	"github.com/barnex/chess/engines/picard"
	"github.com/barnex/chess/engines/riker"
	"github.com/barnex/chess/engines/tarr"
	"github.com/barnex/chess/engines/troi"
	"github.com/barnex/chess/engines/worf"
)

var (
	flagD = flag.Int("d", 2, "depth")
	flagE = flag.String("e", "crusher", "opponent: tarr|riker|worf|troi|picard|crusher")
	flagV = flag.Bool("v", false, "verbose output")
	flagB = flag.Bool("b", false, "play as black")
	flagL = flag.String("l", "chess.log", "log file")
)

var engines = map[string]func() Engine{
	"tarr":  func() Engine { return tarr.New(tarr.Heuristic2) },
	"riker": func() Engine { return riker.New(*flagD - 1) },
	"worf":  func() Engine { return worf.New(*flagD - 1) },
	"troi":  func() Engine { return troi.New(*flagD - 1) },
	"picard": func() Engine {
		e := picard.NewOpts(*flagD)
		e.Weight[0] = 0.001
		e.Weight[1] = 0.002
		return e
	},
	"crusher": func() Engine {
		e := crusher.New()
		e.Weight[0] = 0.001
		e.Weight[1] = 0.002
		return e
	},
}

var (
	b *Board
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	if *flagL != "" {
		f, err := os.Create(*flagL)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}
	log.Println("\nChess")

	ai, ok := engines[*flagE]
	if !ok {
		log.Fatalf("unknown engine: %v", *flagE)
	}

	var msg string
	var players [2]Engine
	if *flagB {
		players = [2]Engine{ai(), Stdin("black: ")}
		msg = fmt.Sprintf("White: %v%v\nBlack: you\n", *flagE, *flagD)
	} else {
		players = [2]Engine{Stdin("white: "), ai()}
		msg = fmt.Sprintf("White: you\nBlack: %v%v\n", *flagE, *flagD)
	}

	b = NewBoard()
	Render(b, map[Pos]bool{}, msg)

	colors := [2]Color{White, Black}
	names := []string{"white", "black"}
	current := 0
	moveNum := 0
	allCap := map[Color][]Piece{}
	for b.Winner() == Nobody {
		log.Println("\nmove:", moveNum)

		var buf bytes.Buffer // text beside the board
		printf := func(f string, x ...interface{}) {
			fmt.Fprintf(&buf, f, x...)
		}

		moveNum++
		NumEvals = 0

		// do the move
		start := time.Now()
		player := players[current]
		color := colors[current]
		m, score := player.Move(b, color)
		duration := time.Since(start)

		// print move
		printf(fgDark+"move %v\n%v: %v%v"+reset, moveNum, names[current], b.At(m.SrcI()), m)
		// print captured this move
		if x := b.At(m.DstI()); x != 00 {
			printf(" x%v", x)
			allCap[x.Color()] = append(allCap[x.Color()], x)
		}
		if x := IsCheck(b.WithMove(m)); x != 00 {
			printf("+ [CHECK!]")
		}
		printf("\n")

		// print duration
		printf(fgDark+"%v\n"+reset, duration.Round(time.Millisecond))

		// print stats
		{
			evals := float64(NumEvals)
			rate := evals / duration.Seconds()
			if *flagV {
				printf("score:%+.2f\nevals:%.3fM\nspeed:%.3fM/s\n",
					score, evals/1e6, rate/1e6)
			} else {
				printf("\n\n\n")
			}
		}

		// print all captured
		for _, c := range []Color{White, Black} {
			for _, x := range allCap[c] {
				printf(bgLight+fgBlack+"%v", x)
			}
			printf(reset + "\n")
		}

		b = b.WithMove(m)

		mark := map[Pos]bool{
			m.Src(): true,
			m.Dst(): true,
		}
		Render(b, mark, buf.String())

		current = 1 - current
	}

	fmt.Println(b.Winner(), "wins")
}
