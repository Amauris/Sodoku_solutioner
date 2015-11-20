package main;

import (
	"fmt"
	"os"
	"time"
	sodoku "sodoku"
	solutions "solutions"
)

func main() {
	
	start := time.Now()
	inputEntries := os.Args[1]

	//for i:=0; i<=10000; i+=1 {

		solutionizer := &solutions.Solutionizer{}
		board := sodoku.GetPreDefinedBoard(inputEntries, 9)

		solutionizer.GetSodokuAnswer(board)
		 fmt.Printf("%v difficulty\n", solutionizer.Difficulty())
	//}
	
    elapsed := time.Since(start)
    fmt.Println("Solutionizer took %s", elapsed)
}
