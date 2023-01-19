package models

import (
	"math"
	"time"
	"strings"
	"strconv"
)

type OptionsRequestData struct {
	search		string	`json:"search"`
	heuristic	string	`json:"heuristic"`
}

type RequestSolveData struct {
	Options		OptionsRequestData	`json:"options"`
	State		string				`json:"state"`
	Analyzer	Analyzer			`json:"analyzer"`
}

func (r *RequestSolveData) extractOptions() (options Options, err error) {
	options.search = r.Options.search
	options.heuristic = r.Options.heuristic

	return
}

func generateNumbersSlice(strPuzzle string) (slice []int) {
	splittedStr := strings.Split(strPuzzle, ",")
	slice = make([]int, len(splittedStr))
	for i, str := range splittedStr {
        _int, err := strconv.Atoi(str)
		if (err != nil) {
			return
		}
		slice[i] = _int
    }
	return
}

func index(needle int, haystack []int) (int) {
	for i, v := range haystack {
		if (v == needle) {
			return i
		}
	}
	return -1
}

func (r *RequestSolveData) extractPuzzle() (puzzle Puzzle, err error) {

	puzzle.numbers = generateNumbersSlice(r.State)
	puzzle.size = int(math.Sqrt(float64((len(puzzle.numbers)))))
	puzzle.parent_move = ""
	puzzle.pos_zero = index(0, puzzle.numbers)

	return
}

func (r *RequestSolveData) ComputeContext() (context *Context, err error) {

	if options, err := r.extractOptions(); err != nil {
		return context, err
	} else {
		context.search = options.search == "greedy"
		context.DistanceHeuristic = // if else set function evaluating string options.heuristic
	}

	if puzzle, err := r.extractPuzzle(); err != nil {
		return context, err
	} else {
		context.Puzzle = puzzle
	}

	r.Analyzer = InitAnalyzer()

	return
}