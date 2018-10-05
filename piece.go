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
