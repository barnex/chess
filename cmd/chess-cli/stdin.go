package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/barnex/chess"
)

// Stdin returns an engine that reads moves from standard input.
func Stdin(prompt string) Engine {
	return &stdin{prompt, bufio.NewScanner(os.Stdin)}
}

type stdin struct {
	prompt  string
	scanner *bufio.Scanner
}

func (e *stdin) Move(b *Board, c Color) (Move, float64) {
	move, err := Parse(e.ReadLine(), b, c)
	for err != nil {
		fmt.Println(err)
		move, err = Parse(e.ReadLine(), b, c)
	}
	return move, 0
}

func Parse(line string, b *Board, c Color) (Move, error) {
	line = strings.TrimSpace(line)

	switch len(line) {
	default:
		return Move{}, fmt.Errorf("syntax error: %q", line)
	case 4:
		return Parse4(line)
	case 2:
		return Parse2(line, b, c)
	}
}

// Parse2 parses a 2-character move, like
// 	a3
func Parse2(line string, b *Board, c Color) (Move, error) {
	dst, err := ParsePos(line)
	if err != nil {
		return Move{}, err
	}

	myPawn := WP * Piece(c)
	var cand []Move
	for _, a := range AllMoves(b, c) {
		if b.At(a.Src) == myPawn && a.Dst == dst {
			cand = append(cand, a)
		}
	}
	switch len(cand) {
	case 0:
		return Move{}, fmt.Errorf("%v not allowed", line)
	case 1:
		return cand[0], nil
	default:
		return Move{}, fmt.Errorf("%v is ambigous: %v match", line, cand)
	}

}

// Parse4 parses a 4-character move, like
// 	a1b2
func Parse4(line string) (Move, error) {
	src, err := ParsePos(line[:2])
	if err != nil {
		return Move{}, err
	}

	dst, err := ParsePos(line[2:4])
	if err != nil {
		return Move{}, err
	}

	m := Move{src, dst}
	return m, nil
}

func ParsePos(txt string) (Pos, error) {
	if len(txt) != 2 {
		return Pos{}, fmt.Errorf("syntax error")
	}
	c := int(txt[0]) - int('a')
	r := int(txt[1]) - int('1')

	if r < 0 || r > 7 || c < 0 || c > 7 {
		return Pos{}, fmt.Errorf("syntax error: %q", txt)
	}
	return RC(r, c), nil
}

func (e *stdin) ReadLine() string {
	fmt.Print(e.prompt)
	if e.scanner.Scan() == false {
		os.Exit(0) // end of stream
	}
	if e.scanner.Err() != nil {
		log.Fatal(e.scanner.Err())
	}
	return e.scanner.Text()
}
