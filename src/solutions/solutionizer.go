package solutions;

import (
	"fmt"
	"math"
	sodoku "sodoku"
)

type Solutionizer struct {
	attempts int
}

type oppertunity struct {
	I int
	J int
	Entries []int
}

//main entry function that calls the right
//functions for gettnig sodoku answer
func (inst *Solutionizer) GetSodokuAnswer(board *sodoku.Board) string {

	inst.attempts = 0;

	inst.SetIndexWithLeastPossibleChoices(board)

	fmt.Println(board.GetStringFormat())

	return board.GetStringFormat()
}

//retrieves the family with the least
//free options to choose from. For example 
//if a particular row only has one blank
//index then we know we can fill that entry
//with 100% certainty(the unused number will go there)
func (inst *Solutionizer) SetIndexWithLeastPossibleChoices(board *sodoku.Board) bool {

	for !board.IsBoardComplete() {

		min := -1
		oppertunities := [9][]oppertunity{}

		//make sure we reset the current cursor
		board.ResetCursor()

		//fmt.Println(board.Entries)
		i, j, v := 0, 0, 0
		oppertunityFound := false
		for  i!=-1 {
			
			i, j, v = board.GetNextEntry()

			if(v!=0) {
				continue
			}

			families := board.GetFamilies(i, j)
			availableNumbers := 987654321

			for _, family := range(families) {
				availableNumbers = inst.availableNumbers(availableNumbers, family)
			}
			
			numAvailable := inst.getPossibilitiesFromAvailableNumbers(availableNumbers)
			length := len(numAvailable)
			//fmt.Println(numAvailable)
			if(length<=0) {
				return false
			//if we only have one choice to choose from then we know 100% we can set it
			} else if(length==1) {
				oppertunityFound = true
				board.SetEntry(i, j, numAvailable[0])
				///insert available number
			//
			} else if(min==-1 || len(numAvailable)<=min) {
				oppertunityFound = true
				min = len(numAvailable)
				oppertunities[min] = append(oppertunities[min], oppertunity{i, j, numAvailable})
				//minFamily = families
			}
		}

		if oppertunityFound==false {
			return false
		}

		if min>-1 {
			for _, ops := range(oppertunities) {
				for _, op := range(ops) {

					if len(op.Entries)<=0 {
						continue
					}

					for _, v := range(op.Entries) {
						board.SetEntry(op.I, op.J, v)
						
						if inst.SetIndexWithLeastPossibleChoices(board) {
							break
						} else {
							board.SetEntry(op.I, op.J, 0)
						}
					}
				}
				
			}
			/**/

			return false
		}
		
		inst.attempts += 1
	}

	return true
}

func (inst *Solutionizer) Difficulty() int {


	return inst.attempts
}

//retrieved numbders from availableNumbers and converts into array of numbers
func (inst *Solutionizer) getPossibilitiesFromAvailableNumbers(availableNumbers int) []int {

	possibilities := []int{}
	//first get the maxinum nth number available
	max := int(math.Log10(float64(availableNumbers))) + 1

	for max > 0 {

		if(!inst.isNumberTaken(availableNumbers, max)) {
			possibilities = append(possibilities, max)
		}

		max -= 1
	}

	return possibilities
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
	
	isTaken := (nth!=nthNumber)

	return isTaken
}