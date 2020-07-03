package generation

import (
	"jcfv.me/go/life/board"
)

/*
NextGeneration computes the next generation of the game. This determines
which cells will be alive or dead in the next iteration.
*/
func NextGeneration(gameGrid *board.Grid) {
	gameGridCopy := getCopyOfGrid(gameGrid)

	for row := range gameGrid.Table {
		for column := range gameGrid.Table[row] {
			currentCell := &gameGrid.Table[row][column]

			aliveNeigbours := getNumberOfAliveNeighbours(gameGridCopy.Table, currentCell)

			if currentCell.Alive {
				if aliveNeigbours < 2 {
					currentCell.Alive = false
				} else if aliveNeigbours > 3 {
					currentCell.Alive = false
				} else {
				}
			} else {
				if aliveNeigbours == 3 {
					currentCell.Alive = true
				}
			}
		}
	}
}

func getNumberOfAliveNeighbours(gridTable [][]board.Cell, cell *board.Cell) int {
	aliveNeighbours := 0
	sizeOfTable := len(gridTable)
	//N
	if cell.XPosition-1 >= 0 {
		//N
		if northNeighbour := gridTable[cell.XPosition-1][cell.YPosition]; northNeighbour.Alive {
			aliveNeighbours++
		}
		//NE
		if cell.YPosition+1 < sizeOfTable {
			if northEastNeighbour := gridTable[cell.XPosition-1][cell.YPosition+1]; northEastNeighbour.Alive {
				aliveNeighbours++
			}
		}
		//NW
		if cell.YPosition-1 >= 0 {
			if northWestNeighbour := gridTable[cell.XPosition-1][cell.YPosition-1]; northWestNeighbour.Alive {
				aliveNeighbours++
			}
		}
	}
	//E
	if cell.YPosition+1 < sizeOfTable {
		if eastNeighbour := gridTable[cell.XPosition][cell.YPosition+1]; eastNeighbour.Alive {
			aliveNeighbours++
		}
	}
	//W
	if cell.YPosition-1 >= 0 {
		if eastNeighbour := gridTable[cell.XPosition][cell.YPosition-1]; eastNeighbour.Alive {
			aliveNeighbours++
		}
	}
	//S
	if cell.XPosition+1 < sizeOfTable {
		//S
		if southNeighbour := gridTable[cell.XPosition+1][cell.YPosition]; southNeighbour.Alive {
			aliveNeighbours++
		}
		//SE
		if cell.YPosition+1 < sizeOfTable {
			if southEastNeighbour := gridTable[cell.XPosition+1][cell.YPosition+1]; southEastNeighbour.Alive {
				aliveNeighbours++
			}
		}
		//SW
		if cell.YPosition-1 >= 0 {
			if southWestNeighbour := gridTable[cell.XPosition+1][cell.YPosition-1]; southWestNeighbour.Alive {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

func getCopyOfGrid(gameGrid *board.Grid) board.Grid {
	size := len(gameGrid.Table)
	table := make([][]board.Cell, size)

	for row := range gameGrid.Table {
		table[row] = make([]board.Cell, size)
		for column := range gameGrid.Table[row] {
			table[row][column].XPosition = gameGrid.Table[row][column].XPosition
			table[row][column].YPosition = gameGrid.Table[row][column].YPosition
			table[row][column].Alive = gameGrid.Table[row][column].Alive
		}
	}
	return board.Grid{Table: table}
}
