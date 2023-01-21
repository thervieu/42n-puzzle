package server

import (
	"fmt"
	// "log"
	// "time"
	"encoding/json"
	// "strings"

	"net/http"
	"sync"

	// "io/ioutil"

	"github.com/thervieu/42n-puzzle/astar"
	"github.com/thervieu/42n-puzzle/models"
)

func solveRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode the data from JSON
	decoder := json.NewDecoder(r.Body)
	data := models.RequestSolveData{}

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
	puzzle, err := data.ExtractPuzzle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	responseData := make([]models.ResponseSolveData, 6)
	if data.Option == 6 {
		var wg sync.WaitGroup
		for i := 0; i < 6; i++ {
			wg.Add(1)
			go func(i int, puzzle models.Puzzle) {
				defer wg.Done()
				responseData[i] = astar.Solve(i, puzzle)
			}(i, puzzle)
		}
		wg.Wait()

	} else {
		responseData = make([]models.ResponseSolveData, 1)
		responseData[0] = astar.Solve(data.Option, puzzle)
	}

	// start := time.Now()
	// solved := models.Solve(context)
	result := "bonjour"

	// elapsed_time := time.Now().Sub(start).Milliseconds()
	// _json, err := solved.ToJSON(elapsed_time)

	// if err != nil {
	// 	// If the context computation failed
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	fmt.Fprintf(w, result)
}
