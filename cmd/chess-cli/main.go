package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/barnex/chess"
)

func main() {
	b := NewBoard()

	engine := Minimax(4, Heuristic2)

	Render(b)
	for {

		m := ReadMove()
		for !Allowed(b, White, m) {
			fmt.Println(m, "not allowed")
			m = ReadMove()
		}
		src := b.At(m.Src)
		fmt.Println("player:", src, m)
		b = b.WithMove(m)
		Render(b)
		if b.Winner() != 0 {
			break
		}

		m2 := engine.Move(b, Black)
		src = b.At(m2.Src)
		b = b.WithMove(m2)
		fmt.Println("computer:", src, m2)
		Render(b)
		if b.Winner() != 0 {
			break
		}
	}

	fmt.Println(b.Winner(), "wins")
}

const (
	halfL   = "\u258C"
	halfR   = "\u2590"
	esc     = "\033["
	fgWhite = esc + "48;5;255;m"
	fgBlack = esc + "48;5;232;m"

	bgLight = esc + "48;5;250;m"
	bgDark  = esc + "48;5;245;m"

	fgLight = esc + "38;5;250;m"
	fgDark  = esc + "38;5;245;m"
)

func Render(b *Board) {
	fmt.Println(b.String(), "\n")
	//for r := 7; r >= 0; r-- {
	//	fmt.Println(r + 1)
	//	for c := 0; c < 8; c++ {

	//		if (r+c)%2 == 0 {
	//			fmt.Print(fgLight, bgDark, halfR, b.At(RC(r, c)))
	//		} else {
	//			fmt.Print(fgDark, bgLight, halfR, b.At(RC(r, c)))
	//		}
	//		//fmt.Println(" ", b.At(RC(r, c)))
	//		fmt.Print(halfL, halfR)

	//	}
	//	fmt.Println()
	//}
	//fmt.Println("  a b c d e f g h")
	//fmt.Println()
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

var scanner = bufio.NewScanner(os.Stdin)

func ReadMove() Move {
	var m Move
	var err error
	f := func() {
		fmt.Print("\n?>")
		scanner.Scan()
		if scanner.Err() != nil {
			os.Exit(0)
		}
		m, err = ParseMove(scanner.Text())
	}
	f()
	for err != nil {
		f()
	}
	return m
}

func ParseMove(txt string) (Move, error) {

	txt = strings.TrimSpace(txt)

	if len(txt) != 4 {
		return Move{}, fmt.Errorf("syntax error: %q, need 4 characters", txt)
	}

	src, err := ParsePos(txt[:2])
	if err != nil {
		return Move{}, err
	}

	dst, err := ParsePos(txt[2:4])
	if err != nil {
		return Move{}, err
	}

	return Move{src, dst}, nil
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
