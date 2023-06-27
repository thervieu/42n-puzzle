package server

import (
	"fmt"

	// "log"
	// "time"
	"encoding/json"
	// "strings"

	"net/http"

	// "io/ioutil"

	"github.com/thervieu/42n-puzzle/solver"
)

func solveRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode the data from JSON
	decoder := json.NewDecoder(r.Body)
	data := RequestSolveData{}

	if err := decoder.Decode(&data); err != nil {
		// If JSON unmarshalling failed
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := data.Sanitize(); err != nil {
		// If the data is corrupted
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Compute a new context from the data
	size, initPuzzle := data.ExtractPuzzle()
	heuristic, search := data.ExtractOptions()

	solveData, err := solver.Solve(size, initPuzzle, heuristic, search, false)
	if err != nil {
		// If the data is corrupted
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_json, err := solveData.ToJSON()
	if err != nil {
		// If the context computation failed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, _json)
}
