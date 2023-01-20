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

func MakeGoal(size uint8) (goal []int) {
	goal := make([]int, size * size)
	for i := range goal {
		goal[i] = -1
	}
	cur := 1
	x, y := 0
	ix, iy := 1, 0
	for {
		goal[x + y*s] = cur
		if (cur == 0) {
			break
		}
		cur += 1
		if (x + ix == s || x + ix < 0 || (ix != 0 && goal[x + ix + y*s] != -1)) {
			iy = ix
			ix = 0
		} else if  (y + iy == s || y + iy < 0 || (iy != 0 && goal[x + (y+iy)*s] != -1)) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if (cur == size * size) {
			cur = 0
		}
	}
	return 
}

func (c *Context) ComputeTimeline() (next *Context) {
	
	return
}

func (c *Context) ToJSON(time int64) (string, error) {
	response := ResponseSolveData{}
	response.computeResponseSolve(c, time, c.Analyzer.time_complexity, c.Analyzer.space_complexity)

	marshalled, err := json.Marshal(response)

	return string(marshalled), err
}
