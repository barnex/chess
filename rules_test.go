package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRules1(t *testing.T) {
	b := NewBoard()

	for _, c := range []struct {
		src  string
		want []string
	}{
		{"a2", []string{"a3", "a4"}},
	} {
		src := MustParse(c.src)
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
