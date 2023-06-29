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
	Time             string    `json:"time"`
	Moves            int       `json:"moves"`
	Time_complexity  int       `json:"time_complexity"`
	Space_complexity int       `json:"space_complexity"`
	States           [][][]int `json:"states"`
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

func makeIntArr(size int, states []string) [][][]int {
	pathArr := make([][][]int, len(states))
	for i := range pathArr {
		pathArr[i] = make([][]int, size)
		strs := strings.Split(states[i], " ")
		for j := 0; j < size; j++ {
			pathArr[i][j] = make([]int, size)
			for k := 0; k < size; k++ {
				pathArr[i][j][k], _ = strconv.Atoi(strs[j*size+k])
			}
		}
	}
	return pathArr
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

func recursive_path(size int, mapmap map[string]string, currToString string) (path []string) {
	if v, ok := mapmap[currToString]; ok == true {
		path = append([]string{currToString}, path...)
		return append(recursive_path(size, mapmap, v), path...)
	}
	return
}

func printPuzzle(size int, puzzleStr string) {
	fmt.Printf("\n\n\n")
	strs := strings.Split(puzzleStr, " ")
	puzzle := make([]int, len(strs))
	for i := range puzzle {
		puzzle[i], _ = strconv.Atoi(strs[i])
	}
	for i := 0; i < len(puzzle); i++ {
		if i != 0 && i%size == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%2d ", puzzle[i])
	}
	fmt.Printf("\n")
	if size == 3 {
		time.Sleep(30 * time.Millisecond)
	} else if size == 4 {
		time.Sleep(20 * time.Millisecond)
	} else if size >= 5 {
		time.Sleep(10 * time.Millisecond)
	}
}

func Solve(size int, initPuzzle []int, heuristic string, search string, print bool) (SolveData, error) {
	start := time.Now()
	timeComplexity := 0
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

		openLength := openQueue.Len()
		closedLength := len(closedMap)
		timeComplexity += 1
		if openLength+closedLength > spaceComplexity {
			spaceComplexity = openLength + closedLength
		}
		for _, children := range getNeighborPuzzles(size, current.puzzle) {
			childToString := PuzzleToString(children)
			isGoal := AreArraysEqual(children, goal)
			if isGoal {
				time_ := time.Since(start)

				path := recursive_path(size, closedMap, currToString)
				path = append(path, childToString)
				if print {
					for _, v := range path {
						printPuzzle(size, v)
					}
					fmt.Println("\nTime complexity  : " + fmt.Sprintf("%d", timeComplexity))
					fmt.Println("Space complexity : " + fmt.Sprintf("%d", spaceComplexity))
					fmt.Println("Number of moves  : " + fmt.Sprintf("%d", current.move+1))
					fmt.Println("Time to solve    : " + fmt.Sprintf("%s", time_) + "")
				}
				test := fmt.Sprint("%s", time_)
				return SolveData{
					Time:             test,
					Time_complexity:  timeComplexity,
					Space_complexity: spaceComplexity,
					Moves:            current.move + 1,
					States:           makeIntArr(size, path),
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
