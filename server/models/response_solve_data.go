package models

import (
)

type SolveComplexityData struct {
	Time        		int64	`json:"time"`
	Time_complexity		int		`json:"time_complexity"`
	Space_complexity	int		`json:"space_complexity"`
}

type ResponseSolveData struct {
	Complexity	SolveComplexityData	`json:"options"`
	States		[]string			`json:"states"`
	Analyzer	Analyzer			`json:"analyzer"`
}

func (r *ResponseSolveData) computeResponseSolve(ctx *Context, time int64, time_complexity int, space_complexity int) {
	r.Complexity = SolveComplexityData{
		Time: time,
		Time_complexity: time_complexity,
		Space_complexity: space_complexity,
	}
	// r.States = ctx.ComputeTimeline()

	return 
}