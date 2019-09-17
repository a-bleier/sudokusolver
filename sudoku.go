package sudokusolver

import (
	"fmt"
	"github.com/a-bleier/util"
)

//Sodoku contains a 2-dimensional field of cells
type sudokuType struct {
	cellsField [][]cell
}


//TODO: Refactor spaghett code
func (sudoku *sudokuType) subsetElimination() {
	flag := true
	
	for flag{
		startAllOverFlag := false
		//Checks every row
		for n,row := range(sudoku.cellsField) {
			newRow := eliminateSubsets(row)			
			sudoku.cellsField[n] = newRow
			if compareCellSlices(newRow, row) == false { //if not equal
				startAllOverFlag = true
				break
			}
		}

		if startAllOverFlag {
			continue
		}
		
		//Checks every block		
		for colIndex := 0; colIndex < len(sudoku.cellsField); colIndex++ {
			var colList []cell
			for rowIndex := 0; rowIndex < len(sudoku.cellsField); rowIndex++ {
				cell := sudoku.cellsField[rowIndex][colIndex]
				colList = append(colList, cell)
			}

			newColList := eliminateSubsets(colList)
			for rowIndex := 0; rowIndex < len(sudoku.cellsField); rowIndex++ {
				cell := newColList[rowIndex]
				sudoku.cellsField[rowIndex][colIndex] = cell
			}


			if compareCellSlices(newColList, colList) == false { //if not equal
				startAllOverFlag = true
				break
			}
		}

		if startAllOverFlag {
			continue
		}
	
		
		//Checks every block
		blockIterator := newBlockSudokuIterator(*sudoku)
		blockIterator.first()	
		for(!blockIterator.isDone()){
			block := blockIterator.getCurrentItem()

			newBlock := eliminateSubsets(block)
			
			codata := blockIterator.getCoordinationData()

			for index,cell := range(newBlock) {
				sudoku.cellsField[codata.y + index / codata.dy][codata.x + index % codata.dx] = cell
			}
			blockIterator.next()
			if compareCellSlices(newBlock, block) == false { //if not equal
				startAllOverFlag = true
				break
			}
		}

		if startAllOverFlag {
			continue
		}
		
		flag = false
	}
}

func eliminateSubsets(cellSlice []cell) []cell {
	size := len(cellSlice)
	var newCellSlice []cell
	//copy cells from cellSLice in newCellSlice
	for _,cell := range(cellSlice){
		newCellSlice = append(newCellSlice, cell)
	}


	//searches in all subsets
	for i := 0; i< (1<<uint(size)); i++ {
		cellSlice = newCellSlice
		//map which holds a specific subset of the cellSlice; the position of a cell is stored as a key
		m := make(map[int]cell)
		for j := 0; j < size; j++ {

			if (i & (1 << uint(j))) > 0 {
				m[j] = cellSlice[j]
			}
		}
		
		
		lenMap := len(m)

		//empty subset
		if lenMap == 0 {
			continue
		}

		var numberList []int //numbers are stored here
		
		//lenMap numbers in lenMap cells
		for _, cell := range(m) {


				for x := 0; x < cell.getNumberOfPossibleValues(); x++ {
					value := cell.getValueAt(x)

					if !util.IntSliceHasValue(numberList, value) {
						numberList = append(numberList, value)
					}
									
				}
		}
	

		if len(numberList) != lenMap {
			continue
		} 

		//cells except the cells in the map will be cleaned from the numbers in the map
		for pos,cell := range(cellSlice) {
			continueFlag := false
			for key := range(m) {
				if(pos == key) {
					continueFlag = true
					break
				}
			}
			if continueFlag {
				continue
			}
			for _,value := range(numberList){
				cell.removeValue(value)
			}
			newCellSlice[pos] = cell

		}

		


	}

	return newCellSlice
	
}

func (sudoku *sudokuType) lineElimination() bool {
	var blockList [][]cell
	blockIterator := newBlockSudokuIterator(*sudoku)
	blockIterator.first()	
	//codata := blockIterator.getCoordinationData()
	for(!blockIterator.isDone()){
		block := blockIterator.getCurrentItem()

		blockList = append(blockList, block)

		blockIterator.next()
	}
	/*
	//Eliminate rows
	var rowBlockList [][]cell
	for index, block := range(blockList) {
		rowBlockList = append(rowBlockList, )
	}
	*/

	return false
}
/*
func (sudoku *sudokuType) blockElimination() bool {
	for _,row := range(sudoku.cellsField) {
		for _,cell := range(row) {
			
		}
	}

	return false
}
*/

//a sudoku is solved successfully when and only when eyery cell has only one value left and the winning conditions of a sudoku game are fullfilled
func (sudoku *sudokuType) isSolvedSuccessfully () bool {

	//Rows and cols check
	for colIndex := 0; colIndex < len(sudoku.cellsField); colIndex++ {
		checkColList := make([]int, len(sudoku.cellsField))
		checkRowList := make([]int, len(sudoku.cellsField))
		for rowIndex := 0; rowIndex < len(sudoku.cellsField); rowIndex++ {
			cellCol := sudoku.cellsField[rowIndex][colIndex]
			cellRow := sudoku.cellsField[colIndex][rowIndex]

			if(cellCol.getNumberOfPossibleValues() != 1 && cellRow.getNumberOfPossibleValues() != 1) {return false}

			value := cellCol.getPossibleValues()[0]
			if(checkColList[value-1] != 0) {return false}
			checkColList[value-1] = value

			value = cellRow.getPossibleValues()[0]
			if(checkRowList[value-1] != 0) {return false}
			checkRowList[value-1] = value

		}
	}

	//Block check
	blockIterator := newBlockSudokuIterator(*sudoku)
	blockIterator.first()	
	for(!blockIterator.isDone()){
		block := blockIterator.getCurrentItem()
		blockIterator.next()
		checkBlockList := make([]int, len(sudoku.cellsField))
		for _, cell := range(block) {
			value := cell.getPossibleValues()[0]
			if(checkBlockList[value-1] != 0) {return false}
			checkBlockList[value-1] = value
		}

	}
	return true
}

func compareCellSlices (cellSlice1 []cell, cellSlice2 []cell) bool {
	for index, cell := range(cellSlice1) {
		intSlice1 := cell.getPossibleValues()
		intSlice2 := cellSlice2[index].getPossibleValues()
		/*
		if(!util.IntSliceDeepCheck(intSlice1, intSlice2)){
			return false
		}
		*/
		if(len(intSlice1) != len(intSlice2)){
			return false
		}
	}

	return true
}

func (sudoku *sudokuType) printSudoku() {
	fmt.Println("------------------------------------------------------------------------")
	for _, row := range(sudoku.cellsField) {
		for _,cell := range(row){
			
			values := cell.getPossibleValues()

			for _,n := range(values) {
				fmt.Printf("%d,", n)
			}
			fmt.Printf(" | ")

		}
		fmt.Printf("\n")
	}
}

