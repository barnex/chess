package chess

func AllPre(b *Board, c Color) []Move {
	var moves []Move
	pos := make([]Pos, 0, 64)
	for p := (Pos{}); p.Valid(); p = p.Next() {
		if b.At(p).Color() == c {
			pos = pos[:0]
			Pre(b, p, &pos)
			for _, dst := range pos {
				moves = append(moves, MoveP(p, dst))
			}

		}
	}
	return moves
}

func Pre(b *Board, src Pos, dst *[]Pos) {
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

func QPre(b *Board, src Pos, dst *[]Pos) {
	RPre(b, src, dst)
	BPre(b, src, dst)
}

func RPre(b *Board, src Pos, dst *[]Pos) {
	pmarch(Front, b, src, dst)
	pmarch(Back, b, src, dst)
	pmarch(Left, b, src, dst)
	pmarch(Right, b, src, dst)
}

func BPre(b *Board, src Pos, dst *[]Pos) {
	pmarch(FrontRight, b, src, dst)
	pmarch(FrontLeft, b, src, dst)
	pmarch(BackRight, b, src, dst)
	pmarch(BackLeft, b, src, dst)
}

func pmarch(d Pos, b *Board, src Pos, dst *[]Pos) {
	for p := src.Add(d); p.Valid(); p = p.Add(d) {
		*dst = append(*dst, p)
		if b.At(p).Color() != 00 {
			break
		}
	}
}

func NPre(b *Board, src Pos, dst *[]Pos) {
	d := []Pos{
		{-2, -1}, {-1, -2},
		{+2, -1}, {+1, -2},
		{-2, +1}, {-1, +2},
		{+2, +1}, {+1, +2}}
	pjump(d, b, src, dst)
}

func KPre(b *Board, src Pos, dst *[]Pos) {
	d := []Pos{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	pjump(d, b, src, dst)
}

func pjump(d []Pos, b *Board, src Pos, dst *[]Pos) {
	for _, d := range d {
		p := src.Add(d)
		if p.Valid() {
			*dst = append(*dst, p)
		}
	}
}

func WPPre(b *Board, src Pos, dst *[]Pos) {
	// one row forward
	if p := src.Add(Front); p.Valid() && b.At(p) == 00 {
		*dst = append(*dst, p)
	}

	// two rows forward
	if p := src.Add(Pos{2, 0}); p.Row() == 3 && b.At(p) == 00 && b.At(src.Add(Front)) == 00 {
		*dst = append(*dst, p)
	}

	if p := src.Add(FrontLeft); p.Valid() {
		*dst = append(*dst, p)
	}

	if p := src.Add(FrontRight); p.Valid() {
		*dst = append(*dst, p)
	}
}

func BPPre(b *Board, src Pos, dst *[]Pos) {
	// one row forward
	if p := src.Add(Back); p.Valid() && b.At(p) == 00 {
		*dst = append(*dst, p)
	}

	// two rows forward
	if p := src.Add(Pos{-2, 0}); p.Row() == 4 && b.At(p) == 00 && b.At(src.Add(Back)) == 00 {
		*dst = append(*dst, p)
	}

	if p := src.Add(BackLeft); p.Valid() {
		*dst = append(*dst, p)
	}

	if p := src.Add(BackRight); p.Valid() {
		*dst = append(*dst, p)
	}
}
