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
	return &stdin{prompt: prompt, scanner: bufio.NewScanner(os.Stdin)}
}

type stdin struct {
	prompt   string
	scanner  *bufio.Scanner
	previous Board
	inited   bool
}

func (e *stdin) Move(_ *Board, c Color) (Move, float64) {
	if !e.inited {
		e.inited = true
		e.previous = *b
	}

	move, err := e.Parse(e.ReadLine(), c)
	for err != nil {
		fmt.Println(err)
		move, err = e.Parse(e.ReadLine(), c)
	}
	e.previous = *b
	return move, 0
}

func (e *stdin) Parse(line string, c Color) (Move, error) {
	if line == "undo" {
		cpy := e.previous
		b = &cpy
		Render(b, nil, "undone move")
		return Move{}, fmt.Errorf("undone")
	}
	return ParseAndValidate(line, b, c)
}

func ParseAndValidate(line string, b *Board, c Color) (Move, error) {
	move, err := Parse(line, b, c)
	if err != nil {
		return Move{}, err
	}
	if !Allowed(b, c, move) {
		return Move{}, illegalMove(move)
	}
	return move, nil
}

func Parse(line string, b *Board, c Color) (Move, error) {
	line = strings.TrimSpace(line)

	switch len(line) {
	default:
		return Move{}, syntaxError(line)
	case 2:
		return Parse2(line, b, c)
	case 3:
		return Parse3(line, b, c)
	case 4:
		return Parse4(line)
	}
}

// Parse2 parses a 2-character move, like
// 	a3
// which means to move a pawn to a3.
// Returns an error if the move is ambiguous or not allowed.
func Parse2(line string, b *Board, c Color) (Move, error) {
	return Parse3("P"+line, b, c)
}

// Parse3 parses a 3-character move, like
// 	Nf3
// which means to move a knight to f3.
// Returns an error if the move is ambiguous or not allowed.
func Parse3(line string, b *Board, c Color) (Move, error) {
	p, err := ParsePiece(line[:1])
	if err != nil {
		return Move{}, err
	}

	dst, err := ParsePos(line[1:])
	if err != nil {
		return Move{}, err
	}

	myPiece := p * Piece(c)
	var cand []Move
	for _, a := range AllMoves(b, c) {
		if b.At(a.Src) == myPiece && a.Dst == dst {
			cand = append(cand, a)
		}
	}
	switch len(cand) {
	case 0:
		return Move{}, illegalMove(line)
	case 1:
		return cand[0], nil
	default:
		return Move{}, ambiguousMove(line, cand)
	}
}

func ParsePiece(line string) (Piece, error) {
	p, ok := map[string]Piece{
		"P": WP, "p": WP,
		"R": WR, "r": WR,
		"N": WN, "n": WN,
		"B": WB, "b": WB,
		"Q": WQ, "q": WQ,
		"K": WK, "k": WK,
	}[line]
	if !ok {
		return 00, syntaxError(line)
	}
	return p, nil
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
		return Pos{}, syntaxError(txt)
	}
	c := int(txt[0]) - int('a')
	r := int(txt[1]) - int('1')

	if r < 0 || r > 7 || c < 0 || c > 7 {
		return Pos{}, syntaxError(txt)
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

func syntaxError(line string) error {
	return fmt.Errorf("i don't understand %q", line)
}

func illegalMove(m interface{}) error {
	return fmt.Errorf("%v is illegal", m)
}

func ambiguousMove(line string, cand []Move) error {
	return fmt.Errorf("%v is ambiguous, matches: %v", line, cand)
}
