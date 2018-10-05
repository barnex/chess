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
	//->1a->2b->3c->4d->5e->6f->7g->8h
	//->1h->1g->1f->1e->1d->1c->1b->1a
	//->8h->7g->6f->5e->4d->3c->2b->1a
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
