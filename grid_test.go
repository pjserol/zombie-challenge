package main

import (
	"reflect"
	"testing"
)

func Test_newGrid(t *testing.T) {
	tests := []struct {
		name         string
		dimensions   int
		moves        string
		zombie       zombie
		creatures    []creature
		expectedGrid Grid
	}{
		{
			name:       "test happy path",
			dimensions: 4,
			moves:      "DLUURR",
			zombie:     zombie{XPos: 2, YPos: 1},
			creatures: []creature{
				creature{XPos: 0, YPos: 1},
				creature{XPos: 1, YPos: 2},
				creature{XPos: 3, YPos: 1},
			},
			expectedGrid: Grid{
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
		},
		{
			name: "test empty value",
			expectedGrid: Grid{
				Zombies: []zombie{zombie{}},
			},
		},
	}
	for _, test := range tests {
		output := newGrid(test.dimensions, test.moves, test.zombie, test.creatures)

		if !reflect.DeepEqual(test.expectedGrid, output) {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedGrid, output)
		}
	}
}
