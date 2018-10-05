package main

import "fmt"

type Pos uint8

const (
	Front Pos = 0x01
	Back  Pos = 0x0F
	Right Pos = 0x10
	Left  Pos = 0xF0
	Diag1 Pos = Front + Right
	Diag2 Pos = Front + Left
	Diag3 Pos = Back + Right
	Diag4 Pos = Back + Left
)

func (p Pos) Row() int {
	return int((p & 0xF0) >> 4)
}

func (p Pos) Col() int {
	return int(p & 0x0F)
}

func (p Pos) Valid() bool {
	return (p & 0x88) == 0
}

func (p Pos) Add(d Pos) Pos {
	return p + d
}

func (p Pos) String() string {
	if !p.Valid() {
		return fmt.Sprintf("(%v,%v)", p.Row(), p.Col())
	}
	return "12345678"[p.Row():p.Row()+1] + "abcdefgh"[p.Col():p.Col()+1]
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
