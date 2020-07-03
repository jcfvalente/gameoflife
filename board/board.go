package board

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Grid represents all Cells
type Grid struct {
	Table [][]Cell
}

func (grid Grid) String() string {
	var builder strings.Builder
	for index := range grid.Table {
		for _, cell := range grid.Table[index] {
			builder.WriteString(fmt.Sprint("", cell))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

// Cell represents life: it has a position (x,y) and a status (alive or dead)
type Cell struct {
	XPosition int
	YPosition int
	Alive     bool
}

func (cell Cell) String() string {
	if cell.Alive {
		return fmt.Sprintf(" * ")
	}

	return fmt.Sprintf(" . ")
}

//CreateGrid initializes the board
func CreateGrid(size int) Grid {
	rand.Seed(time.Now().Unix()) /*Random seed each time  */
	table := make([][]Cell, size)
	for row := range table {
		table[row] = make([]Cell, size)
		for column := range table[row] {
			table[row][column].XPosition = row
			table[row][column].YPosition = column
			setRandomState(&table[row][column])
		}
	}
	return Grid{table}
}

func setRandomState(cell *Cell) {
	if random := rand.Float32(); random < 0.5 {
		cell.Alive = true
	} else {
		cell.Alive = false
	}
}
