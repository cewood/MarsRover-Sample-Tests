package marsrover

import "fmt"

type Coordinates struct {
	x int
	y int
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

func (d Direction) String() string {
	return [...]string{"N", "E", "S", "W"}[d]
}

type Command int

const (
	B Command = iota
	F
	L
	R
)

func (c Command) String() string {
	return [...]string{"B", "F", "L", "R"}[c]
}

type Obstacle struct {
	position Coordinates
}

type Plateau struct {
	maxX      int
	maxY      int
	obstacles []Obstacle
}

type Status int

const (
	OK Status = iota
	NOK
)

func (s Status) String() string {
	return [...]string{"OK", "NOK"}[s]
}

type MarsRover struct {
	plateau  Plateau
	heading  Direction
	position Coordinates
	status   Status
}

func (r MarsRover) printResultHeader() string {
	buf := make([]byte, 0)

	// Add our line padding for alignment
	buf = append(buf, []byte("  ")...)

	// Print the column numbers as a header
	for x := 0; x < r.plateau.maxX; x++ {
		buf = append(buf, []byte(fmt.Sprintf(" %d", x))...)
	}

	// Add our line ending
	buf = append(buf, []byte("\n")...)

	return string(buf)
}

func (r MarsRover) printResultRow(row int, obstacles map[string]bool) string {
	buf := make([]byte, 0)

	// Add our line padding for alignment
	buf = append(buf, []byte(fmt.Sprintf("%d |", row))...)

	obs := make([]Obstacle, 0)

	// Narrow the list of obstacles to check
	for _, val := range r.plateau.obstacles {
		if val.position.y == row {
			obs = append(obs, val)
		}
	}

	for x := 0; x < r.plateau.maxX; x++ {
		// check if there is a rover here
		if r.position.y == row && r.position.x == x {
			// determine the direction to print
			switch heading := r.heading; heading {
			case N:
				buf = append(buf, []byte("ᐱ|")...)
			case E:
				buf = append(buf, []byte("ᐳ|")...)
			case S:
				buf = append(buf, []byte("ᐯ|")...)
			case W:
				buf = append(buf, []byte("ᐸ|")...)
			}
		} else if _, ok := obstacles[fmt.Sprintf("%d:%d", x, row)]; ok {
			// check if there is a obstacle here
			buf = append(buf, []byte("x|")...)
		} else {
			// default empty tile case
			buf = append(buf, []byte(" |")...)
		}
	}

	// Add our line ending
	buf = append(buf, []byte(fmt.Sprintf(" %d\n", row))...)

	return string(buf)
}

func (r MarsRover) printResult() string {
	buf := make([]byte, 0)
	obstacles := make(map[string]bool, 0)

	// Index the obstacles for faster checking
	for _, val := range r.plateau.obstacles {
		obstacles[fmt.Sprintf("%d:%d", val.position.x, val.position.y)] = true
	}

	// Append the header at the top
	buf = append(buf, []byte(r.printResultHeader())...)

	// Printing is top to bottom, left to right
	for y := r.plateau.maxY - 1; y >= 0; y-- {
		buf = append(buf, []byte(r.printResultRow(y, obstacles))...)
	}

	// Append the header at the bottom
	buf = append(buf, []byte(r.printResultHeader())...)

	return string(buf)
}

func (r MarsRover) PrintResult() {
	fmt.Print(r.printResult())
}

func (r *MarsRover) turnLeft() {
	switch {
	case r.heading == N:
		// Turning from N to W, hence reset/jump to 3
		r.heading = 3
	default:
		// Default, decrement the value
		r.heading--
	}
}

func (r MarsRover) currentLocation() interface{} {
	return fmt.Sprintf("%v %v %v", r.position.x, r.position.y, r.heading)
}

func (r *MarsRover) acceptCommands(commands []Command) {
	for _, c := range commands {
		if r.status == NOK {
			return
		}

		switch {
		case c == B:
			r.backward()
		case c == F:
			r.forward()
		case c == L:
			r.turnLeft()
		case c == R:
			r.turnRight()
		}
	}
}

func (r MarsRover) coordinates() Coordinates {
	return r.position
}

func (r *MarsRover) move(direction Direction) {
	pos := r.position

	switch d := direction; d {
	case N:
		if r.position.y == r.plateau.maxY {
			pos.y = 0
		} else {
			pos.y++
		}
	case S:
		if r.position.y == 0 {
			pos.y = r.plateau.maxY
		} else {
			pos.y--
		}
	case E:
		if r.position.x == r.plateau.maxX {
			pos.x = 0
		} else {
			pos.x++
		}
	case W:
		if r.position.x == 0 {
			pos.x = r.plateau.maxX
		} else {
			pos.x--
		}
	}

	// check if the proposed location has an obstacle
	for _, obstacle := range r.plateau.obstacles {
		if pos == obstacle.position {
			r.status = NOK
			return
		}
	}

	r.position = pos
	r.status = OK
}

func (r *MarsRover) forward() {
	r.move(r.heading)
}

func (r *MarsRover) backward() {
	// Nasty inline slice of opposite directions
	r.move([...]Direction{S, W, N, E}[r.heading])
}

func (r *MarsRover) turnRight() {
	switch {
	case r.heading == 3:
		// Turning from W to N, hence reset/jump to 0
		r.heading = 0
	default:
		// Default, increment the value
		r.heading++
	}
}
