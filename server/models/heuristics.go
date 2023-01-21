package models

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func HammingDistance(size int, puzzle []int, goal []Point) int {
	heuristic := 0
	for i, v := range puzzle {
		if i%size != goal[v].x || i/size != goal[v].y {
			heuristic += 1
		}
	}
	return heuristic
}

func ManhattanDistance(size int, puzzle []int, goal []Point) int {
	heuristic := 0
	for i, v := range puzzle {
		if v != 0 && i%size != goal[v].x || i/size != goal[v].y {
			heuristic += abs((i%size)-goal[v].x) + abs((i/size)-goal[v].y)
		}
	}
	return heuristic
}

// https://medium.com/swlh/looking-into-k-puzzle-heuristics-6189318eaca2
func LinearConflict(size int, puzzle []int, goal []Point) int {
	heuristic := 0
	for i, v := range puzzle {
		if v != 0 && i%size != goal[v].x || i/size != goal[v].y {
			heuristic += abs((i%size)-goal[v].x) + abs((i/size)-goal[v].y)
			if i%size == goal[v].x {
				for k := i + size; k < len(puzzle); k += size {
					if puzzle[k] != 0 && goal[puzzle[k]].y != k/size {
						heuristic += 2
					}
				}
			} else {
				for k := i + 1; k < size; k += 1 {
					if puzzle[k] != 0 && goal[puzzle[k]].x != k%size {
						heuristic += 2
					}
				}
			}
		}
	}
	return heuristic
}
