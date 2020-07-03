package main

import (
	"fmt"
	"time"

	"jcfv.me/go/life/board"
	"jcfv.me/go/life/generation"
)

func main() {
	fmt.Print("\033[H\033[2J") // Clear the screen
	gameGrid := board.CreateGrid(50)

	fmt.Print("\033[s") // Save position
	for {
		fmt.Print("\033[u\033[K") // Go back to saved position
		fmt.Print(gameGrid)
		generation.NextGeneration(&gameGrid)
		time.Sleep(time.Second / 4)
	}
}
