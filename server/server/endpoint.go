package server

import (
	"fmt"
	// "log"
	"time"
	"encoding/json"
	// "strings"
	
	"net/http"
	// "io/ioutil"

	// "github.com/thervieu/42n-puzzle/models"
)

func solve(w http.ResponseWriter, r *http.Request) {
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
	context, err := data.ComputeContext()

	if err != nil {
		// If the context computation failed
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	start := time.Now()
	solved := models.Solve(context)
	result := "bonjour"

	elapsed_time := time.Now().Sub(start).Milliseconds()
	_json, err := solved.ToJSON(elapsed_time)
	
	if err != nil {
		// If the context computation failed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, result)
}