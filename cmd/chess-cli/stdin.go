package main

import (
	"bufio"
	"fmt"
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

	if len(line) != 4 {
		return Move{}, fmt.Errorf("syntax error: %q, need 4 characters", line)
	}

	src, err := ParsePos(line[:2])
	if err != nil {
		return Move{}, err
	}

	dst, err := ParsePos(line[2:4])
	if err != nil {
		return Move{}, err
	}

	m := Move{src, dst}
	if !Allowed(b, c, m) {
		return Move{}, fmt.Errorf("%v not allowed", m)
	}

	return m, nil
}

func (e *stdin) ReadLine() string {
	fmt.Print(e.prompt)
	e.scanner.Scan()
	if e.scanner.Err() != nil {
		os.Exit(0) // end of stream
	}
	return e.scanner.Text()
}

func Allowed(b *Board, c Color, m Move) bool {
	var all []Pos
	Moves(b, m.Src, &all)
	for _, a := range all {
		if m.Dst == a {
			return true
		}
	}
	return false
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
