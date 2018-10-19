package chess

func IsCheck(b *Board) Color {
	for _, c := range []Color{White, Black} {
		for _, m := range AllMoves(b, c) {
			if b.At(m.DstI()) == WK*(-Piece(c)) {
				return -c
			}
		}
	}
	return 0
}

func AllMoves(b *Board, c Color) []Move {
	var moves []Move
	pos := make([]Index, 0, 64)
	for i_ := range b {
		i := Index(i_)
		if b.At(i).Color() == c {
			pos = pos[:0]
			Moves(b, i, &pos)
			for _, dst := range pos {
				moves = append(moves, Move{i, dst})
			}

		}
	}
	return moves
}

func Moves(b *Board, src Index, dst *[]Index) {
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
	if b.At(m.SrcI()).Color() != c {
		return false
	}
	var all []Index
	Moves(b, m.src, &all)
	for _, a := range all {
		if m.DstI() == a {
			return true
		}
	}
	return false
}

func QMoves(b *Board, src Index, dst *[]Index) {
	RMoves(b, src, dst)
	BMoves(b, src, dst)
}

func RMoves(b *Board, src Index, dst *[]Index) {
	march(Front, b, src, dst)
	march(Back, b, src, dst)
	march(Left, b, src, dst)
	march(Right, b, src, dst)
}

func BMoves(b *Board, src Index, dst *[]Index) {
	march(FrontRight, b, src, dst)
	march(FrontLeft, b, src, dst)
	march(BackRight, b, src, dst)
	march(BackLeft, b, src, dst)
}

func march(d Pos, b *Board, src Index, dst *[]Index) {
	me := b.At(src)
	for p := src.Pos().Add(d); p.Valid(); p = p.Add(d) {
		// empty
		if b.At(p.Index()).Color() == 00 {
			*dst = append(*dst, p.Index())
		}
		// capture
		if b.At(p.Index()).Color() == -me.Color() {
			*dst = append(*dst, p.Index())
			return
		}
		// blocked by own
		if b.At(p.Index()).Color() == me.Color() {
			return
		}
	}
}

func NMoves(b *Board, src Index, dst *[]Index) {
	d := []Pos{
		{-2, -1}, {-1, -2},
		{+2, -1}, {+1, -2},
		{-2, +1}, {-1, +2},
		{+2, +1}, {+1, +2}}
	jump(d, b, src, dst)
}

func KMoves(b *Board, src Index, dst *[]Index) {
	d := []Pos{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	jump(d, b, src, dst)
}

func jump(d []Pos, b *Board, src Index, dst *[]Index) {
	me := b.At(src)
	for _, d := range d {
		p := src.Pos().Add(d)
		if p.Valid() && b.At(p.Index()).Color() != me.Color() {
			*dst = append(*dst, p.Index())
		}
	}
}

func WPMoves(b *Board, src Index, dst *[]Index) {
	me := b.At(src)

	// one row forward
	if p := src.Pos().Add(Front); p.Valid() && b.At(p.Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	// two rows forward
	if p := src.Pos().Add(Pos{2, 0}); p.Row() == 3 && b.At(p.Index()) == 00 && b.At(src.Pos().Add(Front).Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(FrontLeft); p.Valid() && b.At(p.Index()).Color() == -me.Color() {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(FrontRight); p.Valid() && b.At(p.Index()).Color() == -me.Color() {
		*dst = append(*dst, p.Index())
	}
}

func BPMoves(b *Board, src Index, dst *[]Index) {
	me := b.At(src)

	// one row forward
	if p := src.Pos().Add(Back); p.Valid() && b.At(p.Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	// two rows forward
	if p := src.Pos().Add(Pos{-2, 0}); p.Row() == 4 && b.At(p.Index()) == 00 && b.At(src.Pos().Add(Back).Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(BackLeft); p.Valid() && b.At(p.Index()).Color() == -me.Color() {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(BackRight); p.Valid() && b.At(p.Index()).Color() == -me.Color() {
		*dst = append(*dst, p.Index())
	}
}
