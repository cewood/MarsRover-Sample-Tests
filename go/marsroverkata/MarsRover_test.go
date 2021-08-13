package marsrover

import (
	"testing"
)

func TestMarsRover_turnRight(t *testing.T) {
	tests := []struct {
		name     string
		heading  Direction
		expected Direction
	}{
		{
			"north to east",
			N,
			E,
		},
		{
			"west to north",
			W,
			N,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plateau := Plateau{maxX: 5, maxY: 5}
			startingPosition := Coordinates{1, 2}

			r := &MarsRover{
				plateau:  plateau,
				heading:  tt.heading,
				position: startingPosition,
			}
			r.turnRight()

			if r.heading != tt.expected {
				t.Errorf("Expected '%v', but got '%v'\n", tt.expected, r.heading)
			}
		})
	}
}

func TestMarsRover_turnLeft(t *testing.T) {
	tests := []struct {
		name     string
		heading  Direction
		expected Direction
	}{
		{
			"north to west",
			N,
			W,
		},
		{
			"south to east",
			S,
			E,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plateau := Plateau{maxX: 5, maxY: 5}
			startingPosition := Coordinates{1, 2}

			r := &MarsRover{
				plateau:  plateau,
				heading:  tt.heading,
				position: startingPosition,
			}
			r.turnLeft()

			if r.heading != tt.expected {
				t.Errorf("Expected '%v', but got '%v'\n", tt.expected, r.heading)
			}
		})
	}
}

func TestMarsRover_backward(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	tests := []struct {
		name     string
		heading  Direction
		coords   Coordinates
		expected Coordinates
	}{
		{
			"normal move, N on Y axis",
			N,
			Coordinates{0, 0},
			Coordinates{0, 5},
		},
		{
			"wrap move, N on Y axis",
			N,
			Coordinates{0, 5},
			Coordinates{0, 4},
		},
		{
			"normal move, S on Y axis",
			S,
			Coordinates{0, 5},
			Coordinates{0, 0},
		},
		{
			"wrap move, S on Y axis",
			S,
			Coordinates{0, 0},
			Coordinates{0, 1},
		},
		{
			"normal move, E on X axis",
			E,
			Coordinates{0, 0},
			Coordinates{5, 0},
		},
		{
			"wrap move, E on X axis",
			E,
			Coordinates{5, 0},
			Coordinates{4, 0},
		},
		{
			"normal move, W on X axis",
			W,
			Coordinates{5, 0},
			Coordinates{0, 0},
		},
		{
			"wrap move, W on X axis",
			W,
			Coordinates{0, 0},
			Coordinates{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MarsRover{
				plateau:  plateau,
				heading:  tt.heading,
				position: tt.coords,
			}
			r.backward()

			if r.position != tt.expected {
				t.Errorf("Expected '%v', but got '%v'\n", tt.expected, r.position)
			}
		})
	}
}

func TestMarsRover_forward(t *testing.T) {
	plateau := Plateau{maxX: 5, maxY: 5}
	tests := []struct {
		name     string
		heading  Direction
		coords   Coordinates
		expected Coordinates
	}{
		{
			"normal move, N on Y axis",
			N,
			Coordinates{0, 0},
			Coordinates{0, 1},
		},
		{
			"wrap move, N on Y axis",
			N,
			Coordinates{0, 5},
			Coordinates{0, 0},
		},
		{
			"normal move, S on Y axis",
			S,
			Coordinates{0, 5},
			Coordinates{0, 4},
		},
		{
			"wrap move, S on Y axis",
			S,
			Coordinates{0, 0},
			Coordinates{0, 5},
		},
		{
			"normal move, E on X axis",
			E,
			Coordinates{0, 0},
			Coordinates{1, 0},
		},
		{
			"wrap move, E on X axis",
			E,
			Coordinates{5, 0},
			Coordinates{0, 0},
		},
		{
			"normal move, W on X axis",
			W,
			Coordinates{5, 0},
			Coordinates{4, 0},
		},
		{
			"wrap move, W on X axis",
			W,
			Coordinates{0, 0},
			Coordinates{5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MarsRover{
				plateau:  plateau,
				heading:  tt.heading,
				position: tt.coords,
			}
			r.forward()

			if r.position != tt.expected {
				t.Errorf("Expected '%v', but got '%v'\n", tt.expected, r.position)
			}
		})
	}
}

func TestMarsRover_move(t *testing.T) {
	tests := []struct {
		name      string
		heading   Direction
		obstacles []Obstacle
		coords    Coordinates
		expected  Coordinates
	}{
		{
			"normal move, N on Y axis",
			N,
			[]Obstacle{},
			Coordinates{0, 0},
			Coordinates{0, 1},
		},
		{
			"wrap move, N on Y axis",
			N,
			[]Obstacle{},
			Coordinates{0, 5},
			Coordinates{0, 0},
		},
		{
			"normal move, S on Y axis",
			S,
			[]Obstacle{},
			Coordinates{0, 5},
			Coordinates{0, 4},
		},
		{
			"wrap move, S on Y axis",
			S,
			[]Obstacle{},
			Coordinates{0, 0},
			Coordinates{0, 5},
		},
		{
			"normal move, E on X axis",
			E,
			[]Obstacle{},
			Coordinates{0, 0},
			Coordinates{1, 0},
		},
		{
			"wrap move, E on X axis",
			E,
			[]Obstacle{},
			Coordinates{5, 0},
			Coordinates{0, 0},
		},
		{
			"normal move, W on X axis",
			W,
			[]Obstacle{},
			Coordinates{5, 0},
			Coordinates{4, 0},
		},
		{
			"wrap move, W on X axis",
			W,
			[]Obstacle{},
			Coordinates{0, 0},
			Coordinates{5, 0},
		},
		{
			"normal move, N on Y axis with obstacle",
			N,
			[]Obstacle{Obstacle{Coordinates{0, 1}}},
			Coordinates{0, 0},
			Coordinates{0, 0},
		},
		{
			"wrap move, N on Y axis with obstacle",
			N,
			[]Obstacle{Obstacle{Coordinates{0, 0}}},
			Coordinates{0, 5},
			Coordinates{0, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plateau := Plateau{maxX: 5, maxY: 5, obstacles: tt.obstacles}
			r := &MarsRover{
				plateau:  plateau,
				heading:  tt.heading,
				position: tt.coords,
			}
			r.move(r.heading)

			if r.position != tt.expected {
				t.Errorf("Expected '%v', but got '%v'\n", tt.expected, r.position)
			}
		})
	}
}

func TestMarsRover_acceptCommands(t *testing.T) {
	tests := []struct {
		name       string
		obstacles  []Obstacle
		commands   []Command
		expHeading Direction
		expCoords  Coordinates
		expStatus  Status
	}{
		{
			"normal case, moving in single direction",
			[]Obstacle{},
			[]Command{F, F},
			N,
			Coordinates{0, 2},
			OK,
		},
		{
			"obstruction case, moving in single direction",
			[]Obstacle{{Coordinates{0, 2}}},
			[]Command{F, F},
			N,
			Coordinates{0, 1},
			NOK,
		},
		{
			"normal case, moving multiple directions",
			[]Obstacle{},
			[]Command{F, F, R, F, F, R, F, F, R, F, F},
			W,
			Coordinates{0, 0},
			OK,
		},
		{
			"obstruction case, moving multiple directions",
			[]Obstacle{{Coordinates{1, 0}}},
			[]Command{F, F, R, F, F, R, F, F, R, F, F},
			W,
			Coordinates{2, 0},
			NOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plateau := Plateau{maxX: 5, maxY: 5, obstacles: tt.obstacles}
			r := &MarsRover{
				plateau:  plateau,
				heading:  N,
				position: Coordinates{0, 0},
			}
			r.acceptCommands(tt.commands)

			if r.position != tt.expCoords || r.heading != tt.expHeading || r.status != tt.expStatus {
				t.Errorf("Expected '%v' '%v' '%v', but got '%v' '%v' '%v'\n", tt.expCoords, tt.expHeading, tt.expStatus, r.position, r.heading, r.status)
			}
		})
	}
}

func TestMarsRover_printResultHeader(t *testing.T) {
	tests := []struct {
		name  string
		input Plateau
		want  string
	}{
		{
			"5x5 plateau",
			Plateau{maxX: 5, maxY: 5, obstacles: []Obstacle{}},
			"   0 1 2 3 4\n",
		},
		{
			"10x10 plateau",
			Plateau{maxX: 10, maxY: 10, obstacles: []Obstacle{}},
			"   0 1 2 3 4 5 6 7 8 9\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MarsRover{
				plateau:  tt.input,
				heading:  N,
				position: Coordinates{0, 0},
				status:   OK,
			}

			if got := r.printResultHeader(); got != tt.want {
				t.Errorf("MarsRover.printResultHeader() = '%v', want '%v'\n", got, tt.want)
			}
		})
	}
}

func TestMarsRover_printResultRow(t *testing.T) {
	type input struct {
		plateau   Plateau
		heading   Direction
		position  Coordinates
		obstacles map[string]bool
	}
	tests := []struct {
		name  string
		input input
		want  string
	}{
		{
			"5x5 plateau",
			input{
				plateau:  Plateau{maxX: 5, maxY: 5, obstacles: []Obstacle{}},
				heading:  N,
				position: Coordinates{0, 0},
				obstacles: map[string]bool{
					"3:0": true,
				},
			},
			"0 |ᐱ| | |x| | 0\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &MarsRover{
				plateau:  tt.input.plateau,
				heading:  tt.input.heading,
				position: tt.input.position,
				status:   OK,
			}
			if got := r.printResultRow(0, tt.input.obstacles); got != tt.want {
				t.Errorf("MarsRover.printResultRow() = '%v', want '%v'\n", got, tt.want)
			}
		})
	}
}

func TestMarsRover_printResult(t *testing.T) {
	tests := []struct {
		name  string
		input *MarsRover
		want  string
	}{
		{
			"5x5 plateau",
			&MarsRover{
				plateau:  Plateau{maxX: 5, maxY: 5, obstacles: []Obstacle{{Coordinates{2, 2}}}},
				heading:  N,
				position: Coordinates{0, 0},
				status:   OK,
			},
			`   0 1 2 3 4
4 | | | | | | 4
3 | | | | | | 3
2 | | |x| | | 2
1 | | | | | | 1
0 |ᐱ| | | | | 0
   0 1 2 3 4
`,
		},
		{
			"10x10 plateau",
			&MarsRover{
				plateau:  Plateau{maxX: 10, maxY: 10, obstacles: []Obstacle{{Coordinates{8, 8}}}},
				heading:  N,
				position: Coordinates{0, 0},
				status:   OK,
			},
			`   0 1 2 3 4 5 6 7 8 9
9 | | | | | | | | | | | 9
8 | | | | | | | | |x| | 8
7 | | | | | | | | | | | 7
6 | | | | | | | | | | | 6
5 | | | | | | | | | | | 5
4 | | | | | | | | | | | 4
3 | | | | | | | | | | | 3
2 | | | | | | | | | | | 2
1 | | | | | | | | | | | 1
0 |ᐱ| | | | | | | | | | 0
   0 1 2 3 4 5 6 7 8 9
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.printResult(); got != tt.want {
				t.Errorf("MarsRover.printResult() = '%v', want '%v'\n", got, tt.want)
			}
		})
	}
}
