package models

import (

)

type Analyzer struct {
	time_complexity int
	space_complexity int
}



func InitAnalyzer() (Analyzer) {
	return Analyzer{0, 0}
}
