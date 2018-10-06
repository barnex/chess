package main

type Piece int8

const (
	wP Piece = 1
	wR Piece = 2
	wN Piece = 3
	wB Piece = 4
	wQ Piece = 5
	wK Piece = 6
	bP Piece = -wP
	bR Piece = -wR
	bN Piece = -wN
	bB Piece = -wB
	bQ Piece = -wQ
	bK Piece = -wK
)

const (
	White = 1
	Black = -1
)

// https://en.wikipedia.org/wiki/Chess_symbols_in_Unicode
var pieceStr = map[Piece]rune{
	0:  '·',
	wP: '\u2659',
	wR: '\u2656',
	wN: '\u2658',
	wB: '\u2657',
	wQ: '\u2655',
	wK: '\u2654',
	bP: '\u265F',
	bR: '\u265C',
	bN: '\u265E',
	bB: '\u265D',
	bQ: '\u265B',
	bK: '\u265A',
}

func (p Piece) String() string {
	if r, ok := pieceStr[p]; ok {
		return string(r)
	}
	return "x" // for position marking in tests
}

func (p Piece) Color() int {
	switch {
	case p > 0:
		return White
	case p < 0:
		return Black
	default:
		return 0
	}
}
