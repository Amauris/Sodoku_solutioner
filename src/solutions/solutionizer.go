package solutions;

import (
	//"fmt"
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

	inst.SetIndicesWithLeastPossibleChoices(board)

	return board.GetStringFormat()
}

//retrieves the family with the least
//free options to choose from. For example 
//if a particular row only has one blank
//indices then we know we can fill that entry
//with 100% certainty(the unused number will go there)
func (inst *Solutionizer) SetIndicesWithLeastPossibleChoices(board *sodoku.Board) bool {

	toBeFilled := board.GetEmptyIndices()
	
	if len(toBeFilled)<=0 && board.IsBoardComplete() {
		return true
	}

	oppertunityFound := false
	oppertunities := [9][]oppertunity{}
	boardChange := false

	for  _, v := range(toBeFilled) {
		
		i, j := v[0], v[1]

		numAvailable := inst.getPossibilities(i, j, board)
		length := len(numAvailable)
		//fmt.Println(i, j, numAvailable)
		//if no oppertunities or certainties were found then
		//we are dealing with a faulty/broken board
		if(length<=0) {
			return false
		//if we only have one choice to choose from then we know 100% we can set it
		} else if(length==1) {
			//set indices for relatives
			board.SetEntry(i, j, numAvailable[0])
			boardChange = true
			//now that we have filled index at i,j
			//that means relatives were affected
			//and their might be new relatives we can fill with certainty
			//toBeFilledRelatives := board.GetFamilyEmptyIndices(i, j)
			//inst.SetIndicesWithCertainty(toBeFilledRelatives, board)

			//recall board recusrion
			//return inst.SetIndicesWithLeastPossibleChoices(board)
		//other wise we track available oppurtunities thats ordered
		//based on amount of numbers available
		} else {
			oppertunityFound = true
			oppertunities[length] = append(oppertunities[length], oppertunity{i, j, numAvailable})
		}
	}

	//if board was changed we call recursion on updated board
	if boardChange {
		return inst.SetIndicesWithLeastPossibleChoices(board)
	//else if at least one oppurtunity was found
	//we insert oppurtunity and recompute recursion
	} else if oppertunityFound {

		//make copy of current entries before any alterations
		origEntry := inst.copy(board.Entries)
		for _, ops := range(oppertunities) {

			for _, op := range(ops) {

				if len(op.Entries)<=0 {
					continue
				}

				for _, v := range(op.Entries) {
					
					board.SetEntry(op.I, op.J, v)
					//fmt.Printf("Trying %v %v %v\n", op.I, op.J, v)
					pass := inst.SetIndicesWithLeastPossibleChoices(board)
					//if recursion returns true
					if pass {
						//fmt.Println(board.Entries)
						return true
					//otherwise if this oppertunity wasnt the best choice
					//we set it back to 0 and try next oppertunity
					} else {
						board.SetEntries(origEntry)
					}
				}
				
			}
			
		}
	}

	inst.attempts += 1

	return false
}

func (inst *Solutionizer) copy(values [][]int) [][]int {

	a := make([][]int, 9)

	// manual deep copy
	for i := range(values) {
	    a[i] = make([]int, len(values[i]))
	    copy(a[i], values[i])
	}
	return a
}

//fills the board with answers it can get 100% right
//in other words where possible entries for emptyt(0) indices is equal 1
func (inst *Solutionizer) SetIndicesWithCertainty(indices [][]int, board *sodoku.Board) {

	for _, index := range(indices) {

		i, j := index[0], index[1]

		numAvailable := inst.getPossibilities(i, j, board)
		
		length := len(numAvailable)

		if(length==1) {
				
			board.SetEntry(i, j, numAvailable[0])

			//set indices for relatives
			toBeFilledRelatives := board.GetFamilyEmptyIndices(i, j)
			inst.SetIndicesWithCertainty(toBeFilledRelatives, board)	
		}
	}

	return
}

func (inst *Solutionizer) getPossibilities(i, j int, board *sodoku.Board) []int {

	families := board.GetFamilies(i, j)
	availableNumbers := 987654321

	for _, family := range(families) {
		availableNumbers = inst.availableNumbers(availableNumbers, family)
	}
		
	//fmt.Println(board.GetStringFormat())
	//fmt.Println(families)
	numsAvailable := inst.getPossibilitiesFromAvailableNumbers(availableNumbers)
	
	//fmt.Printf("%v %v %v \n", i, j, numsAvailable)

	return numsAvailable
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