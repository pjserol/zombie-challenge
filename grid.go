package main

// Grid data of the programs
type Grid struct {
	Dimensions int
	Moves      string
	Zombies    []zombie
	Creatures  []creature
	Scores     int
}

func newGrid(dimensions int, moves string, z zombie, creatures []creature) Grid {
	return Grid{
		Dimensions: dimensions,
		Moves:      moves,
		Zombies:    []zombie{z},
		Creatures:  creatures,
	}
}
