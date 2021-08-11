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
