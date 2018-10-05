package main

import "fmt"

type Pos uint8

func (p Pos) Row() int {
	return int((p & 0xF0) >> 4)
}

func (p Pos) Col() int {
	return int(p & 0x0F)
}

func (p Pos) Valid() bool {
	return (p & 0x88) == 0
}

func (p Pos) Index() int {
	if !p.Valid() {
		panic(fmt.Errorf("pos out of bounds: %v", p))
	}
	return int(((p & 0xF0) >> 1) | (p & 0x0F))
}

func RC(r, c int) Pos {
	//if r < 0 || r > 7 || c < 0 || c > 7 {
	//	panic(fmt.Sprintf("pos out of bounds: (%v, %v)", r, c))
	//}
	return Pos(r<<4 | c)
}
