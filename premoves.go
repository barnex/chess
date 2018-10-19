package chess

func AllPre(b *Board, c Color) []Move {
	var moves []Move
	pos := make([]Index, 0, 64)
	for i, p := range b {
		if p.Color() == c {
			pos = pos[:0]
			Pre(b, Index(i), &pos)
			for _, dst := range pos {
				moves = append(moves, Move{Index(i), dst})
			}

		}
	}
	return moves
}

func Pre(b *Board, src Index, dst *[]Index) {
	switch b.At(src) {
	case 00:
		return
	case WP:
		WPPre(b, src, dst)
	case BP:
		BPPre(b, src, dst)
	case WN, BN:
		NPre(b, src, dst)
	case WR, BR:
		RPre(b, src, dst)
	case WB, BB:
		BPre(b, src, dst)
	case WQ, BQ:
		QPre(b, src, dst)
	case WK, BK:
		KPre(b, src, dst)
	}
}

func QPre(b *Board, src Index, dst *[]Index) {
	RPre(b, src, dst)
	BPre(b, src, dst)
}

func RPre(b *Board, src Index, dst *[]Index) {
	pmarch(Front, b, src, dst)
	pmarch(Back, b, src, dst)
	pmarch(Left, b, src, dst)
	pmarch(Right, b, src, dst)
}

func BPre(b *Board, src Index, dst *[]Index) {
	pmarch(FrontRight, b, src, dst)
	pmarch(FrontLeft, b, src, dst)
	pmarch(BackRight, b, src, dst)
	pmarch(BackLeft, b, src, dst)
}

func pmarch(d Pos, b *Board, src Index, dst *[]Index) {
	for p := src.Pos().Add(d); p.Valid(); p = p.Add(d) {
		*dst = append(*dst, p.Index())
		if b.At(p.Index()).Color() != 00 {
			break
		}
	}
}

func NPre(b *Board, src Index, dst *[]Index) {
	d := []Pos{
		{-2, -1}, {-1, -2},
		{+2, -1}, {+1, -2},
		{-2, +1}, {-1, +2},
		{+2, +1}, {+1, +2}}
	pjump(d, b, src, dst)
}

func KPre(b *Board, src Index, dst *[]Index) {
	d := []Pos{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	pjump(d, b, src, dst)
}

func pjump(d []Pos, b *Board, src Index, dst *[]Index) {
	for _, d := range d {
		p := src.Pos().Add(d)
		if p.Valid() {
			*dst = append(*dst, p.Index())
		}
	}
}

func WPPre(b *Board, src Index, dst *[]Index) {
	// one row forward
	if p := src.Pos().Add(Front); p.Valid() && b.At(p.Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	// two rows forward
	if p := src.Pos().Add(Pos{2, 0}); p.Row() == 3 && b.At(p.Index()) == 00 && b.At(src.Pos().Add(Front).Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(FrontLeft); p.Valid() {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(FrontRight); p.Valid() {
		*dst = append(*dst, p.Index())
	}
}

func BPPre(b *Board, src Index, dst *[]Index) {
	// one row forward
	if p := src.Pos().Add(Back); p.Valid() && b.At(p.Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	// two rows forward
	if p := src.Pos().Add(Pos{-2, 0}); p.Row() == 4 && b.At(p.Index()) == 00 && b.At(src.Pos().Add(Back).Index()) == 00 {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(BackLeft); p.Valid() {
		*dst = append(*dst, p.Index())
	}

	if p := src.Pos().Add(BackRight); p.Valid() {
		*dst = append(*dst, p.Index())
	}
}
