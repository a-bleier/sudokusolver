package sudokusolver

import ("fmt")
//TODO: Can be made into a package with more generic two-dimensional structures
//TODO: Implement an index
//Maybe implement iterator for rows and cols ? 
type sudokuIterator interface {
	first()
	next()
	isDone() bool
	getCurrentItem() []cell
}



//Will implement iterator; iterates through the sudoku and returns a block as slice of cells
type blockSudokuIterator struct {
	currentBlock []cell
	sudoku sudokuType
	codata coordinationData
}

func newBlockSudokuIterator (sudoku sudokuType) blockSudokuIterator {
	return blockSudokuIterator{
		currentBlock : nil, 
		sudoku : sudoku, 
		codata : newCoordinationData(len(sudoku.cellsField))}
}

func (bi *blockSudokuIterator) first(){
	field := bi.sudoku.cellsField

	var cellblock []cell

	for y := 0; y < bi.codata.dy; y++ {
		for x := 0; x < bi.codata.dx; x++ {
			cellblock = append(cellblock, field[y][x])
		}
	}

	bi.currentBlock = cellblock
	
}

func (bi *blockSudokuIterator) next(){
	field := bi.sudoku.cellsField

	var cellblock []cell


	//Moves the block
	if(bi.codata.x == len(field)-bi.codata.dx){
		bi.codata.x = 0
		bi.codata.y += bi.codata.dy
	}else {
		bi.codata.x += bi.codata.dx
	}

	//Out of bounds
	if(bi.codata.y == len(field)){
		bi.currentBlock = nil
	}

	for y := bi.codata.y; y < bi.codata.y + bi.codata.dy; y++ {
		for x := bi.codata.x; x < bi.codata.x + bi.codata.dx; x++ {
			cellblock = append(cellblock, field[y][x])
		}
	}

	bi.currentBlock = cellblock

}

func (bi *blockSudokuIterator) isDone() bool {
	size := len(bi.sudoku.cellsField)
	data := bi.codata
	return data.x == size - data.dx && data.y == size - data.dy
}

func (bi *blockSudokuIterator) getCurrentItem() []cell {
	return bi.currentBlock
}

func (bi *blockSudokuIterator) getCoordinationData () coordinationData {
	return bi.codata
}

type coordinationData struct {
	x, y, dx, dy int
}

func newCoordinationData (size int) coordinationData{
	switch size {
	case 4:
		return coordinationData {0,0,2,2}
	case 6:
		return coordinationData {0,0,3,2}
	case 9:
		return coordinationData {0,0,3,3}
	default:
		return coordinationData {0,0,0,0}
	}

}

func printCoData(cd coordinationData ) {
	fmt.Printf(" %d %d\n", cd.x, cd.y)
}