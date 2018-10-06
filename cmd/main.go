package main

import (
	"fmt"
	"time"

	"github.com/barnex/chess"
)

func main() {
	start := time.Now()
	w, m := chess.Game(chess.ERandom(), chess.ERandom())
	fmt.Println("the winner is", w, "in", m, "moves and", time.Since(start))
}
