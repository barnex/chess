package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
