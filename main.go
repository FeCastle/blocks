package main

import (
	"fmt"

	"github.com/FeCastle/blocks/cube"
	"github.com/FeCastle/blocks/solve"
)

// 
func main() {
	fmt.Printf("--------------------------------------------------------------------------\n")
	cube4 := cube.SetCubeSize(4, 3)
	snake := []int{1,2,1,2,1,1,3,1,2,1,2,1,2,1,1,1,1,1,1,1,1,2,2,1,1,1,1,1,2,3,1,1,1,3,1,2,1,1,1,1,1,1,1,1,1,3 }
	if result := solve.Solve(cube4,snake); result != true {
		fmt.Errorf("Solve of Cube 4 should be OK!\n")
	}
}
