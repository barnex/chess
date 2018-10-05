package main

import (
	"fmt"
	"testing"
)

func ExamplePos() {
	for p := RC(0, 0); p.Valid(); p = p.Add(Diag1) {
		fmt.Println(p)
	}

	//Output:
	//1a
	//2b
	//3c
	//4d
	//5e
	//6f
	//7g
	//8h
}

func TestRC(t *testing.T) {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			p := RC(r, c)
			if p.Row() != r || p.Col() != c {
				t.Errorf("RC: have: %v %v, want: %v %v", p.Row(), p.Col(), r, c)
			}
			if !p.Valid() {
				t.Errorf("valid: %v %v: have false, want true", r, c)
			}
		}
	}
}

func TestPos_Invalid(t *testing.T) {
	for _, p := range []Pos{
		RC(-1, 0),
		RC(0, -1),
		RC(-1, -1),
		RC(8, 0),
		RC(0, 8),
		RC(8, 8),
	} {
		if p.Valid() {
			t.Errorf("valid: %v: have true, want false", p)
		}
	}
}
