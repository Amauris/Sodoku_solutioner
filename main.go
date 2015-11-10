package main;

import (
	//"fmt"
	"os"
	sodoku "sodoku"
	solutions "solutions"
)

func main() {
	
	inputEntries := os.Args[1]

	board := sodoku.GetPreDefinedBoard(inputEntries, 9)
	solutions.GetSodokuAnswer(board)
}
