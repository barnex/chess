package chess

import "fmt"

func Game(we, be Engine) (winner Color, moves int) {

	b := NewBoard()

	max := 100
	currPlayer := White
	players := map[Color]Engine{White: we, Black: be}
	for i := 0; i < max; i++ {
		if w := b.Winner(); w != 0 {
			return w, i
		}

		m := players[currPlayer].Move(b, currPlayer)
		b.Move(m.Src, m.Dst)
		fmt.Println(currPlayer, m, "\n", b)

		currPlayer = -currPlayer
	}
	return 0, max
}
