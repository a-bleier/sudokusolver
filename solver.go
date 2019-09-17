package sudokusolver
//The sudoku will be solved here
//Will only support sudokus with the size 4, 6 and 9

import (
	"fmt"
)


//RunSolver will be called from the frame work user
//TODO: Will return a [][]int later on
func RunSolver (loader Loader) int {
	//Loads a 2-dimensional array which resembles a game of sudoku
	sudokuField, sudokuSize := loader.LoadSudokuField()
	//checks whether the sudokuField is valid
	if(!sudokuFieldValidityCheck(sudokuField, sudokuSize)){
		fmt.Println("sudoku field is invalid")
		return 2

	}

	//a new sudoku object is initialized
	sudoku := newSudoku(sudokuField, sudokuSize)
	sudoku.printSudoku()
	solve(sudoku)

	return 0
}

func solve (sudoku sudokuType) sudokuType{

	for(!sudoku.isSolvedSuccessfully()){
		
		sudoku.subsetElimination()
		sudoku.printSudoku()

		

		//Guessing
		
	}

	
	sudoku.subsetElimination()
	sudoku.printSudoku()
	//fmt.Println("isSolved")

	
	return sudoku
}

func newSudoku (sudokuField [][]int, size int) sudokuType {
	var cellsField [][]cell
	for  row := 0; row < size; row++{		
		for col := 0; col < size; col++{

			//The array property of a cell is generated here

			if row == 0 {
				cellsField = append(cellsField, make([]cell, 0))
			}

			var values []int

			if sudokuField[row][col] == 0{
				for i:= 1; i < size+1; i++ {					
					values = append(values, i)
				}
			} else {
				values = append(values, sudokuField[row][col])
			}

			//cellArray added to cellsField
			cellsField[row] = append (cellsField[row], cell {possibleValues : values})

		}
	}

	

	return sudokuType{cellsField : cellsField}
}


func sudokuFieldValidityCheck (sudokuField [][]int, size int) bool {
	//Number of rows
	if(len(sudokuField) != size && size != 9){
		return false
	}
	for  row := 0; row < len(sudokuField); row++{
		
		//Number of col
		if(len(sudokuField[row]) != size){
			return false
		}
		for col := 0; col < len(sudokuField[row]); col++{
			
			
			//Cells
			if(sudokuField[row][col] < 0|| sudokuField[row][col] > size) {
				return false
			}
		}
	}
	return true
}





