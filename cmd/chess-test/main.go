//Command chess-test plays two engines against each other.
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/barnex/chess"
	"github.com/barnex/chess/engines/picard"
	"github.com/barnex/chess/engines/worf"
)

func main() {

	engineA := worf.New(1)
	engineB := picard.New(1)

	var (
		totalMoves int
		wins       [3]int
	)

	numRounds := 5000

	for i := 0; i < numRounds; i++ {
		w1, m1 := Game(engineA, engineB)
		totalMoves += m1
		w2, m2 := Game(engineB, engineA)
		totalMoves += m2

		wins[w1+1]++
		wins[-w2+1]++

		winA := float64(wins[2])
		winB := float64(wins[0])
		winX := float64(wins[1])
		totalGames := wins[0] + wins[1] + wins[2]

		score := winB / (winA + winB)
		err := 1 * math.Sqrt(winB+1) / (winA + winB) // TODO: use bernouilli stats
		draw := winX / (winA + winB + winX)
		movesPerGame := float64(totalMoves) / (winA + winB + winX)

		fmt.Printf("%.1f%% +/- %.1f%% (%d games, %.1f%% draw, %.1f moves/game)\n",
			100*score, 100*err, totalGames, 100*draw, movesPerGame)

	}
}

func Game(we, be chess.Engine) (winner chess.Color, moves int) {

	b := chess.NewBoard()

	max := 200
	currPlayer := chess.White
	players := map[chess.Color]chess.Engine{chess.White: we, chess.Black: be}
	for i := 0; i < max; i++ {
		if w := b.Winner(); w != 0 {
			return w, i
		}

		m, _ := players[currPlayer].Move(b, currPlayer)
		if m == (chess.Move{}) { // no possible moves, game over
			return -currPlayer, i
		}

		if !chess.Allowed(b, currPlayer, m) {
			log.Fatalf("BUG: illegal move by %T %v %v\n%v", players[currPlayer], currPlayer, m, b)
		}

		b = b.WithMove(m)
		//	fmt.Println(currPlayer, m, "\n", b)

		currPlayer = -currPlayer
	}
	return 0, max
}
