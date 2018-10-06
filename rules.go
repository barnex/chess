package main

func Moves(b *Board, src Pos, dst *[]Pos) {
	switch b.At(src) {
	case 00:
		return
	case wP:
		WPMoves(b, src, dst)
	case bP:
		BPMoves(b, src, dst)
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
