package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_runProgram(t *testing.T) {
	tests := []struct {
		name           string
		grid           *Grid
		expectedOutput outputData
	}{
		{
			name: "test happy path",
			grid: &Grid{
				Dimensions: 4,
				Moves:      "DLUURR",
				Zombies: []zombie{
					zombie{XPos: 2, YPos: 1},
				},
				Creatures: []creature{
					creature{XPos: 0, YPos: 1},
					creature{XPos: 1, YPos: 2},
					creature{XPos: 3, YPos: 1},
				},
			},
			expectedOutput: outputData{
				Scores: 3,
				Zombies: []zombie{
					zombie{XPos: 3, YPos: 0},
					zombie{XPos: 2, YPos: 1},
					zombie{XPos: 1, YPos: 0},
					zombie{XPos: 0, YPos: 0},
				},
			},
		},
	}
	for _, test := range tests {
		output := runProgram(test.grid)

		if !reflect.DeepEqual(test.expectedOutput, output) {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, output)
		}
	}
}

func Test_convertMoveToCoordinates(t *testing.T) {
	tests := []struct {
		name        string
		move        string
		expectedX   int
		expectedY   int
		expectedErr error
	}{
		{
			name:      "test happy path U",
			move:      "U",
			expectedX: 0,
			expectedY: -1,
		},
		{
			name:      "test happy path D",
			move:      "D",
			expectedX: 0,
			expectedY: 1,
		},
		{
			name:      "test happy path L",
			move:      "L",
			expectedX: -1,
			expectedY: 0,
		},
		{
			name:      "test happy path R",
			move:      "R",
			expectedX: 1,
			expectedY: 0,
		},
		{
			name:        "test unknown move",
			move:        "A",
			expectedErr: errors.New("unknown move 'A'"),
		},
	}
	for _, test := range tests {
		x, y, err := convertMoveToCoordinates(test.move)

		if test.expectedX != x || test.expectedY != y {
			t.Errorf("for %s, expected result (%d,%d), but got (%d,%d)", test.name, test.expectedX, test.expectedY, x, y)
		}

		if err != nil && test.expectedErr.Error() != err.Error() {
			t.Errorf("for %s, expected error %v, but got %v", test.name, test.expectedErr, err)
		}
	}
}
