package solve

import (
	"testing"
	"fmt"

	"github.com/FeCastle/blocks/cube"
)

func TestSolveOK2 (t *testing.T) {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube1 := cube.SetCubeSize(2, 2)
	snake := []int{1, 1, 1, 1}
	if result := Solve(cube1, snake); result != true {
		t.Errorf("Solve of 2x2 square should be OK!\n")
	}
}

func TestSolveFAIL2 (t *testing.T) {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube1 := cube.SetCubeSize(2, 2)
	snake := []int{1, 2, 1}
	if result := Solve(cube1, snake); result != false {
		t.Errorf("Solve of 2x2 square should be FAIL!\n")
	}
}

func TestSolveFail (t *testing.T) {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube2 := cube.SetCubeSize(3,2)
	snake2 := []int{1,2,2,2,2}
	if result := Solve(cube2, snake2);  result != false {
		t.Errorf("Solve of 3x3 square should FAIL!\n")
	}
}

func TestSolveOK (t *testing.T) {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube3 := cube.SetCubeSize(3, 2)
	snake3 := []int{1, 2, 2, 2, 1, 1}
	if result := Solve(cube3, snake3); result != true {
		t.Errorf("Solve of 3x3 square should be OK!\n")
	}
}

func TestCubeOK2 (t *testing.T) {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube2 := cube.SetCubeSize(2, 3)
	snake := []int{1, 1, 1, 1, 1, 1, 1, 1}
	if result := Solve(cube2, snake); result != true {
		t.Errorf("Solve of Cube 2x2x2 should be OK!\n")
	}
}

func TestCubeFAIL2 (t *testing.T) {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube2 := cube.SetCubeSize(2, 2)
	snake := []int{1, 1, 1, 1, 2, 2}
	if result := Solve(cube2, snake); result != false {
		t.Errorf("Solve of Cube 3x3x3 should be FAIL!\n")
	}
}

func TestCubeOK3 (t *testing.T) {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube3 := cube.SetCubeSize(3, 3)
	snake := []int{1, 2, 2, 1, 1, 2, 1, 2, 1, 2, 1, 1, 1, 2, 2, 1, 2, 2}
	if result := Solve(cube3, snake); result != true {
		t.Errorf("Solve of Cube 3x3x3 should be OK!\n")
	}
}
