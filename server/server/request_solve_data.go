package server

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type RequestSolveData struct {
	Heuristic string `json:"heuristic"`
	Search    string `json:"search"`
	Puzzle    []int  `json:"puzzle"`
}

func SanitizeHeuristic(heuristic string) (err error) {
	if heuristic != "hamming" && heuristic != "manhattan" && heuristic != "euclidean" {
		return errors.New("Bad Request: heuristic has a wrong value")
	}
	return nil
}

func SanitizeSearch(search string) (err error) {
	if search != "uniform" && search != "greedy" {
		return errors.New("Bad Request: search has a wrong value")
	}
	return nil
}

func VerifValuesSlice(arr []int, typeError bool) error {

	begError := ""
	if typeError == true {
		begError = "Bad Request: puzzle: "
	} else {
		begError = "error: file: "
	}
	slice := make([]int, len(arr))
	for _, v := range arr {
		if v < 0 {
			return errors.New(begError + "negative number")
		}
		if v >= len(arr) {
			return errors.New(begError + "value bigger than its length")
		}
		if slice[v] == 1 {
			return errors.New(begError + "value " + strconv.Itoa(v) + " duplicate ")
		}
		slice[v] = 1
	}
	return nil
}

func (r *RequestSolveData) SanitizePuzzle() error {
	int_sqrt := int(math.Sqrt(float64(len(r.Puzzle))))
	if int_sqrt*int_sqrt != len(r.Puzzle) {
		return errors.New("Bad Request: Puzzle length is not a squared number")
	}
	if int_sqrt*int_sqrt != len(r.Puzzle) {
		return errors.New("Bad Request: Puzzle length is not a squared number")
	}

	return VerifValuesSlice(r.Puzzle, true)
}

func (r *RequestSolveData) Sanitize() error {
	if err := SanitizeHeuristic(r.Heuristic); err != nil {
		return err
	}
	if err := SanitizeSearch(r.Search); err != nil {
		return err
	}
	if err := r.SanitizePuzzle(); err != nil {
		return err
	}
	return nil
}

func generateNumbersSlice(strPuzzle string) (slice []int) {
	splittedStr := strings.Split(strPuzzle, ",")
	slice = make([]int, len(splittedStr))
	for i, str := range splittedStr {
		_int, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		slice[i] = _int
	}
	return
}

func index(needle int, haystack []int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func (r *RequestSolveData) ExtractPuzzle() (size int, puzzle []int) {
	size = int(math.Sqrt(float64((len(r.Puzzle)))))
	puzzle = r.Puzzle
	return
}

func (r *RequestSolveData) ExtractOptions() (heuristic string, search string) {
	heuristic = r.Heuristic
	search = r.Search
	return
}
