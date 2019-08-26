package sudokusolver

import (
	"fmt"
)

//Sodoku contains a 2-dimensional field of cells
type sudokuType struct {
	cellsField [][]cell
}

func (sudoku sudokuType) printSudoku() {
	for _, row := range(sudoku.cellsField) {
		for _,cell := range(row){
			fmt.Printf("%v\n", cell.getPossibleValues())
		}
	}
}

