package models

import (
)

type SolveComplexityData struct {
	Time        		int64	`json:"time"`
	time_complexity		int		`json:"time_complexity"`
	space_complexity	int		`json:"space_complexity"`
}

type ResponseSolveData struct {
	Complexity	SolveComplexityData	`json:"options"`
	States		[]string			`json:"states"`
	Analyzer	Analyzer			`json:"analyzer"`
}

func (r *ResponseSolveData) computeResponseSolve(ctx *Context, time int64, time_complexity int, space_complexity int) {
	r.Complexity = SolveComplexityData{
		Time: time,
		time_complexity: time_complexity,
		space_complexity: space_complexity,
	}
	r.States = ctx.ComputeTimeline()

	return 
}