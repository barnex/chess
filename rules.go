package main

import "log"

func Moves(b *Board, src Pos, dst *[]Pos) {
	log.Println("moves", b.At(src))
	switch b.At(src) {
	case 00:
		return
	case WP:
		WPMoves(b, src, dst)
	}
}

func WPMoves(b *Board, src Pos, dst *[]Pos) {
	log.Println("wp moves", b.At(src))
	if p := src.Add(Front); p.Valid() && b.At(p) == 00 {
		*dst = append(*dst, p)
	}

	if p := src.Add(Pos{2, 0}); p.Row() == 3 && b.At(p) == 00 {
		*dst = append(*dst, p)
	}
}
