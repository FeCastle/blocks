package solve

import (
	"fmt"

	"github.com/FeCastle/blocks/cube"
	"github.com/FeCastle/blocks/move"
)

func Solve ( c *cube.Cube, snake []int) bool {
	// cSize := c.getBlockCount()
	// TODO Check that snake is correct size
	var moveStack []*move.MoveList

	maxDepth := len(snake)

	best := 4
	// Always start with a block of 1 UP
	ml := move.GenMoveList(c, snake[0])

	// Push the initial list onto the stack and pop it off, continue to pop until it is empty

	for moveStack = append(moveStack, ml); len(moveStack) > 0; {
		depth := len(moveStack)


		// The current move was made
		res := ml.MakeMove(c, depth)
		if res == true {
			if depth >= maxDepth {
				fmt.Printf("Hit Max depth of %d is this solved?\n", depth)
				for i := 0; i < len(moveStack); i++ {
					ml = moveStack[i]
					ml.PrintCurrent(i)
				}

				return true
			}

			if depth > best {
				best = depth
				fmt.Printf("------------ New Best at Depth %d? -------\n", best)
				for i := 0; i < len(moveStack); i++ {
					ml = moveStack[i]
					ml.PrintCurrent(i)
				}
			}
			// fmt.Printf("OK - Generating new move list.\n")
			ml = move.GenMoveList(c, snake[depth])
			moveStack = append(moveStack, ml)
		} else {
			// Pop this current and chuck it
			// fmt.Printf("Failed - Popping old move stack size:.\n")
			moveStack = moveStack[:len(moveStack) - 1]
			if len(moveStack) > 0 {
				ml = moveStack[len(moveStack)-1]
				ml.UndoMove(c)
			}
		}
	}

	fmt.Printf("Out of moves (FAIL)!!!\n")
	return false
}

