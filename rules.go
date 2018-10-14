package chess

func IsCheck(b *Board) Color {
	for _, c := range []Color{White, Black} {
		for _, m := range AllMoves(b, c) {
			if b.At(m.Dst) == WK*(-Piece(c)) {
				return -c
			}
		}
	}
	return 0
}

func AllMoves(b *Board, c Color) []Move {
	var moves []Move
	pos := make([]Pos, 0, 64)
	for p := (Pos{}); p.Valid(); p = p.Next() {
		if b.At(p).Color() == c {
			pos = pos[:0]
			Moves(b, p, &pos)
			for _, dst := range pos {
				moves = append(moves, Move{p, dst})
			}

		}
	}
	return moves
}

func Moves(b *Board, src Pos, dst *[]Pos) {
	switch b.At(src) {
	case 00:
		return
	case WP:
		WPMoves(b, src, dst)
	case BP:
		BPMoves(b, src, dst)
	case WN, BN:
		NMoves(b, src, dst)
	case WR, BR:
		RMoves(b, src, dst)
	case WB, BB:
		BMoves(b, src, dst)
	case WQ, BQ:
		QMoves(b, src, dst)
	case WK, BK:
		KMoves(b, src, dst)
	}
}

func Allowed(b *Board, c Color, m Move) bool {
	if b.At(m.Src).Color() != c {
		return false
	}
	var all []Pos
	Moves(b, m.Src, &all)
	for _, a := range all {
		if m.Dst == a {
			return true
		}
	}
	return false
}

func QMoves(b *Board, src Pos, dst *[]Pos) {
	RMoves(b, src, dst)
	BMoves(b, src, dst)
}

func RMoves(b *Board, src Pos, dst *[]Pos) {
	march(Front, b, src, dst)
	march(Back, b, src, dst)
	march(Left, b, src, dst)
	march(Right, b, src, dst)
}

func BMoves(b *Board, src Pos, dst *[]Pos) {
	march(FrontRight, b, src, dst)
	march(FrontLeft, b, src, dst)
	march(BackRight, b, src, dst)
	march(BackLeft, b, src, dst)
}

func march(d Pos, b *Board, src Pos, dst *[]Pos) {
	me := b.At(src)
	for p := src.Add(d); p.Valid(); p = p.Add(d) {
		// empty
		if b.At(p).Color() == 00 {
			*dst = append(*dst, p)
		}
		// capture
		if b.At(p).Color() == -me.Color() {
			*dst = append(*dst, p)
			return
		}
		// blocked by own
		if b.At(p).Color() == me.Color() {
			return
		}
	}
}

func NMoves(b *Board, src Pos, dst *[]Pos) {
	d := []Pos{
		{-2, -1}, {-1, -2},
		{+2, -1}, {+1, -2},
		{-2, +1}, {-1, +2},
		{+2, +1}, {+1, +2}}
	jump(d, b, src, dst)
}

func KMoves(b *Board, src Pos, dst *[]Pos) {
	d := []Pos{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	jump(d, b, src, dst)
}

func jump(d []Pos, b *Board, src Pos, dst *[]Pos) {
	me := b.At(src)
	for _, d := range d {
		p := src.Add(d)
		if p.Valid() && b.At(p).Color() != me.Color() {
			*dst = append(*dst, p)
		}
	}
}

func WPMoves(b *Board, src Pos, dst *[]Pos) {
	me := b.At(src)

	// one row forward
	if p := src.Add(Front); p.Valid() && b.At(p) == 00 {
		*dst = append(*dst, p)
	}

	// two rows forward
	if p := src.Add(Pos{2, 0}); p.Row() == 3 && b.At(p) == 00 && b.At(src.Add(Front)) == 00 {
		*dst = append(*dst, p)
	}

	if p := src.Add(FrontLeft); p.Valid() && b.At(p).Color() == -me.Color() {
		*dst = append(*dst, p)
	}

	if p := src.Add(FrontRight); p.Valid() && b.At(p).Color() == -me.Color() {
		*dst = append(*dst, p)
	}
}

func BPMoves(b *Board, src Pos, dst *[]Pos) {
	me := b.At(src)

	// one row forward
	if p := src.Add(Back); p.Valid() && b.At(p) == 00 {
		*dst = append(*dst, p)
	}

	// two rows forward
	if p := src.Add(Pos{-2, 0}); p.Row() == 4 && b.At(p) == 00 && b.At(src.Add(Back)) == 00 {
		*dst = append(*dst, p)
	}

	if p := src.Add(BackLeft); p.Valid() && b.At(p).Color() == -me.Color() {
		*dst = append(*dst, p)
	}

	if p := src.Add(BackRight); p.Valid() && b.At(p).Color() == -me.Color() {
		*dst = append(*dst, p)
	}
}
