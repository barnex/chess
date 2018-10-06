package main

func Moves(b *Board, src Pos, dst *[]Pos) {
	switch b.At(src) {
	case 00:
		return
	case wP:
		WPMoves(b, src, dst)
	case bP:
		BPMoves(b, src, dst)
	case wN, bN:
		NMoves(b, src, dst)
	case wR, bR:
		RMoves(b, src, dst)
	case wB, bB:
		BMoves(b, src, dst)
	case wQ, bQ:
		QMoves(b, src, dst)
	case wK, bK:
		KMoves(b, src, dst)
	}
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
