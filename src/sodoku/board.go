package sodoku;

import (
	//"fmt"
	"strings"
	"strconv"
	"math"
)

type Board struct {
	cursor int
	Entries [][]int
	preDefinedLayout string
	dimensions int
	rowFamilyCache map[string][]int
	columnFamilyCache map[string][]int
	quadrantFamilyCache map[string][]int
}

func GetCleanBoard(dimensions int) *Board {

	newSodokuBoard := createBoard("", dimensions)
	//instantiate board
	newSodokuBoard.initBoard()
	//cache that makes it easy for 
	//accessing rows, columns, quadrants
	//newSodoku.setFamilyCache()

	return newSodokuBoard
}

func GetPreDefinedBoard(layout string, dimensions int) *Board {

	newPreDefinedSodokuBoard := createBoard(layout, dimensions)

	//instantiate board
	newPreDefinedSodokuBoard.initBoard()
	//fill board with the right values
	newPreDefinedSodokuBoard.fillBoard()
	//cache that makes it easy for 
	//accessing rows, columns, quadrants
	//newSodoku.setFamilyCache()

    return newPreDefinedSodokuBoard
}

func createBoard(layout string, dimensions int) *Board {

	newBoard := &Board{1, [][]int{}, layout, dimensions, make(map[string][]int), make(map[string][]int), make(map[string][]int)}

	return newBoard
}

func (inst *Board) initBoard() {

	entries := make([][]int, inst.dimensions)

    for i, _ := range(entries) {
            entries[i] = make([]int, inst.dimensions)
    }

    inst.Entries = entries
}

//returns the current cursor position
//as well as the entry value
func (inst *Board) GetNextEntry() (i, j, v int) {

	if inst.cursor>(inst.dimensions*inst.dimensions) {

		return -1, -1, -1
	}

	i = int((inst.cursor-1)/inst.dimensions)
	j = (inst.cursor-1)%inst.dimensions
	v = inst.Entries[i][j]

	inst.cursor += 1

	return i, j, v
}

func (inst *Board) ResetCursor() {

	inst.cursor = 1
}

func (inst *Board) SetCursor(i, j int) {

	j += 1

	inst.cursor = (i*inst.dimensions) + j
}

func (inst *Board) fillBoard() {

	rows := strings.Split(inst.preDefinedLayout, "\n")
	entries := make([][]int, inst.dimensions)

	i := 0
    for _, row := range(rows) {

    	row := strings.TrimSpace(row)

    	if(len(row)<=0) {
    		continue;
    	}

		columns := strings.Split(row, " ")

		if i>=inst.dimensions {
			break
		}

		entries[i] = make([]int, inst.dimensions)
		j := 0
		for _, entry := range(columns) {

			entry = strings.TrimSpace(entry)

			if(len(entry)<=0) {
    			continue;
    		}

			if j>=inst.dimensions {
				break
			}

			if entry=="_" {
				entries[i][j] = 0
			} else {
				entryInt, err := strconv.Atoi(entry)

				if err!=nil {
					entries[i][j] = 0
				} else {
					entries[i][j] = entryInt
				}
			}

			j += 1
		}

		i += 1
    }

	inst.Entries = entries
}

//traverse through all family(rows, columns, and quadrants)
//and insert to private family type cache
func (inst *Board) setFamilyCache(i, j int) {

	//set columns cache
	//for every 
}

func (inst *Board) GetRow(i int) []int {
	return inst.Entries[i]
}

func (inst *Board) SetEntry(i, j, value int) {
	inst.Entries[i][j] = value
}

func (inst *Board) GetColumn(j int) []int {

	row := make([]int, inst.dimensions)

	for i, r:= range(inst.Entries) {
		row[i] = r[j]
	}

	return row

}

func (inst *Board) GetQuadrant(i, j int) []int {

	quadrant := make([]int, inst.dimensions)
	quadrantX := int(math.Floor(float64(i/3)))*3
	quadrantY := int(math.Floor(float64(j/3)))*3

	quadrantTemp := inst.Entries[quadrantX:quadrantX+3]
	for i, v := range(quadrantTemp) {
		tempSubRow := v[quadrantY:quadrantY+3]
		for j, v2 := range(tempSubRow) {
			quadrant[(i*3)+j] = v2
		}
	}

	return quadrant
}

//map i, j to hash so you can retrieve
//from appropriate rowFamilyCache index
func (inst *Board) GetFamilies(i, j int) [][]int {

	rowFamily := inst.GetRow(i)
	columnFamily := inst.GetColumn(j)
	quadrantFamily := inst.GetQuadrant(i, j)

	return [][]int{rowFamily, columnFamily, quadrantFamily}
}

//map i, j to hash so you can retrieve
//from appropriate columnFamilyCache index
func (inst *Board) getColumnFamily(i, j int) []int {

	return []int{}
}

//map i, j to hash so you can retrieve
//from appropriate quadrantFamilyCache index
func (inst *Board) getQuadrantFamily(i, j int) []int {

	return []int{}
}


//make sure for every entry, its corresponding
//family is unique
func (inst *Board) IsBoardComplete() bool {

	for i, row := range(inst.Entries) {
		for j, _ := range(row) {
			families := inst.GetFamilies(i, j)
			for _, family := range(families) {
				numsToCount := make(map[int]int, 9)
				for _, value := range(family) {
					if numsToCount[value]==0 {
						numsToCount[value] = 1
					} else {
						return false
					}
				}
			}
		}
	}

	return true
}

func (inst *Board) GetStringFormat() string {

	boardString := ""

	for _, row := range(inst.Entries) {
		for _, entry := range(row) {
			if entry==0 {
				boardString += "_"
			} else {
				boardString += strconv.Itoa(entry)
			}
			boardString += " "
		}
		boardString += "\n" 
	}

	return boardString
}