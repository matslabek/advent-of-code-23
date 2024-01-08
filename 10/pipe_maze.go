package pipemaze

type Coordinate struct {
	X int
	Y int
}
type Pipe struct {
	C   Coordinate
	Dir rune
}

func PathFinder(inputs []string) int {
	li := len(inputs)
	S := Coordinate{}
	// Find where we start
	for i := 0; i < li; i++ {
		row := inputs[i]
		for j := 0; j < li; j++ {
			if row[j] == 'S' {
				S.X = j
				S.Y = i
				break
			}
		}
		if S.X != 0 {
			break
		}
	}
	startingPoints := findStartingDirections(inputs, S)
	counter := 1
	A := startingPoints[0]
	Z := startingPoints[1]
	for {
		switch A.Dir {
		case '^':
			A.C.Y--
		case 'v':
			A.C.Y++
		case '>':
			A.C.X++
		case '<':
			A.C.X--
		}
		switch inputs[A.C.Y][A.C.X] {
		case 'J':
			if A.Dir == 'v' {
				A.Dir = '<'
			} else {
				A.Dir = '^'
			}
		case 'F':
			if A.Dir == '^' {
				A.Dir = '>'
			} else {
				A.Dir = 'v'
			}
		case '7':
			if A.Dir == '>' {
				A.Dir = 'v'
			} else {
				A.Dir = '<'
			}
		case 'L':
			if A.Dir == 'v' {
				A.Dir = '>'
			} else {
				A.Dir = '^'
			}
		}

		switch Z.Dir {
		case '^':
			Z.C.Y--
		case 'v':
			Z.C.Y++
		case '>':
			Z.C.X++
		case '<':
			Z.C.X--
		}
		switch inputs[Z.C.Y][Z.C.X] {
		case 'J':
			if Z.Dir == 'v' {
				Z.Dir = '<'
			} else {
				Z.Dir = '^'
			}
		case 'F':
			if Z.Dir == '^' {
				Z.Dir = '>'
			} else {
				Z.Dir = 'v'
			}
		case '7':
			if Z.Dir == '>' {
				Z.Dir = 'v'
			} else {
				Z.Dir = '<'
			}
		case 'L':
			if Z.Dir == 'v' {
				Z.Dir = '>'
			} else {
				Z.Dir = '^'
			}
		}

		counter++
		// Check if we are in the same point
		if A.C.X == Z.C.X && A.C.Y == Z.C.Y {
			return counter
		}
	}
}

func findStartingDirections(plane []string, S Coordinate) []Pipe {
	pipes := make([]Pipe, 0)
	//Check left
	if plane[S.Y][S.X-1] == 'F' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X - 1, Y: S.Y},
			Dir: 'v',
		})
	} else if plane[S.Y][S.X-1] == '-' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X - 1, Y: S.Y},
			Dir: '<',
		})
	} else if plane[S.Y][S.X-1] == 'L' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X - 1, Y: S.Y},
			Dir: '^',
		})
	}
	// Check right
	if plane[S.Y][S.X+1] == '-' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X + 1, Y: S.Y},
			Dir: '>',
		})
	} else if plane[S.Y][S.X+1] == 'J' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X + 1, Y: S.Y},
			Dir: '^',
		})
	} else if plane[S.Y][S.X+1] == '7' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X + 1, Y: S.Y},
			Dir: 'v',
		})
	}
	// Check up
	if plane[S.Y-1][S.X] == '|' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X, Y: S.Y - 1},
			Dir: '^',
		})
	} else if plane[S.Y-1][S.X] == 'F' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X, Y: S.Y - 1},
			Dir: '>',
		})
	} else if plane[S.Y-1][S.X] == '7' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X, Y: S.Y - 1},
			Dir: '<',
		})
	}
	//Check down
	if plane[S.Y+1][S.X] == '|' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X, Y: S.Y + 1},
			Dir: 'v',
		})
	} else if plane[S.Y+1][S.X] == 'L' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X, Y: S.Y + 1},
			Dir: '>',
		})
	} else if plane[S.Y+1][S.X] == 'J' {
		pipes = append(pipes, Pipe{
			C:   Coordinate{X: S.X, Y: S.Y + 1},
			Dir: '<',
		})
	}
	return pipes
}
