package models

import (
)

func abs(x int) (int) {
	if (x < 0) {
		return -x
	}
	return x
}

func HammingDistance(size int, puzzle []int, goal []Point) (int) {
	heuristic := 0
	for i, v := range puzzle {
		if (i % 3 != goal[v].x || i / 3 != goal[v].y) {
			heuristic += 1
		}
	}
	return heuristic
}

func ManhattanDistance(size int, puzzle []int, goal []Point) (int) {
	heuristic := 0
	for i, v := range puzzle {
		if (v != 0 && i % 3 != goal[v].x || i / 3 != goal[v].y) {
			heuristic += abs((i % 3) - goal[v].x) + abs((i / 3) - goal[v].y)
		}
	}
	return heuristic
}

// https://medium.com/swlh/looking-into-k-puzzle-heuristics-6189318eaca2
func LinearConflict(size int, puzzle []int, goal []Point) (int) {
	heuristic := 0
	for i, v := range puzzle {
		if (v != 0 && i % 3 != goal[v].x || i / 3 != goal[v].y) {
			heuristic += abs((i % 3) - goal[v].x) + abs((i / 3) - goal[v].y)
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