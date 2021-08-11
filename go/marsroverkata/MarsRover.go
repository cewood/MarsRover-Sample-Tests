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

func (r MarsRover) acceptCommands(commands []Command) {

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
