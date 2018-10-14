package main

import (
	"fmt"
	"strings"

	. "github.com/barnex/chess"
)

const (
	halfL = "\u258C"
	halfR = "\u2590"

	bgLight = "\033[48;5;255m"
	bgDark  = "\033[48;5;250m"
	fgLight = "\033[38;5;255m"
	fgDark  = "\033[38;5;250m"
	fgBlack = "\033[38;5;232m"

	bg = "\033[48;5;"
	fg = "\033[38;5;"

	dark  = "252m"
	light = "231m"
	black = "232m"

	markDark  = "49m"
	markLight = "193m"

	reset = "\033[39;49m"
)

func Render(b *Board, mark map[Pos]bool, sideText string) {

	if mark == nil {
		mark = map[Pos]bool{}
	}

	colorOf := func(r, c int) string {
		if (r+c)%2 == 0 {
			if mark[RC(r, c)] {
				return markDark
			}
			return dark
		} else {
			if mark[RC(r, c)] {
				return markLight
			}
			return light
		}
	}

	side := strings.Split(sideText, "\n")
	for len(side) < 8 {
		side = append(side, "")
	}

	i := 0
	for r := 7; r >= 0; r-- {
		// print row number
		fmt.Print(reset, fgDark, r+1, reset)

		for c := 0; c < 8; c++ {

			piece := b.At(RC(r, c)).String()
			if b.At(RC(r, c)) == 0 {
				piece = " "
			}

			// print transition from previous to new tile+pice
			fmt.Print(reset)
			if c == 0 {
				fmt.Print(fg, colorOf(r, c), halfR, bg, colorOf(r, c), fg, black, piece)
			} else {
				fmt.Print(bg, colorOf(r, c-1), fg, colorOf(r, c), halfR, bg, colorOf(r, c), fg, black, piece)
			}
		}
		// print final transition, end of row
		fmt.Print(reset, fg, colorOf(r, 7), halfL)
		fmt.Println(reset, side[i])
		i++
	}
	// print column numbers
	fmt.Print(reset, fgDark, "  a b c d e f g h", reset, "\n")
}
