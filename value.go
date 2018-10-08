package chess

import (
	"fmt"
	"math"
)

type Value struct {
	Win       Color
	Heuristic float64
}

func (v Value) String() string {
	if v.Win != 0 {
		return v.Win.String()
	}
	return fmt.Sprint(v.Heuristic)
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

func Max(a, b Value) Value {
	if a.GT(b) {
		return a
	}
	return b
}

func Min(a, b Value) Value {
	if b.GT(a) {
		return a
	}
	return b
}

var MinusInf = Value{-1, math.Inf(-1)}
