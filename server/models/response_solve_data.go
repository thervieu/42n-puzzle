package models

import (
)

type ResponseSolveOptionsData struct {
	Time        		int64	`json:"time"`
	time_complexity		int		`json:"time_complexity"`
	space_complexity	int		`json:"space_complexity"`
}

type ResponseSolveData struct {
	Complexity	ResponseSolveOptionsData	`json:"options"`
	States						[]string	`json:"states"`
}

func (r *ResponseSolveData) computeResponseSolve(context *Context, time int64, time_complexity int, space_complexity int) {
	r.Complexity = ResponseSolveOptionsData{
		Time: time,
		time_complexity: time_complexity,
		space_complexity: space_complexity,
	}
	r.States = c.ComputeTimeline()

	return 
}