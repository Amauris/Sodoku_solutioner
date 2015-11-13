package main;

import (
	"fmt"
	"os"
	sodoku "sodoku"
	solutions "solutions"
)

func main() {
	
	inputEntries := os.Args[1]

	solutionizer := &solutions.Solutionizer{}
	board := sodoku.GetPreDefinedBoard(inputEntries, 9)

	fmt.Println(solutionizer.GetSodokuAnswer(board))
}
