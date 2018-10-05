package main

import "fmt"

type Pos int

func (p Pos) Row() int {
	return int((p & 0x70) >> 3)
}

func (p Pos) Col() int {
	return int(p & 0x07)
}

func RC(r, c int) Pos {
	if r < 0 || r > 7 || c < 0 || c > 7 {
		panic(fmt.Sprintf("pos out of bounds: (%v, %v)", r, c))
	}
	return Pos(r<<3 | c)
}
