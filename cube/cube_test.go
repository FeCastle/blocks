package cube

import (
  "testing"
   "fmt"
)

func TestConfig (t *testing.T) {
        fmt.Printf("Testing 2 by 2")
	// line
	sizeTest(3,1,3,t)

	// square
	sizeTest(2,2,4,t)
	sizeTest(3,2,9,t)
	sizeTest(4,2,16,t)

	// cube
	sizeTest(4,3,64,t)

	// error
	sizeTest(3,4,0, t)

}

func TestEmpty (t *testing.T) {
	cube := SetCubeSize(3,3)
	empty := cube.empty(0)
	if empty != true {
		t.Errorf("A new cube should be empty at pos 0")
	}
}

func TestTryLeftLine (t *testing.T) {
	cube := SetCubeSize(3,1)

	if cube.TryLeft(0,1) != false {
		t.Errorf("3x1 going left from 0 should fail")
	}

	if cube.TryLeft(2,2) != true {
		t.Errorf("3x1 going left from Pos2 should be OK.")
	}

	if cube.TryLeft(2,3) != false {
		t.Errorf("3x1 going left from Pos2 should fail because there are too many blocks.")
	}
}

func TestTryLRCube (t *testing.T) {
	cube := SetCubeSize(3, 3)
	if cube.TryRight(21, 2) != true {
		t.Errorf("3x3x3 going Right from Pos 21 should be OK.")
	}

	if cube.TryRight(22, 2) != false {
		t.Errorf("3x3x3 going Right from Pos 22 should fail.")
	}

	if cube.TryRight(24, 3) != false {
		t.Errorf("3x3x3 going Right from Pos 24 should fail (TOO LONG).")
	}

	if cube.TryRight(13, 2) != false {
		t.Errorf("3x3x3 going Left from Pos 13 should fail (TOO LONG).")
	}
	if cube.TryRight(27, 1) != false {
		t.Errorf("3x3x3 going starting from POS 27 should fail (OOB).")
	}
}

func TestTryRightSquare (t *testing.T) {
	cube := SetCubeSize(3, 2)

	if cube.TryRight(2, 2) != false {
		t.Errorf("3x3 going right from Pos2 fail.")
	}
	if cube.TryRight(4, 2) != false {
		t.Errorf("3x3 going right should fail from center (new row)")
	}
	if cube.TryRight(6, 2) != true {
		t.Errorf("3x3 going right should be OK from pos 6")
	}
	if cube.TryRight(8, 3) != false {
		t.Errorf("3x3 going left should fail from pos 8 (too long) ")
	}
}

func TestTryLeftSquare (t *testing.T) {
	cube := SetCubeSize(3, 2)

	if cube.TryLeft(2, 2) != true {
		t.Errorf("3x3 going left from Pos2 should be OK.")
	}
	if cube.TryLeft(4, 2) != false {
		t.Errorf("3x3 going left should fail from center (new row)")
	}
	if cube.TryLeft(8, 2) != true {
		t.Errorf("3x3 going left should be OK from pos 8")
	}
	if cube.TryLeft(8, 3) != false {
		t.Errorf("3x3 going left should fail from pos 8 (too long) ")
	}
}

func TestTryUp (t *testing.T) {
	cube := SetCubeSize(3, 3)
	if cube.TryUp(0, 3) != false {
		t.Errorf("3x3 going Up 3 from 0 should fail")
	}
	if cube.TryUp(16, 1) != true {
		t.Errorf("3x3 going Up from 16 should be OK")
	}
	if cube.TryUp(25, 2) != false {
		t.Errorf("3x3 going Up from 25 should fail (len 2)")
	}
}

func TestTryDown (t *testing.T) {
	cube := SetCubeSize(3, 3)
	if cube.TryDown(20, 2) != true {
		t.Errorf("3x3 going down 2 from Pos 20 should be OK")
	}
	if cube.TryDown(15, 1) != true {
		t.Errorf("3x3 going down 1 from 15 should Fail")
	}
	if cube.TryDown(7, 1) != false {
		t.Errorf("3x3 going Down 1 from 7 should Fail")

	}
}

func TestTryNorth (t *testing.T) {
	cube := SetCubeSize(3, 3)
	if cube.TryNorth(1, 1) != false {
		t.Errorf("3x3 going North should fail from 1")
	}
	if cube.TryNorth(16, 2) != true {
		t.Errorf("3x3 going North from 16 should be OK")
	}
	if cube.TryNorth(26, 3) != false {
		t.Errorf("3x3 going North from 26 should fail (len 3)")
	}
}

func TestTrySouth (t *testing.T) {
	cube := SetCubeSize(3, 3)
	if cube.TrySouth(0, 2) != true {
		t.Errorf("3x3 going South should be OK from 0")
	}
	if cube.TrySouth(17, 1) != false {
		t.Errorf("3x3 going South from 17 should fail")
	}
	if cube.TrySouth(22, 2) != false {
		t.Errorf("3x3 going South from 22 should fail")
	}

}


func sizeTest (size int, dim int, expected int, t *testing.T) {
	cube := SetCubeSize(size, dim)
	bc := cube.getBlockCount()
	if bc != expected {
		t.Errorf("Expected %d x %d to have %d blocks, instead it had %d\n", size, size, expected, bc)
	}
}


