package models

import (
	"math"
	"strings"
	"strconv"
	"errors"
)

type RequestOptionsData struct {
	Search		string	`json:"search"`
	Heuristic	string	`json:"heuristic"`
}

type RequestSolveData struct {
	Options		RequestOptionsData	`json:"options"`
	State		[]int			`json:"state"`
}

func (r *RequestOptionsData) SanitizeOptions() (err error) {
	if (r.Search != "uniform_cost" && r.Search != "greedy") {
		return errors.New("Bad Request: Options.Search has a wrong value")
	}
	if (r.Heuristic != "linear_conflict" && r.Heuristic != "manhattan" ||
		r.Heuristic != "hamming") {
		return errors.New("Bad Request: Options.Search has a wrong value")
	}

	return nil
}

func VerifValuesSlice(arr []int, typeError bool) (error) {

	begError := ""
	if (typeError == true)  {
		begError = "Bad Request: State: " 
	} else {
		begError = "error: file: "
	}
	slice := make([]int, len(arr))
	for _,v := range arr {
		if (v < 0) {
			return errors.New(begError + "negative number")
		}
		if (v >= len(arr)) {
			return errors.New(begError + "value bigger than its length")
		}
		if (slice[v] == 1) {
			return errors.New(begError + "value " + strconv.Itoa(v) + " duplicate ")
		}
		slice[v] = 1
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

	return VerifValuesSlice(r.State, true)
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
	options.search = r.Options.Search
	options.heuristic = r.Options.Heuristic

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
		context.Puzzle = puzzle
	}

	context.Analyzer = InitAnalyzer()

	return
}