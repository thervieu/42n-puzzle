package models

type Puzzle struct {
	size        int
	numbers     []int
	pos_zero    int
	parent_move string
}

func InitializePuzzle(size int, order []int, pos_zero int) (p Puzzle) {
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

	if move == "U" {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero+p.size]
		p.numbers[p.pos_zero+p.size] = 0
		p.pos_zero = p.pos_zero + p.size
		p.parent_move = "U"
	}
	if move == "D" {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero-p.size]
		p.numbers[p.pos_zero-p.size] = 0
		p.pos_zero = p.pos_zero - p.size
		p.parent_move = "D"
	}
	if move == "L" {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero+1]
		p.numbers[p.pos_zero+1] = 0
		p.pos_zero = p.pos_zero + 1
		p.parent_move = "L"
	}
	if move == "R" {
		p.numbers[p.pos_zero] = p.numbers[p.pos_zero-1]
		p.numbers[p.pos_zero-1] = 0
		p.pos_zero = p.pos_zero - 1
		p.parent_move = "R"
	}
	return
}

func (p *Puzzle) next_states() []*Puzzle {
	next_states := []*Puzzle{}
	if p.pos_zero >= p.size && p.parent_move != "U" {
		copy := p.Copy()
		copy.move("D")
		next_states = append(next_states, copy)
	}
	if p.pos_zero <= (p.size*p.size)-1-p.size && p.parent_move != "D" {
		copy := p.Copy()
		copy.move("U")
		next_states = append(next_states, copy)
	}
	if p.pos_zero%p.size != 0 && p.parent_move != "L" {
		copy := p.Copy()
		copy.move("R")
		next_states = append(next_states, copy)
	}
	if (p.pos_zero+1)%p.size != 0 && p.parent_move != "R" {
		copy := p.Copy()
		copy.move("L")
		next_states = append(next_states, copy)
	}
	return (next_states)
}

func MakeGoal(size int) (goal []Point) {
	goal = make([]Point, size*size)
	puzzle := make([]int, size*size)
	for i := range goal {
		goal[i] = Point{-1, -1}
		puzzle[i] = -1
	}

	cur := 1
	x, y := 0, 0
	ix, iy := 1, 0
	for {
		puzzle[x+y*size] = cur
		goal[cur] = Point{x, y}

		if cur == 0 {
			break
		}
		cur += 1
		if x+ix == size || x+ix < 0 || (ix != 0 && puzzle[x+ix+y*size] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == size || y+iy < 0 || (iy != 0 && puzzle[x+(y+iy)*size] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
		if cur == size*size {
			cur = 0
		}
	}

	return
}

func MakeGoalArray(size int) (goal []int) {
	goalPoints := MakeGoal(size)
	goal = make([]int, size*size)
	for i := range goalPoints {
		goal[i] = goalPoints[i].x + (goalPoints[i].y * size)
	}
	return
}

// func (p *Puzzle) IsSolvable() (bool) {
// 	goal := MakeGoalArray(p.size)

// 	nbInversionsStart := numberInversions(p.numbers)
// 	nbInversionsGoal := numberInversions(goal)

// 	if (p.size % 2 == 0) { // also handle 0

// 	}

// 	return
// }
