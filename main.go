package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Start program...")

	if len(os.Args) != 2 {
		panic("Please specify an input path")
	}

	input, err := readFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	grid := newGrid(input.Dimensions, input.Moves, input.Zombie, input.Creatures)

	res := runProgram(&grid)

	log.Printf("Result: %v", res)

	if err := writeFile(res, "data/response.json"); err != nil {
		fmt.Println("Error:", err)
	}
}

func runProgram(grid *Grid) outputData {
	position := 0
	for {
		grid.moveZombie(position)

		position++

		if position > grid.Scores {
			break
		}
	}

	return outputData{
		Scores:  grid.Scores,
		Zombies: grid.Zombies,
	}
}

func convertMoveToCoordinates(move string) (x, y int, err error) {
	switch move {
	case "U":
		return 0, -1, nil
	case "D":
		return 0, 1, nil
	case "L":
		return -1, 0, nil
	case "R":
		return 1, 0, nil
	}

	return 0, 0, fmt.Errorf("unknown move '%s'", move)
}
