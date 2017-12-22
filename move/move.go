package move

import (
	"github.com/FeCastle/blocks/cube"
	"fmt"
)

type MoveList struct {
	moves []int             // Starting position
	directions []cube.Direction  // Direction
	current int             // The current move to try
	count int               // The number of blocks in the move
}

func (m *MoveList) addMove(pos int, d cube.Direction) {
	m.moves = append (m.moves, pos)
	m.directions = append(m.directions, d)
}

func (m *MoveList) UndoMove(c *cube.Cube) {
	// Set the depth to -1 -> empty blocks are -1
	prev := m.current - 1
	// fmt.Printf("Undo Previous move at %d\n", prev)
	c.MakeMove(-1, m.count, m.moves[prev], m.directions[prev])
}

func (m *MoveList) MakeMove(c *cube.Cube, depth int) bool {
	if m.current >= len(m.moves) {
		return false
	}
	c.MakeMove(depth, m.count, m.moves[m.current], m.directions[m.current])
	// m.print()
	m.current++
	return true
}

func GenMoveList (c *cube.Cube, count int) *MoveList {
	m := &MoveList{
		current: 0,
		count: count,
	}
	// Initial move list is large
	// Only need to do a single direction for all of the blocks
	// Actually only need to do 1/2 because we could invers
	// The first move is always 1 block up
	if c.LastMove < 0 {
		totalMoves := c.BC/2 + 1
		m.moves = make([]int, totalMoves)
		m.directions = make([]cube.Direction, totalMoves)
		for i:=0; i < totalMoves; i++ {
			m.moves[i] = i
			m.directions[i] = cube.INITIAL
		}
	} else {
		switch c.LastMoveDir {
		// Try North South Up Down
		case cube.RIGHT:
			fallthrough
		case cube.LEFT:
			if c.TryNorth(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.NORTH)
			}
			if c.TrySouth(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.SOUTH)
			}
			if c.TryUp(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.UP)
			}
			if c.TryDown(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.DOWN)
			}
		// Try Left Right Up Down
		case cube.NORTH:
			fallthrough
		case cube.SOUTH:
			if c.TryLeft(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.LEFT)
			}
			if c.TryRight(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.RIGHT)
			}
			if c.TryUp(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.UP)
			}
			if c.TryDown(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.DOWN)
			}
		case cube.INITIAL:
			fallthrough
		case cube.UP:
			fallthrough
		case cube.DOWN:
			if c.TryLeft(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.LEFT)
			}
			if c.TryRight(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.RIGHT)
			}
			if c.TryNorth(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.NORTH)
			}
			if c.TrySouth(c.LastMove, count) == true {
				m.addMove(c.LastMove, cube.SOUTH)
			}
		default:
			fmt.Errorf("ERROR: Can't Gen moves. LastMove NOT set!!!!")
		}
	}

	return m
}

func dirToA (d cube.Direction) string {
	switch d {
	case cube.LEFT:
		return "L"
	case cube.RIGHT:
		return "R"
	case cube.NORTH:
		return "N"
	case cube.SOUTH:
		return "S"
	case cube.UP:
		return "U"
	case cube.DOWN:
		return "D"
	case cube.INITIAL:
		return "I"
	default:
		return "ERROR! (Dir To A)!!"
	}
}

func (m *MoveList) PrintCurrent(i int) {
	fmt.Printf("Move %d: %d %s\n", i, m.count, dirToA(m.directions[m.current - 1]))
}

func (m *MoveList) print () {
	fmt.Printf ("Current Move: %d \n", m.current)
	for i:=0; i < len(m.moves); i++ {
		fmt.Printf(" - From %d: %s%d", m.moves[i], dirToA(m.directions[i]), m.count)
		if i == m.current {
			fmt.Printf(" *")
		}
		fmt.Printf("\n")
	}
}
