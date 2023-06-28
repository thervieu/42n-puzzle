package solver

import (
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SolveData struct {
	Time             int64    `json:"time"`
	Time_complexity  int      `json:"time_complexity"`
	Space_complexity int      `json:"space_complexity"`
	States           []string `json:"states"`
}

func (data SolveData) ToJSON() (string, error) {

	marshalled, err := json.Marshal(data)

	return string(marshalled), err
}

func getHeuristic(size int, heuristic string, search string, puzzle []int, goal []int) float32 {
	if heuristic == "hamming" {
		return HammingDistance(size, puzzle, goal)
	} else if heuristic == "manhattan" {
		return ManhattanDistance(size, puzzle, goal)
	} else {
		return EuclidianDistance(size, puzzle, goal)
	}
}

func initPriorityQueue(size int, heuristic string, search string, puzzle []int, goal []int) PriorityQueue {
	pq := make(PriorityQueue, 1)
	pq[0] = &State{
		priority: getHeuristic(size, heuristic, search, puzzle, goal),
		puzzle:   puzzle,
		move:     0,
		prev_pos: "00000",
	}
	heap.Init(&pq)
	return pq
}

func getNeighborPuzzles(size int, p []int) map[int][]int {
	neighbors := make(map[int][]int)
	indexZero := findPos(0, p)
	if indexZero/size > 0 && indexZero > size {
		neighbors[0] = make([]int, len(p))
		copy(neighbors[0], p)
		neighbors[0][indexZero] = neighbors[0][indexZero-size]
		neighbors[0][indexZero-size] = 0
	}
	if indexZero/size < size-1 {
		neighbors[1] = make([]int, len(p))
		copy(neighbors[1], p)
		neighbors[1][indexZero] = neighbors[1][indexZero+size]
		neighbors[1][indexZero+size] = 0
	}
	if indexZero%size > 0 {
		neighbors[2] = make([]int, len(p))
		copy(neighbors[2], p)
		neighbors[2][indexZero] = neighbors[2][indexZero-1]
		neighbors[2][indexZero-1] = 0
	}
	if indexZero%size < size-1 {
		neighbors[3] = make([]int, len(p))
		copy(neighbors[3], p)
		neighbors[3][indexZero] = neighbors[3][indexZero+1]
		neighbors[3][indexZero+1] = 0
	}
	return neighbors
}

func recursive_print(size int, mapmap map[string]string, currToString string) {
	if v, ok := mapmap[currToString]; ok == true {
		recursive_print(size, mapmap, v)
	}
	strs := strings.Split(currToString, ",")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	printPuzzle(size, ary)
	if size == 3 {
		time.Sleep(30 * time.Millisecond)
	} else if size == 4 {
		time.Sleep(20 * time.Millisecond)
	} else if size >= 5 {
		time.Sleep(10 * time.Millisecond)
	}
	return
}

func printPuzzle(size int, puzzle []int) {
	fmt.Printf("\n\n\n")
	for i := 0; i < len(puzzle); i++ {
		if i != 0 && i%size == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%2d ", puzzle[i])
	}
	fmt.Printf("\n")
}

func Solve(size int, initPuzzle []int, heuristic string, search string, print bool) (SolveData, error) {
	start := time.Now()
	// timeComplexity := 1
	spaceComplexity := 0

	goal := MakeGoal(size)
	if IsSolvable(size, initPuzzle, goal) == false {
		fmt.Println("This puzzle is unsolvable ! Please try another one")
		return SolveData{}, errors.New("Unsolvable")
	}

	openQueue := initPriorityQueue(size, heuristic, search, initPuzzle, goal)
	closedMap := make(map[string]string)

	for len(openQueue) != 0 {
		current := heap.Pop(&openQueue).(*State)
		currToString := PuzzleToString(current.puzzle)
		closedMap[currToString] = current.prev_pos

		queueLength := openQueue.Len()
		if queueLength > spaceComplexity {
			spaceComplexity = queueLength
		}
		for _, children := range getNeighborPuzzles(size, current.puzzle) {
			childToString := PuzzleToString(children)
			isGoal := AreArraysEqual(children, goal)

			if isGoal {
				time_ := time.Since(start)
				if print {
					recursive_print(size, closedMap, currToString)
					printPuzzle(size, children)
					// fmt.Println("Time complexity is" + fmt.Sprintf("%d", len()))
				}
				fmt.Println("Took " + fmt.Sprintf("%s", time_) + " seconds")
				return SolveData{
					Time: int64(time_),
					// Time_complexity: ,
					// Space_complexity: ,
					// States: ,
				}, nil
			} else if _, ok := closedMap[childToString]; ok == true {
				continue // if child in closeMap, do nothing
			} else {

				var priority float32
				if search == "greedy" {
					priority = getHeuristic(size, heuristic, search, children, goal)
				}
				if search == "uniform" {
					priority = float32(current.move)
				} else {
					priority = (float32(current.move) * 2) + getHeuristic(size, heuristic, search, children, goal)
				}
				newPuzzle := &State{
					priority: priority,
					puzzle:   children,
					prev_pos: currToString,
					move:     current.move + 1,
				}
				heap.Push(&openQueue, newPuzzle)
			}
		}
	}
	return SolveData{}, errors.New("Finished the open queue")
}
