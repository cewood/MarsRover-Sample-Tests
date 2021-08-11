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
