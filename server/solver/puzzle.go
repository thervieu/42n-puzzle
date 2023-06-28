package solver

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func AreArraysEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func MakeGoal(size int) []int {
	goal := make([]int, size*size, size*size)

	cur := 1
	x, y := 0, 0
	ix, iy := 1, 0
	for {
		goal[x+y*size] = cur

		cur += 1
		if cur == size*size {
			break
		}
		if x+ix == size || x+ix < 0 || (ix != 0 && goal[x+ix+y*size] != 0) {
			iy = ix
			ix = 0
		} else if y+iy == size || y+iy < 0 || (iy != 0 && goal[x+(y+iy)*size] != 0) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
	}
	return goal
}

// def make_puzzle(s, solvable, iterations):
// def swap_empty(p):
// 	idx = p.index(0)
// 	poss = []
// 	if idx % s > 0:
// 		poss.append(idx - 1)
// 	if idx % s < s - 1:
// 		poss.append(idx + 1)
// 	if idx / s > 0 and idx - s >= 0:
// 		poss.append(idx - s)
// 	if idx / s < s - 1:
// 		poss.append(idx + s)
// 	swi = random.choice(poss)
// 	p[idx] = p[swi]
// 	p[swi] = 0

// p = make_goal(s)
// for i in range(iterations):
// 	swap_empty(p)

func MakeShuffle(size int) []int {
	rand.Seed(time.Now().Unix())
	puzzle := MakeGoal(size)
	var iterations int
	if size == 3 {
		iterations = 40
	}
	if size == 4 {
		iterations = 120
	}
	if size > 4 {
		iterations = 1000
	}
	for i := 0; i < iterations; i++ {
		idx := findPos(0, puzzle)
		var slice []int
		if idx%size > 0 {
			slice = append(slice, idx-1)
		}
		if idx%size < size-1 {
			slice = append(slice, idx+1)
		}
		if idx/size > 0 && idx > size {
			slice = append(slice, idx-size)
		}
		if idx/size < size-1 {
			slice = append(slice, idx+size)
		}
		swapIndex := slice[rand.Intn(len(slice))]
		puzzle[idx] = puzzle[swapIndex]
		puzzle[swapIndex] = 0
	}
	return puzzle
}

func findPos(value int, puzzle []int) int {
	for i, v := range puzzle {
		if v == value {
			return i
		}
	}
	return -1
}

func numberOfInversions(puzzle []int) int {
	inv_count := 0
	for i := 0; i < len(puzzle)-1; i++ {
		for j := i + 1; j < len(puzzle); j++ {
			if puzzle[j] != 0 &&
				puzzle[i] != 0 && puzzle[i] > puzzle[j] {
				inv_count += 1
			}
		}
	}
	return inv_count
}

// https://www.cs.princeton.edu/courses/archive/spring21/cos226/assignments/8puzzle/specification.php#:~:text=In%20summary%2C%20when%20n%20is,number%20of%20inversions%20is%20even.
func IsSolvable(size int, p []int, goal []int) bool {
	inv_puzzle := numberOfInversions(p)
	inv_goal := numberOfInversions(goal)
	if size%2 == 0 {
		inv_puzzle += findPos(0, p) / size
		inv_goal += findPos(0, goal) / size
	}
	return inv_puzzle%2 == inv_goal%2
}

func PuzzleToString(p []int) string {
	if len(p) == 0 {
		return ""
	}

	str := make([]string, len(p))
	for i, v := range p {
		str[i] = strconv.Itoa(v)
	}
	return strings.Join(str, " ")
}
