package main

import (
	"os"
	"fmt"
	"github.com/thervieu/42n-puzzle/server"
	"github.com/thervieu/42n-puzzle/parser"
	"github.com/thervieu/42n-puzzle/models"
)

func main() {
	args := os.Args[1:]

	if (len(args) == 0) {
		server.Start();
	} else {
		search, heuristic, fileName, err := parser.ParseArgs(args)

		if (err != nil) {
			fmt.Println(err)
			os.Exit(1)
		}
		
		size, startArray, err := parser.ParseFile(fileName)
		if (err != nil) {
			fmt.Println(err)
			os.Exit(1)
		}

		ctx := models.BuildContextCLI(size, search, heuristic, startArray)
		fmt.Println(ctx.Puzzle)
	}
}
