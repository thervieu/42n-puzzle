package models

import (
)

const (
	euclidean = iota // 0
	manhattan = iota // 1
	hamming = iota // 2
)

type Options struct {
	search string
	heuristic string
}
