package chess

import (
	"math"
)

var NumEvals int64

func Inf(sign Color) float64 {
	if sign == 0 {
		return 0
	}
	return math.Inf(int(sign))
}
