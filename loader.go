package sudokusolver

//Loader is used for dependency injection later on
type Loader interface  {
	//returns a 3-dimensional array
	LoadSudokuField() ([][]int, int)
}

//CsvLoader can load a sudoku in csv format. Visit the website https://qqwing.com/generate.html
type CsvLoader struct {
	csvText string
}


//LoadSudokuField is the implementation of interface Loader
//TODO: Hardcoded numbers here are really bad
func (cl CsvLoader) LoadSudokuField () ([][]int, int) {
	var sudokuField [][]int
	for i, symbol := range (cl.csvText) {

		if(i%9 == 0){
			sudokuField = append(sudokuField, make([]int, 0))
		}

		if(symbol == '.'){
			sudokuField[i/9] = append(sudokuField[i/9], 0)
		} else {
			sudokuField[i/9] = append(sudokuField[i/9], int(symbol) - 48)
		}
	}

	return sudokuField, 9
}

//NewCsvLoader creates a new CsvLoader 
func NewCsvLoader (text string) CsvLoader {
	return CsvLoader {csvText : text}
}

//SetCsvFormattedText sets new csvText
func (cl CsvLoader) SetCsvFormattedText (text string) {
	cl.csvText = text
}