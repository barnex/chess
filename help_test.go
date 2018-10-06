package main

import (
	"fmt"
	"sort"
)

func markAllMoves(b *Board, typ Piece) {
	for p := RC(0, 0); p.Valid(); p = p.Next() {
		if b.At(p) == typ {
			fmt.Printf("moves for %v%v\n", b.At(p), p)
			var moves []Pos
			Moves(b, p, &moves)
			mark(b, moves)
			fmt.Println()
		}
	}
}

func mark(b *Board, p []Pos) {
	c := b.Copy()
	for _, p := range p {
		c[p.Index()] = 100
	}
	fmt.Println(c)
}

func Sort(p []Pos) {
	sort.Slice(p, func(i, j int) bool { return p[i].String() < p[j].String() })
}