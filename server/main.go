package main

import (
	"fmt"
	"os"

	"github.com/thervieu/42n-puzzle/parser"
	"github.com/thervieu/42n-puzzle/server"
	"github.com/thervieu/42n-puzzle/solver"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		server.Start()
	} else {
		size, search, heuristic, fileName, err := parser.ParseArgs(args)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var initPuzzle []int
		if size != 0 {
			initPuzzle = solver.MakeShuffle(size)
		} else {
			size, initPuzzle, err = parser.ParseFile(fileName)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = solver.Solve(size, initPuzzle, heuristic, search, true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
