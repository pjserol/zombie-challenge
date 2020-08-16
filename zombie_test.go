package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_nextPosition(t *testing.T) {
	tests := []struct {
		name             string
		x, y, dimensions int
		zombie           zombie
		expectedZombie   zombie
		expectedErr      error
	}{
		{
			name:           "test happy path",
			x:              1,
			y:              1,
			dimensions:     4,
			zombie:         zombie{XPos: 1, YPos: 1},
			expectedZombie: zombie{XPos: 2, YPos: 2},
		},
		{
			name:           "test happy path x max",
			x:              1,
			y:              0,
			dimensions:     4,
			zombie:         zombie{XPos: 3, YPos: 1},
			expectedZombie: zombie{XPos: 0, YPos: 1},
		},
		{
			name:           "test happy path x min",
			x:              -1,
			y:              0,
			dimensions:     4,
			zombie:         zombie{XPos: 0, YPos: 1},
			expectedZombie: zombie{XPos: 3, YPos: 1},
		},
		{
			name:           "test happy path y max",
			x:              0,
			y:              1,
			dimensions:     4,
			zombie:         zombie{XPos: 1, YPos: 3},
			expectedZombie: zombie{XPos: 1, YPos: 0},
		},
		{
			name:           "test happy path y min",
			x:              0,
			y:              -1,
			dimensions:     4,
			zombie:         zombie{XPos: 1, YPos: 0},
			expectedZombie: zombie{XPos: 1, YPos: 3},
		},
		{
			name:        "test wrong x",
			x:           2,
			y:           1,
			dimensions:  3,
			expectedErr: errors.New("x not allowed"),
		},
		{
			name:        "test wrong y",
			x:           1,
			y:           2,
			dimensions:  3,
			expectedErr: errors.New("y not allowed"),
		},
	}
	for _, test := range tests {
		err := test.zombie.nextPosition(test.x, test.y, test.dimensions)

		if !reflect.DeepEqual(test.expectedZombie, test.zombie) {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedZombie, test.zombie)
		}

		if err != nil && test.expectedErr.Error() != err.Error() {
			t.Errorf("for %s, expected error %v, but got %v", test.name, test.expectedErr, err)
		}
	}
}

func Test_isInfected(t *testing.T) {
	tests := []struct {
		name         string
		creature     creature
		zombie       zombie
		expectedBool bool
	}{
		{
			name:         "test happy path",
			zombie:       zombie{XPos: 3, YPos: 0},
			creature:     creature{XPos: 3, YPos: 0},
			expectedBool: true,
		},
		{
			name:         "test happy path zombie creature already infected",
			zombie:       zombie{XPos: 3, YPos: 0},
			creature:     creature{XPos: 3, YPos: 0, infected: true},
			expectedBool: false,
		},
		{
			name:         "test happy path with different position",
			zombie:       zombie{XPos: 3, YPos: 0},
			creature:     creature{XPos: 2, YPos: 1},
			expectedBool: false,
		},
	}
	for _, test := range tests {
		output := test.zombie.isInfected(test.creature)

		if test.expectedBool != output {
			t.Errorf("for %s, expected result %t, but got %t", test.name, test.expectedBool, output)
		}
	}
}
