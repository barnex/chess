package main

type Piece int8

const (
	WP Piece = 1
	WR Piece = 2
	WN Piece = 3
	WB Piece = 4
	WQ Piece = 5
	WK Piece = 6
	BP Piece = -WP
	BR Piece = -WR
	BN Piece = -WN
	BB Piece = -WB
	BQ Piece = -WQ
	BK Piece = -WK
)

// https://en.wikipedia.org/wiki/Chess_symbols_in_Unicode
var pieceStr = map[Piece]rune{
	0:  '·',
	WP: '\u2659',
	WR: '\u2656',
	WN: '\u2658',
	WB: '\u2657',
	WQ: '\u2655',
	WK: '\u2654',
	BP: '\u265F',
	BR: '\u265C',
	BN: '\u265E',
	BB: '\u265D',
	BQ: '\u265B',
	BK: '\u265A',
}

func (p Piece) String() string {
	return string(pieceStr[p])
}
