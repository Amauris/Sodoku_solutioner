package main;

import (
	//"fmt"
	"os"
	sodoku "sodoku"
	solutions "solutions"
)

func main() {
	
	inputTable := os.Args[1]

	table := sodoku.GetPreDefinedTable(inputTable, 9)
	solutions.GetSodokuAnswer(table)
}
