package main

import (
	"fmt"

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

	reset = "\033[39;49m"
)

func Render(b *Board) {
	for r := 7; r >= 0; r-- {

		fmt.Print(reset, fgDark, r+1, reset)

		for c := 0; c < 8; c++ {

			switch c {
			case 0:
				if (c+r)%2 == 0 {
					fmt.Print(reset, fgDark, halfR, bgDark, fgBlack)
				} else {
					fmt.Print(reset, fgLight, halfR, bgLight, fgBlack)
				}
			default:
				if (c+r)%2 == 0 {
					fmt.Print(reset, bgLight, fgDark, halfR, bgDark, fgBlack)
				} else {
					fmt.Print(reset, bgLight, fgDark, halfL, bgLight, fgBlack)
				}
			}
			piece := b.At(RC(r, c)).String()
			if b.At(RC(r, c)) == 0 {
				piece = " "
			}
			fmt.Print(piece)

		}
		if (r)%2 == 1 {
			fmt.Print(reset, fgDark, halfL, reset)
		} else {
			fmt.Print(reset, fgLight, halfL, reset)
		}
		fmt.Println()
	}
	fmt.Print(reset, fgDark, "  a b c d e f g h", reset, "\n")
}
