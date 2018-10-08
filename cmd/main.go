package main

import (
	"fmt"
	"math"

	"github.com/barnex/chess"
)

func main() {

	engineA := chess.Minimax(0, chess.Heuristic1)
	engineB := chess.Minimax(1, chess.Heuristic1)

	var (
		totalMoves int
		wins       [3]int
	)

	numGames := 100

	for i := 0; i < numGames; i++ {
		w1, m1 := chess.Game(engineA, engineB)
		totalMoves += m1
		w2, m2 := chess.Game(engineB, engineA)
		totalMoves += m2

		wins[w1+1]++
		wins[-w2+1]++

		winA := float64(wins[2])
		winB := float64(wins[0])
		winX := float64(wins[1])
		totalGames := wins[0] + wins[1] + wins[2]

		score := winB / (winA + winB)
		err := 3 * math.Sqrt(winB+1) / (winA + winB)
		draw := winX / (winA + winB + winX)
		movesPerGame := float64(totalMoves) / (winA + winB + winX)

		fmt.Printf("%.1f%% +/- %.1f%% (%d games, %.1f%% draw, %.1f moves/game)\n",
			100*score, 100*err, totalGames, 100*draw, movesPerGame)

	}
}
