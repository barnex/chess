package chess

import "math"

type Value struct {
	Win       Color
	Heuristic float64
}

func (a Value) GT(b Value) bool {
	if a.Win > b.Win {
		return true
	}
	return a.Heuristic > b.Heuristic
}

func (a Value) Neg() Value {
	return a.Mul(-1)
}

func (a Value) Mul(c Color) Value {
	return Value{c * a.Win, float64(c) * a.Heuristic}
}

var MinusInf = Value{-1, math.Inf(-1)}
