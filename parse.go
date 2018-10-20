package chess

import (
	"fmt"
	"strings"
)

func MustParse(line string, b *Board, c Color) Move {
	m, err := Parse(line, b, c)
	if err != nil {
		panic(err)
	}
	return m
}

func Parse(line string, b *Board, c Color) (Move, error) {
	line = strings.TrimSpace(line)

	switch len(line) {
	default:
		return Move{}, syntaxError(line)
	case 2:
		return parse2(line, b, c)
	case 3:
		return parse3(line, b, c)
	case 4:
		return parse4(line)
	}
}

func ParseAndValidate(line string, b *Board, c Color) (Move, error) {
	move, err := Parse(line, b, c)
	if err != nil {
		return Move{}, err
	}
	if !Allowed(b, c, move) {
		return Move{}, illegalMove(move)
	}
	return move, nil
}

// parse2 parses a 2-character move, like
// 	a3
// which means to move a pawn to a3.
// Returns an error if the move is ambiguous or not allowed.
func parse2(line string, b *Board, c Color) (Move, error) {
	return parse3("P"+line, b, c)
}

// parse3 parses a 3-character move, like
// 	Nf3
// which means to move a knight to f3.
// Returns an error if the move is ambiguous or not allowed.
func parse3(line string, b *Board, c Color) (Move, error) {
	p, err := parsePiece(line[:1])
	if err != nil {
		return Move{}, err
	}

	dst, err := parsePos(line[1:])
	if err != nil {
		return Move{}, err
	}

	myPiece := p * Piece(c)
	var cand []Move
	for _, a := range AllMoves(b, c) {
		if b.At(a.SrcI()) == myPiece && a.Dst() == dst {
			cand = append(cand, a)
		}
	}
	switch len(cand) {
	case 0:
		return Move{}, illegalMove(line)
	case 1:
		return cand[0], nil
	default:
		return Move{}, ambiguousMove(line, cand)
	}
}

func parsePiece(line string) (Piece, error) {
	p, ok := map[string]Piece{
		"P": WP, "p": WP,
		"R": WR, "r": WR,
		"N": WN, "n": WN,
		"B": WB, "b": WB,
		"Q": WQ, "q": WQ,
		"K": WK, "k": WK,
	}[line]
	if !ok {
		return 00, syntaxError(line)
	}
	return p, nil
}

// parse4 parses a 4-character move, like
// 	a1b2
func parse4(line string) (Move, error) {
	src, err := parsePos(line[:2])
	if err != nil {
		return Move{}, err
	}

	dst, err := parsePos(line[2:4])
	if err != nil {
		return Move{}, err
	}

	m := MoveP(src, dst)
	return m, nil
}

func parsePos(txt string) (Pos, error) {
	if len(txt) != 2 {
		return Pos{}, syntaxError(txt)
	}
	c := int(txt[0]) - int('a')
	r := int(txt[1]) - int('1')

	if r < 0 || r > 7 || c < 0 || c > 7 {
		return Pos{}, syntaxError(txt)
	}
	return RC(r, c), nil
}

func syntaxError(line string) error {
	return fmt.Errorf("i don't understand %q", line)
}

func illegalMove(m interface{}) error {
	return fmt.Errorf("%v is illegal", m)
}

func ambiguousMove(line string, cand []Move) error {
	return fmt.Errorf("%v is ambiguous, matches: %v", line, cand)
}
