package chess

import (
	"fmt"
	"strings"
)

type Pos [2]int8

var (
	Front      = Pos{1, 0}
	Back       = Pos{-1, 0}
	Right      = Pos{0, 1}
	Left       = Pos{0, -1}
	FrontRight = Front.Add(Right)
	FrontLeft  = Front.Add(Left)
	BackRight  = Back.Add(Right)
	BackLeft   = Back.Add(Left)
)

func (p Pos) Row() int {
	return int(p[0])
}

func (p Pos) Col() int {
	return int(p[1])
}

func (p Pos) Valid() bool {
	return (uint8(p[0])|uint8(p[1]))&0xF8 == 0
}

func (p Pos) Add(d Pos) Pos {
	return Pos{p[0] + d[0], p[1] + d[1]}
}

//func (p Pos) Next() Pos {
//	n := p.Add(Pos{0, 1})
//	if n.Col() == 8 {
//		n = RC(p.Row()+1, 0)
//	}
//	return n
//}

func (p Pos) String() string {
	if !p.Valid() {
		return fmt.Sprintf("(%v,%v)", p.Row(), p.Col())
	}
	return "abcdefgh"[p.Col():p.Col()+1] + "12345678"[p.Row():p.Row()+1]
}

func (p Pos) index() int {
	//if !p.Valid() {
	//	panic(fmt.Errorf("pos out of bounds: %v", p))
	//}
	return int(p[0]<<3 | p[1])
}

func (p Pos) Index() Index {
	return Index(p.index())
}

// TODO: rm
//func I2P(i int) Pos {
//	c := i & 0x7
//	r := i >> 3
//	return RC(r, c)
//}

func RC(r, c int) Pos {
	//if r < 0 || r > 7 || c < 0 || c > 7 {
	//	panic(fmt.Sprintf("pos out of bounds: (%v, %v)", r, c))
	//}
	return Pos{int8(r), int8(c)}
}

func P(p string) Pos {
	return Pos{
		int8(strings.Index("12345678", p[1:2])),
		int8(strings.Index("abcdefgh", p[0:1])),
	}
}
