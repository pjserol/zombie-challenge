package main

type inputData struct {
	Dimensions int        `json:"dimensions"`
	Zombie     zombie     `json:"zombie"`
	Creatures  []creature `json:"creature"`
	Moves      string     `json:"moves"`
}

type creature struct {
	XPos     int `json:"x"`
	YPos     int `json:"y"`
	infected bool
}

type outputData struct {
	Scores  int
	Zombies []zombie `json:"zombie"`
}
