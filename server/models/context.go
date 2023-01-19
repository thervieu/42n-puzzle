package models

import (
	"encoding/json"
)

type Context struct {
	Puzzle				*Puzzle
	Depth				int
	HeuristicFunc		func ([]int, []int) (int)
	SearchFormula		bool // false is uniform-cost, true is greedy
	HeuristicValue		int
	Analyzer 			Analyzer
	Parent				*Context
}

func (c *Context) Next(move string) (next *Context) {
	// move here
	next.puzzle = c.puzzle.CopyAndMove()
	next.Depth = c.Depth + 1
	next.HeuristicFunc = c.HeuristicFunc
	next.SearchFormula = c.SearchFormula

	heurValue := next.HeuristicFunc()
	if (next.SearchFormula) { // greedy
		next.HeuristicValue = heurValue
	} else {
		next.HeuristicValue = heurValue + c.HeuristicValue
	}

	next.Analyzer = c.Analyzer
	next.parent = c

	return
}

// implement heuristics
// set function pointer in extractOptions

func (c *Context) ComputeTimeline() (next *Context) {
	
	return
}

func (c *Context) ToJSON(time int64) (string, error) {
	response := ResponseSolveData{}
	response.computeResponseSolve(c, time, c.Analyzer.time_complexity, c.Analyzer.space_complexity)

	marshalled, err := json.Marshal(response)

	return string(marshalled), err
}
