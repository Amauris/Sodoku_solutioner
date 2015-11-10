package solutions;

import (
	"fmt"
	"math"
	sodoku "sodoku"
)

type Solutionizer struct {
	familyDensityOrdered []family
}

type family struct {
	members []int
	locX int
	locY int
}

//main entry function that calls the right
//functions for gettnig sodoku answer
func GetSodokuAnswer(board *sodoku.Board) {

	solutionizer := &Solutionizer{}
	
	try := 0
	//fmt.Println(board.Entries)
	for try<10 {
		solutionizer.GetIndexWithLeastPossibleChoices(board)
		try += 1
	}
	//fmt.Println(board.Entries)
	//InsertBestOption(mostDensedFamily)

	solutionizer.OutputAnswer(board)
	//fmt.Println(table.GetFamilies(1, 1))
}

func (inst *Solutionizer) Traverse() {
}

func (inst *Solutionizer) ParseInput() {

}

func (inst *Solutionizer) isNumberTaken(availableNumbers, nth int) bool {

	//assumes nth start from 0
	//check if nth position of 987654321 is zero
	//by dividing 987654321 by 10^nth and retrieving remainder
	//then with remainder you divide by 10^(n-1) in order to retrieve
	//nth digit
	nthFloat := float64(nth)
	remainder := availableNumbers%int(math.Pow(10, nthFloat))

	denominator := int(math.Pow(10, nthFloat-1))

	if(denominator==0) {
		return false
	}

	nthNumber := remainder/denominator
	//fmt.Printf("%v %v **\n", nthNumber, nth)
	
	isTaken := (nth!=nthNumber)

	return isTaken
}

func (inst *Solutionizer) availableNumbers(availableNumbers int, family []int) int {

	for _, v := range(family) {

		if(v==0 || inst.isNumberTaken(availableNumbers, v)) {
			continue
		}

		reducer := v*int(math.Pow(10, float64(v-1)))
		availableNumbers -= reducer
	}

	return availableNumbers
}

//retrieved numbders from availableNumbers and converts into array of numbers
func (inst *Solutionizer) getPossibilitiesFromAvailableNumbers(availableNumbers int) []int {

	possibilities := []int{}
	//first get the maxinum nth number available
	max := int(math.Log10(float64(availableNumbers))) + 1

	//fmt.Printf("%v %v\n", availableNumbers, max)

	for max > 0 {
		//fmt.Printf("values are %v %v %v \n", availableNumbers, max, inst.isNumberTaken(availableNumbers, max))
		if(!inst.isNumberTaken(availableNumbers, max)) {
			possibilities = append(possibilities, max)
		}

		max -= 1
	}

	return possibilities
}

//retrieves the family with the least
//free options to choose from. For example 
//if a particular row only has one blank
//index then we know the number that goes in that index
func (inst *Solutionizer) GetIndexWithLeastPossibleChoices(board *sodoku.Board) {

	min := -1
	//minFamily := [][]int{}

	for i, row := range(board.Entries) {

		for j, v := range(row) {

			//if we dont have an empty entry then we continue since this
			//entry is already filled
			if(v!=0) {
				continue
			}

			families := board.GetFamilies(i, j)
			availableNumbers := 987654321

			for _, family := range(families) {
				//fmt.Println(family)
				availableNumbers = inst.availableNumbers(availableNumbers, family)
			}

			numAvailable := inst.getPossibilitiesFromAvailableNumbers(availableNumbers)
			length := len(numAvailable)
			
			if(length<=0) {
				continue
			} else if(length==1) {
				fmt.Println(families)
				fmt.Println(numAvailable)
				board.SetEntry(i, j, numAvailable[0])
				///insert available number
			} else if(min==-1 || len(numAvailable)<=min) {
				min = len(numAvailable)
				//minFamily = families
			}
		}
	}

	//fmt.Printf("%v %v\n", min, minFamily)
	//fmt.Println(board.Entries)
}

func (inst *Solutionizer) OutputAnswer(board *sodoku.Board) {
	for _, row := range(board.Entries) {
		for _, entry := range(row) {
			if(entry==0) {
				fmt.Print("_")
			} else {
				fmt.Print(entry)
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}