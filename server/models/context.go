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

func MakeGoal(size int) (goal []Point) {
	goal = make([]Point, size * size)
	for i := range goal {
		goal[i] = Point{-1, -1}
	}
	cur := 1
	x, y := 0, 0
	ix, iy := 1, 0
	for {
		goal[cur] = Point{x, y}
		if (cur == 0) {
			break
		}
		cur += 1
		if (x + ix == size || x + ix < 0 || (ix != 0 && goal[x + ix + y*size] != Point{-1, -1})) {
			iy = ix
			ix = 0
		} else if  (y + iy == size || y + iy < 0 || (iy != 0 && goal[x + (y+iy)*size] != Point{-1, -1})) {
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
