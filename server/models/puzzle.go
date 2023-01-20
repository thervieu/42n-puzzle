package models

import (

)

type Puzzle struct {
	size int
	numbers []int
	pos_zero int
	parent_move string
}

func InitializePuzzle(size int, order []int, pos_zero int) (p *Puzzle)  {
	p.size = size
	p.numbers = order
	p.pos_zero = pos_zero
	p.parent_move = ""

	return
}

func (p *Puzzle) Copy() *Puzzle {
	newPuzzle := &Puzzle{}
	newPuzzle.parent_move = p.parent_move
	newPuzzle.size = p.size
	newPuzzle.numbers = p.numbers

	return (newPuzzle)
}

func (p *Puzzle) move(move string) {
	// swap numbers positions
	// set pos_zero to new index
	// set parent_move to move

	if (move == "U") {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero + p.size]
		p.numbers[p.pos_zero + p.size] = 0
		p.pos_zero = p.pos_zero + p.size
		p.parent_move = "U"
	}
	if (move == "D") {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero - p.size]
		p.numbers[p.pos_zero - p.size] = 0
		p.pos_zero = p.pos_zero - p.size
		p.parent_move = "D"
	}
	if (move == "L") {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero + 1]
		p.numbers[p.pos_zero + 1] = 0
		p.pos_zero = p.pos_zero + 1
		p.parent_move = "L"
	}
	if (move == "R") {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero - 1]
		p.numbers[p.pos_zero - 1] = 0
		p.pos_zero = p.pos_zero - 1
		p.parent_move = "R"
	}
	return
}

func (p *Puzzle) next_states() []*Puzzle {
	next_states := []*Puzzle{}
	if (p.pos_zero >= p.size && p.parent_move != "U") {
		copy := p.Copy()
		copy.move("D")
		next_states = append(next_states, copy)
	}
	if (p.pos_zero <= (p.size * p.size) - 1 - p.size && p.parent_move != "D") {
		copy := p.Copy()
		copy.move("U")
		next_states = append(next_states, copy)
	}
	if (p.pos_zero % p.size != 0 && p.parent_move != "L") {
		copy := p.Copy()
		copy.move("R")
		next_states = append(next_states, copy)
	}
	if ((p.pos_zero + 1) % p.size != 0 && p.parent_move != "R") {
		copy := p.Copy()
		copy.move("L")
		next_states = append(next_states, copy)
	}
	return (next_states)
}
