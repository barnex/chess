package chess

func Game(we, be Engine) (winner Color, moves int) {

	b := NewBoard()

	max := 200
	currPlayer := White
	players := map[Color]Engine{White: we, Black: be}
	for i := 0; i < max; i++ {
		if w := b.Winner(); w != 0 {
			return w, i
		}

		m, _ := players[currPlayer].Move(b, currPlayer)
		b = b.WithMove(m)
		//	fmt.Println(currPlayer, m, "\n", b)

		currPlayer = -currPlayer
	}
	return 0, max
}
