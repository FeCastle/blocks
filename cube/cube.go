package cube 

 import (
	"fmt"
)

type Direction int
const (
	LEFT Direction = 1 + iota
	RIGHT
	NORTH
	SOUTH
	UP
	DOWN
	INITIAL
)

type Cube struct {
	size int
	dim int  // 2D or 3D
	LastMove int // the last spot
	LastMoveDir Direction // the last spot - Used to generate the next set of moves.
	                      // Will be broken after undo, but that is OK, because list is generated.
	Blocks []int
	BC int // size * size *size (for 3D)
}

func main() {
	fmt.Printf("Starting Cube.\n")
}

func SetCubeSize (s int, d int) *Cube {
	c := &Cube {
		size: s,
		dim: d,
		LastMove: -1, // No moves yet.
	}

	c.BC = c.getBlockCount()
	c.Blocks = make([]int, c.BC)
	for i := 0; i < c.BC; i++ {
		c.Blocks[i] = -1
	}
	return c
}

func (c *Cube) getBlockCount () int {
	if c.dim == 1 {
		return c.size
	} else if c.dim == 2 {
		return c.size * c.size
	} else if c.dim == 3 {
		return c.size * c.size *c.size
	} else {
		return 0
	}
}

func (c *Cube) setBlock (depth int, pos int) {
	// fmt.Printf("Setting block at pos %d to %d\n", pos, depth)
	c.Blocks[pos] = depth
}

func (c *Cube) MakeMove (depth int, count int, pos int, d Direction) {
	// fmt.Printf("***** Depth %d LastMove %d - Moving from %d - %d %s\n", depth, c.LastMove, pos, count, dirToA(d))
	// On undo set LastMove back to the start
	if depth < 0 {
		c.LastMove = pos
	}
	if c.LastMove < 0 {
		if depth != 1 {
			fmt.Printf("ERROR:  Initial Move should only be at depth 0\n")
		}
		if d != INITIAL {
			fmt.Printf("ERROR:  Initial Move should always be UP\n")
		}
		if count != 1 {
			fmt.Printf("ERROR:  Initial Move should always be length 1 (But it was %d).\n", count)
		}
		c.setBlock(depth, pos)
	} else {
		for i := 0; i < count; i++ {
			switch d {
			case LEFT:
				pos = pos - 1
			case RIGHT:
				pos = pos + 1
			case NORTH:
				pos = pos - c.size
			case SOUTH:
				pos = pos + c.size
			case UP:
				pos = pos + (c.size * c.size)
			case DOWN:
				pos = pos - (c.size * c.size)
			case INITIAL:
				pos = pos
			default:
				fmt.Printf("ERROR: Unkown Drection in makeMove.\n")
			}
			c.setBlock(depth, pos)
		}
	}

	// If depth < 0 then this is a reverse, this will not be correct after UNDO,
	// But that is oK, because this is only used to generate the next move list
	if depth >= 0 {
		c.LastMoveDir = d
		c.LastMove = pos
	}

}

func (c *Cube) empty (cPos int) bool{
	// Chekc out of bounds first
	if cPos < 0 || cPos >= c.BC {
		return false
	}

	if c.Blocks[cPos] == -1 {
		return true
	} else {
		return false
	}
}

// Try to lay some number of Blocks to the left
func (c *Cube) TryLeft (cPos int, count int) bool {
	return c.tryLeftRight(cPos, count, false)
}
// Try to lay some number of Blocks to the right
func (c *Cube) TryRight (cPos int, count int) bool {
	return c.tryLeftRight(cPos, count, true)
}

// Try to put some Blocks in the cube at cPos (Left)
func (c *Cube) tryLeftRight (cPos int, count int, right bool) bool {
	nPos := cPos
	row := nPos/c.size

	for try := 0; try < count; try++ {
		if right == true {
			nPos++
		} else {
			nPos--
		}
		// fmt.Printf("Trying Left at pos %d row %d) \n", nPos, row)
		if c.empty(nPos) == false {
			return false
		}
		if nPos/c.size != row {
			return false
		}
	}

	return true
}

func (c *Cube) TryNorth (cPos int, count int) bool {
	return c.tryNorthSouth(cPos, count, true)
}

func (c *Cube) TrySouth (cPos int, count int) bool {
	return c.tryNorthSouth(cPos, count, false)
}

// Try to put some Blocks in the cube N/S
func (c *Cube) tryNorthSouth (cPos int, count int, north bool) bool {
	nPos := cPos
	squareSize := c.size * c.size
	level := nPos / squareSize

	// fmt.Printf("---- Trying N/S from Pos %d Count %d (Level %d) Total BS=%d\n", cPos, count, level, c.bc)
	for try := 0; try < count; try++ {
		if north == true {
			nPos = nPos - c.size
		} else {
			nPos = nPos + c.size
		}

		// fmt.Printf("Trying N/S from Pos %d Count %d\n", nPos, count)
		if c.empty(nPos) == false {
			// fmt.Printf("Failed - Empty!\n")
			return false
		}
		if nPos / squareSize != level {
			// nextLevel := nPos / squareSize
			// fmt.Printf("Failed Level %d\n", nextLevel)
			return false
		}
	}

	// TODO MARK AS USED
	return true
}


func (c *Cube) TryUp (cPos int, count int) bool {
	return c.tryUpDown(cPos, count, true)
}
func (c *Cube) TryDown (cPos int, count int) bool {
	return c.tryUpDown(cPos, count, false)
}
// Try to put some Blocks in the cube N/S
func (c *Cube) tryUpDown (cPos int, count int, up bool) bool {
	nPos := cPos
	squareSize := c.size * c.size

	// fmt.Printf("---- Trying U/D from Pos %d Count %d (Level %d) Total BS=%d\n", cPos, count, level, c.bc)
	for try := 0; try < count; try++ {
		if up == true {
			nPos = nPos + squareSize
		} else {
			nPos = nPos - squareSize
		}

		// fmt.Printf("Trying U/D from Pos %d Count %d\n", nPos, count)
		if c.empty(nPos) == false {
			// fmt.Printf("Failed - Empty!\n")
			return false
		}
	}

	// TODO MARK AS USED
	return true
}
