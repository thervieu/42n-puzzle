package models

import (
	"encoding/json"
)

type Point struct {
	x int
	y int
}

type Context struct {
	Goal				[]Point
	Puzzle				Puzzle
	Depth				int
	HeuristicFunc		func (int, []int, []Point) (int)
	SearchFormula		bool // false is uniform-cost, true is greedy
	HeuristicValue		int
	Analyzer 			Analyzer
	Parent				*Context
}

func (c *Context) Next(p Puzzle) (next *Context) {
	// move here
	next.Puzzle = p
	next.Goal = c.Goal
	next.Depth = c.Depth + 1
	next.Parent = c

	return
}

func BuildContextCLI(size int, search string, heuristic string, arr []int) (c Context) {

	c.Puzzle = InitializePuzzle(size, arr, index(0, arr))
	c.Goal = MakeGoal(c.Puzzle.size)
	c.Depth = 0

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
