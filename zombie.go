package main

import (
	"errors"
	"log"
)

type zombie struct {
	XPos int `json:"x"`
	YPos int `json:"y"`
}

// move from one position
func (z *zombie) nextPosition(x, y int, dimensions int) error {
	if !(x == -1 || x == 0 || x == 1) {
		return errors.New("x not allowed")
	}

	if !(y == -1 || y == 0 || y == 1) {
		return errors.New("y not allowed")
	}

	z.XPos = z.XPos + x
	if z.XPos < 0 {
		z.XPos = dimensions - 1
	} else if z.XPos > dimensions-1 {
		z.XPos = 0
	}

	z.YPos = z.YPos + y
	if z.YPos < 0 {
		z.YPos = dimensions - 1
	} else if z.YPos > dimensions-1 {
		z.YPos = 0
	}

	return nil
}

func (z *zombie) makeAllTheMoves(grid *Grid) {
	for _, move := range grid.Moves {
		x, y, err := convertMoveToCoordinates(string(move))
		if err != nil {
			log.Fatal(err)
		}

		if err := z.nextPosition(x, y, grid.Dimensions); err != nil {
			log.Fatal(err)
		}

		z.infect(grid)
	}
}

func (z *zombie) infect(grid *Grid) {
	for i, creature := range grid.Creatures {
		infected := z.isInfected(creature)
		if infected {
			grid.Creatures[i].infected = true

			newZombie := zombie{
				XPos: creature.XPos,
				YPos: creature.YPos,
			}

			grid.Zombies = append(grid.Zombies, newZombie)
			grid.Scores++
		}
	}
}

func (z *zombie) isInfected(c creature) bool {
	if c.XPos == z.XPos && c.YPos == z.YPos && c.infected == false {
		return true
	}

	return false
}
