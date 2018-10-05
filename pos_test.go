package main

import (
	"fmt"
	"testing"
)

func ExamplePos() {
	for p := RC(0, 0); p.Valid(); p = p.Add(FrontRight) {
		fmt.Print("->", p)
	}
	fmt.Println()
	for p := RC(0, 7); p.Valid(); p = p.Add(Left) {
		fmt.Print("->", p)
	}
	fmt.Println()
	for p := RC(7, 7); p.Valid(); p = p.Add(BackLeft) {
		fmt.Print("->", p)
	}

	//Output:
	//->a1->b2->c3->d4->e5->f6->g7->h8
	//->h1->g1->f1->e1->d1->c1->b1->a1
	//->h8->g7->f6->e5->d4->c3->b2->a1
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
