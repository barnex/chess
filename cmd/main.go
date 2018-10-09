package main

import (
	"fmt"
	"math"

	"github.com/barnex/chess"
)

func main() {

	//engineA := chess.Random()
	//engineB := chess.Random()

	//engineA := chess.Random()
	//engineB := chess.Greedy(chess.Heuristic0)

	//engineA := chess.Greedy(chess.Heuristic0)
	//engineB := chess.Greedy(chess.Heuristic1)

	//engineA := chess.Greedy(chess.Heuristic1)
	//engineB := chess.Greedy(chess.Heuristic2)

	//engineA := chess.Greedy(chess.Heuristic2)
	//engineB := chess.Greedy(chess.Heuristic2)

	//engineA := chess.Greedy(chess.Heuristic2)
	//engineB := chess.Minimax(0, chess.Heuristic2)

	//engineA := chess.Minimax(0, chess.Heuristic2)
	//engineB := chess.Minimax(1, chess.Heuristic2)

	engineA := chess.Minimax(1, chess.Heuristic2)
	engineB := chess.Minimax(2, chess.Heuristic2)

	var (
		totalMoves int
		wins       [3]int
	)

	numRounds := 5000

	for i := 0; i < numRounds; i++ {
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
		err := 1 * math.Sqrt(winB+1) / (winA + winB) // TODO: use bernouilli stats
		draw := winX / (winA + winB + winX)
		movesPerGame := float64(totalMoves) / (winA + winB + winX)

		fmt.Printf("%.1f%% +/- %.1f%% (%d games, %.1f%% draw, %.1f moves/game)\n",
			100*score, 100*err, totalGames, 100*draw, movesPerGame)

	}
}
