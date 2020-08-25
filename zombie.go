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

func (g *Grid) moveZombie(position int) {
	for _, move := range g.Moves {
		x, y, err := convertMoveToCoordinates(string(move))
		if err != nil {
			log.Fatal(err)
		}

		if err := g.Zombies[position].nextPosition(x, y, g.Dimensions); err != nil {
			log.Fatal(err)
		}

		g.infectCreatures(position)
	}
}

func (g *Grid) infectCreatures(position int) {
	for i, creature := range g.Creatures {
		infected := g.Zombies[position].isInfected(creature)
		if infected {
			g.Creatures[i].infected = true

			newZombie := zombie{
				XPos: creature.XPos,
				YPos: creature.YPos,
			}

			g.Zombies = append(g.Zombies, newZombie)
			g.Scores++
		}
	}
}

func (z *zombie) isInfected(c creature) bool {
	if c.XPos == z.XPos && c.YPos == z.YPos && c.infected == false {
		return true
	}

	return false
}
