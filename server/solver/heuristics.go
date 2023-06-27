package solver

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func HammingDistance(size int, puzzle []int, goal []int) float32 {
	var heuristic float32
	for i := 0; i < len(puzzle); i++ {
		if puzzle[i] != 0 && puzzle[i] != goal[i] {
			heuristic += 1
		}
	}
	return heuristic
}

func ManhattanDistance(size int, puzzle []int, goal []int) float32 {
	var heuristic float32
	for i := 0; i < len(puzzle); i++ {
		if puzzle[i] != 0 {
			for j := 0; j < len(goal); j++ {
				if puzzle[i] == goal[j] {
					heuristic += float32(abs(i%size-j%size) + abs(i/size-j/size))
				}
			}

		}
	}
	return heuristic
}

func EuclidianDistance(size int, puzzle []int, goal []int) float32 {
	var heuristic float32
	for i := 0; i < len(puzzle); i++ {
		if puzzle[i] != 0 {
			for j := 0; j < len(goal); j++ {
				if puzzle[i] == goal[j] {
					heuristic += float32(math.Sqrt(math.Pow(float64(i%size-j%size), 2) + math.Pow(float64(i/size-j/size), 2)))
				}
			}

		}
	}
	return heuristic
}
