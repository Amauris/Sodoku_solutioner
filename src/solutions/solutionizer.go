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
func GetSodokuAnswer(table *sodoku.Sodoku) {

	solutionizer := &Solutionizer{}
	solutionizer.GetIndexWithLeastPossibleChoices(table)
	
	//InsertBestOption(mostDensedFamily)

	//InsertFamilyToTable(table, mostDensedFamily)
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
func (inst *Solutionizer) GetIndexWithLeastPossibleChoices(table *sodoku.Sodoku) {

	min := -1
	//minFamily := [][]int{}
	fmt.Println(table.Table)
	for i, row := range(table.Table) {

		for j, _ := range(row) {

			families := table.GetFamilies(i, j)
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
				//fmt.Println(families)
				//fmt.Println(numAvailable)
				table.SetEntry(i, j, numAvailable[0])
				///insert available number
			} else if(min==-1 || len(numAvailable)<=min) {
				min = len(numAvailable)
				//minFamily = families
			}
		}
	}

	//fmt.Printf("%v %v\n", min, minFamily)
	fmt.Println(table.Table)
}

func (inst *Solutionizer) InsertFamilyToTable() {

}

func (inst *Solutionizer) outputAnswer() {

}

//check if sodoku table passes(confirms with the rules)
func (inst *Solutionizer) passesTable() {
}