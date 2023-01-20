package models

import (
	"encoding/json"
)

type Point struct {
	x int
	y int
}

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

func MisplacedTiles(size int, puzzle []int, goal []int) (uint16) {
	heuristic := 0
	for i, v := range puzzle {
		if (i % 3 != goal[v].x || i / 3 != goal[v].y) {
			heuristic += 1
		}
	}
	return heuristic
}

func ManhattanDistance(size int, puzzle []int, goal []int) (uint16) {
	heuristic := 0
	for i, v := range puzzle {
		if (v != 0 && i % 3 != goal[v].x || i / 3 != goal[v].y) {
			heuristic += math.Abs((i % 3) - goal[v].x) + math.Abs((i / 3) - goal[v].y)
		}
	}
	return heuristic
}

// https://medium.com/swlh/looking-into-k-puzzle-heuristics-6189318eaca2
func LinearConflict(size int, puzzle []int, goal []int) (uint16) {
	heuristic := 0
	for i, v := range puzzle {
		if (v != 0 && i % 3 != goal[v].x || i / 3 != goal[v].y) {
			heuristic += math.Abs((i % 3) - goal[v].x) + math.Abs((i / 3) - goal[v].y)
			if (i % 3 == goal[v].x) {
				for k := i + 3; k < len(puzzle); k += 3 {
					if (puzzle[k] != 0 && goal[puzzle[k]].y != k / 3) {
						heuristic += 2
					}
				}
			} else {
				for k := i + 1; k < size; k += 1 {
					if (puzzle[k] != 0 && goal[puzzle[k]].x != k % 3) {
						heuristic += 2
					}
				}
			}
		}
	}
	return heuristic
}

func MakeGoal(size int) (goal []Point) {
	goal := make([]Point, size * size)
	for i := range goal {
		goal[i] = -1
	}
	cur := 1
	x, y := 0
	ix, iy := 1, 0
	for {
		goal[cur] = Point{x, y}
		if (cur == 0) {
			break
		}
		cur += 1
		if (x + ix == size || x + ix < 0 || (ix != 0 && goal[x + ix + y*size] != -1)) {
			iy = ix
			ix = 0
		} else if  (y + iy == size || y + iy < 0 || (iy != 0 && goal[x + (y+iy)*size] != -1)) {
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
