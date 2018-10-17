package chess

import (
	"bytes"
	"fmt"
)

type Bits uint64

func (x Bits) At(pos uint8) uint64 {
	return uint64((x >> pos) & 0x1)
}

func (x Bits) Row(r uint8) uint8 {
	return uint8((x >> r) & 0xFF)
}

func (x Bits) String() string {
	var buf bytes.Buffer
	for r := 7; r >= 0; r-- {
		for c := 0; c < 8; c++ {
			if x.At(uint8(RC(r, c).index())) == 1 {
				fmt.Fprint(&buf, "x")
			} else {
				fmt.Fprint(&buf, ".")
			}
		}
		fmt.Fprintln(&buf)
	}
	return buf.String()
}
