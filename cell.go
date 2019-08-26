package sudokusolver

type cell struct {
	possibleValues []int
}

func (c cell) hasValue(value int) bool {
	for _, cellVal := range c.possibleValues {
		if(cellVal == value){
			return true
		}
	}
	return false
}

func (c cell) removeValue (value int){
	var newpossibleValues []int
	for _, cellVal := range c.possibleValues {
		if(cellVal != value){
			newpossibleValues = append(newpossibleValues, cellVal)
		}
	}
	c.possibleValues = newpossibleValues
}

func (c cell) getPossibleValues () []int {
	return c.possibleValues
}