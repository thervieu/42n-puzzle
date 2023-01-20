package models

import (
	"math"
	"strings"
	"strconv"
	"errors"
)

type RequestOptionsData struct {
	search		string	`json:"search"`
	heuristic	string	`json:"heuristic"`
}

type RequestSolveData struct {
	Options		RequestOptionsData	`json:"options"`
	State		[]int			`json:"state"`
}

func (r *RequestOptionsData) SanitizeOptions() (err error) {
	if (r.search != "uniform_cost" || r.search != "greedy") {
		return errors.New("Bad Request: Options.search has a wrong value")
	}
	if (r.heuristic != "linear_conflict" || r.heuristic != "manhattan" ||
		r.heuristic != "hamming") {
		return errors.New("Bad Request: Options.search has a wrong value")
	}

	return nil
}

func (r *RequestSolveData) SanitizeState() (error) {

	int_sqrt := int(math.Sqrt(float64(len(r.State))))
	if (int_sqrt * int_sqrt != len(r.State)) {
		return errors.New("Bad Request: State length is not a squared number")
	}
	if (int_sqrt * int_sqrt != len(r.State)) {
		return errors.New("Bad Request: State length is not a squared number")
	}

	slice := make([]int, len(r.State))
	for _,v := range r.State {
		if (v < 0) {
			return errors.New("Bad Request: State has a negative number")
		} 
		if (v >= len(r.State)) {
			return errors.New("Bad Request: State has a value bigger than its length")
		} else if (slice[v] == 1) {
			return errors.New("Bad Request: State has a value bigger than its length")
		} else { // slice[v] exists and is equal to 0
			slice[v] = 1
		}
	}

	return nil
}

func (r *RequestSolveData) Sanitize() (error) {
	if err := r.Options.SanitizeOptions(); err != nil {
		return err
	}
	if err := r.SanitizeState(); err != nil {
		return err
	}
	return nil
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

	puzzle.numbers = r.State
	puzzle.size = int(math.Sqrt(float64((len(puzzle.numbers)))))
	puzzle.parent_move = ""
	puzzle.pos_zero = index(0, puzzle.numbers)

	return
}

func (r *RequestSolveData) ComputeContext() (context *Context, err error) {

	if options, err := r.extractOptions(); err != nil {
		return context, err
	} else {
		context.SearchFormula = options.search == "greedy"
		if (options.heuristic == "hamming") {
			context.HeuristicFunc = HammingDistance // impl heuristic and give reference here
		} else if (options.heuristic == "manhattan") {
			context.HeuristicFunc = ManhattanDistance// same
		} else if (options.heuristic == "linear_conflict") {
			context.HeuristicFunc = LinearConflict// same
		}
	}

	if puzzle, err := r.extractPuzzle(); err != nil {
		return context, err
	} else {
		context.Puzzle = &puzzle
	}

	context.Analyzer = InitAnalyzer()

	return
}