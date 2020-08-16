package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_readFile(t *testing.T) {
	tests := []struct {
		name              string
		path              string
		expectedInputData inputData
		expectedErr       error
	}{
		{
			name: "test happy path",
			path: "data/input.json",
			expectedInputData: inputData{
				Dimensions: 4,
				Moves:      "DLUURR",
				Zombie:     zombie{XPos: 2, YPos: 1},
				Creatures: []creature{
					creature{XPos: 0, YPos: 1},
					creature{XPos: 1, YPos: 2},
					creature{XPos: 3, YPos: 1},
				},
			},
		},
		{
			name:        "test wrong path",
			path:        "test.json",
			expectedErr: errors.New("open test.json: no such file or directory"),
		},
	}
	for _, test := range tests {
		output, err := readFile(test.path)

		if !reflect.DeepEqual(test.expectedInputData, output) {
			t.Errorf("for %s, expected result %v, but got %v", test.name, test.expectedInputData, output)
		}

		if err != nil && test.expectedErr.Error() != err.Error() {
			t.Errorf("for %s, expected error %v, but got %v", test.name, test.expectedErr, err)
		}
	}
}
