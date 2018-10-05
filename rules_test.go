package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestRules1(t *testing.T) {
	b := NewBoard()
	b.Move(P("a7"), P("a3"))

	fmt.Println(b)

	for _, c := range []struct {
		src  string
		want []string
	}{
		{"a2", nil},
		{"b2", []string{"a3", "b3", "b4"}},
		{"c2", []string{"c3", "c4"}},
	} {
		src := P(c.src)
		var h []Pos
		Moves(b, src, &h)
		have := make([]string, len(h))
		for i, h := range h {
			have[i] = h.String()
		}
		sort.Strings(have)
		sort.Strings(c.want)
		if !reflect.DeepEqual(have, c.want) {
			t.Errorf("moves %v: have: %v, want: %v", src, have, c.want)
		}
	}
}

func Sort(p []Pos) {
	sort.Slice(p, func(i, j int) bool { return p[i].String() < p[j].String() })
}
