package move

import (
	"testing"
	"fmt"

	"github.com/FeCastle/blocks/cube"
)

func TestMoveUndo (t *testing.T) {
	fmt.Printf("Test Moves!\n")
	c1 := cube.SetCubeSize(3,2)
	//snake := []int{2,2,2,1,1}
	snake := []int{1,2,2,1,1}

	ml1 := GenMoveList(c1, snake[0])
	ml1.print()

	depth := 1

	if c1.Blocks[0] >= 0  {
		t.Errorf("Before initial move, pos 0 should be set to -1!\n")
	}
	ml1.MakeMove(c1, depth)

	if c1.LastMoveDir != cube.INITIAL {
		t.Errorf("After initial move, the LastMove dir should be INITIAL!\n")
	}

	// After the move, the current should be 1.
	if ml1.current != 1 {
		t.Errorf("After initial move, the current should be 1!\n")
	}
	if c1.LastMove != 0 {
		t.Errorf("After initial move, we should be at pos 0!\n")
	}
	if c1.Blocks[0] != depth {
		t.Errorf("After initial move, we should be pos 0 should be set to 0!\n")
	}

	// If we move again R2 (or S2)
	depth = 2
	ml2 := GenMoveList(c1, snake[1])
	ml2.print()

	// Before the move, the current should be 0.
	if ml2.current != 0 {
		t.Errorf("After initial move, the current should be 1!\n")
	}
	ml2.MakeMove(c1, depth)
	// Before the move, the current should be 0.
	if ml2.current != 1 {
		t.Errorf("After initial move, the current should be 1, but it was %d!\n", ml2.current)
	}
	if c1.LastMove != 2 {
		t.Errorf("After second move, we should be at pos 2!, but it was at %d\n", c1.LastMove)
	}
	if c1.Blocks[1] != depth || c1.Blocks[2] != depth {
		t.Errorf("After second move, block 1 was %d should be %d!\n", c1.Blocks[1], depth)
		t.Errorf("After second move, block 2 was %d should be %d!\n", c1.Blocks[2], depth)
	}

	depth = 3
	ml3 := GenMoveList(c1, snake[1])
	ml3.print()
	ml3.MakeMove(c1, depth)
	ml3.UndoMove(c1)

	// ********** UNDO ************** Now Undo the 2nd move
	ml2.UndoMove(c1)
	// We lose this after an undo
	/*
	if c1.LastMoveDir != UP {
		t.Errorf("After undo move, the LastMove dir should be UP!\n")
	}
	*/

	// After the move, the current should be 1.
	if ml2.current != 1 {
		t.Errorf("After Undo - We should still be at move 1!\n")
	}
	if c1.LastMove != 0 {
		t.Errorf("After Undo - We should be back at pos 0, but we were at pos %d !\n", c1.LastMove)
	}
	if c1.Blocks[0] != 1 {
		t.Errorf("After Undo - the block at pos 0 should still be 1!\n")
	}
	if c1.Blocks[1] >= 0 || c1.Blocks[2] >= 0 {
		t.Errorf("After second move, block 1 was %d should be %d!\n", c1.Blocks[1], -1)
		t.Errorf("After second move, block 2 was %d should be %d!\n", c1.Blocks[2], -1)
	}

	// Undo the first move (back to start)
	ml1.UndoMove(c1)
	if c1.Blocks[0] >= 0  {
		t.Errorf("Before initial move, pos 0 should be set to -1!\n")
	}
	if ml1.current != 1 {
		t.Errorf("Atfer 2nd undo we should be on try 1, but we were on %d\n", ml1.current)
	}

	// Remake first
	depth = 1
	ml1.MakeMove(c1, depth)
		// After the move, the current should be 1.
	if ml1.current != 2 {
		t.Errorf("After 2nd initial move, the current should be 2!\n")
	}
	if c1.LastMove != 1 {
		t.Errorf("After 2nd initial move, we should be at pos 1!\n")
	}
	if c1.Blocks[1] != depth {
		t.Errorf("After 2nd initial move, block at pos 1 should be set to 1!\n")
	}
	if c1.Blocks[0] !=  -1 {
		t.Errorf("After 2nd initial move, we should be pos 0 should be set to -1!\n")
	}

	ml2 = GenMoveList(c1, snake[1])
	ml2.print()
}
