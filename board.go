package main

type Board [64]Piece

func NewBoard() *Board {
	return &Board{
		WR, WP, 0, 0, 0, 0, BP, BR,
		WN, WP, 0, 0, 0, 0, BP, BN,
		WB, WP, 0, 0, 0, 0, BP, BB,
		WQ, WP, 0, 0, 0, 0, BP, BQ,
		WK, WP, 0, 0, 0, 0, BP, BK,
		WB, WP, 0, 0, 0, 0, BP, BB,
		WN, WP, 0, 0, 0, 0, BP, BN,
		WR, WP, 0, 0, 0, 0, BP, BR,
	}
}
