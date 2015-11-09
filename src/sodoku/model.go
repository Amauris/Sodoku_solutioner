package sodoku;

import (
	//"fmt"
	"strings"
	"strconv"
	"math"
)

type Sodoku struct {
	Table [][]int
	preDefinedLayout string
	dimensions int
	rowFamilyCache map[string][]int
	columnFamilyCache map[string][]int
	quadrantFamilyCache map[string][]int
}

func GetCleanTable(dimensions int) *Sodoku {

	newSodoku := createTable("", dimensions)
	//instantiate table
	newSodoku.initTable()
	//cache that makes it easy for 
	//accessing rows, columns, quadrants
	//newSodoku.setFamilyCache()

	return newSodoku
}

func GetPreDefinedTable(layout string, dimensions int) *Sodoku {

	newPreDefinedSodoku := createTable(layout, dimensions)

	//instantiate table
	newPreDefinedSodoku.initTable()
	//fill table with the right values
	newPreDefinedSodoku.fillTable()
	//cache that makes it easy for 
	//accessing rows, columns, quadrants
	//newSodoku.setFamilyCache()

    return newPreDefinedSodoku
}

func createTable(layout string, dimensions int) *Sodoku {

	newTable := &Sodoku{[][]int{}, layout, dimensions, make(map[string][]int), make(map[string][]int), make(map[string][]int)}

	return newTable
}

func (inst *Sodoku) initTable() {

	table := make([][]int, inst.dimensions)

        for i, _ := range(table) {
                table[i] = make([]int, inst.dimensions)
        }
}

func (inst *Sodoku) fillTable() {

	rows := strings.Split(inst.preDefinedLayout, "\n")
	table := make([][]int, inst.dimensions)

        for i, _ := range(table) {

			columns := strings.Split(rows[i], " ")

			if i>=inst.dimensions {
				continue
			}

			table[i] = make([]int, inst.dimensions)

			for j, v := range(columns) {

				if j>=inst.dimensions {
	        	                continue
		                }

				if v=="_" {
					table[i][j] = 0
				} else {
					vInt, err := strconv.Atoi(v)

					if err!=nil {
						table[i][j] = 0
					} else {
						table[i][j] = vInt
					}
				}
			}
        }

	inst.Table = table
}

//traverse through all family(rows, columns, and quadrants)
//and insert to private family type cache
func (inst *Sodoku) setFamilyCache(i, j int) {

	//set columns cache
	//for every 
}

func (inst *Sodoku) GetRow(i int) []int {
	return inst.Table[i]
}

func (inst *Sodoku) GetColumn(j int) []int {

	row := make([]int, inst.dimensions)

	for i, r:= range(inst.Table) {
		row[i] = r[j]
	}

	return row

}

func (inst *Sodoku) GetQuadrant(i, j int) []int {

	quadrant := make([]int, inst.dimensions)
	quadrantX := int(math.Floor(float64(i/3)))
	quadrantY := int(math.Floor(float64(j/3)))

	quadrantTemp := inst.Table[quadrantX:quadrantX+3]
	for i, v := range(quadrantTemp) {
		tempSubRow := v[quadrantY:quadrantY+3]

		for j, v2 := range(tempSubRow) {
			quadrant[i+j] = v2
		}
	}

	return quadrant
}

//map i, j to hash so you can retrieve
//from appropriate rowFamilyCache index
func (inst *Sodoku) GetFamilies(i, j int) [][]int {

	rowFamily := inst.GetRow(i)
	columnFamily := inst.GetColumn(j)
	quadrantFamily := inst.GetQuadrant(i, j)

	return [][]int{rowFamily, columnFamily, quadrantFamily}
}

//map i, j to hash so you can retrieve
//from appropriate columnFamilyCache index
func (inst *Sodoku) getColumnFamily(i, j int) []int {

	return []int{}
}

//map i, j to hash so you can retrieve
//from appropriate quadrantFamilyCache index
func (inst *Sodoku) getQuadrantFamily(i, j int) []int {

	return []int{}
}

func (inst *Sodoku) ValidTable() {

}
