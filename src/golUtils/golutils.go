// Use arrays to simulate the game of life

package golutils

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// GridSize: N x N size of the grid
const GridSize = 30

// Generations: How long to run the simulation for
const Generations = 500

func GetNeighbours(coords []int) [][]int {
	directions := [][]int{{-1, -1}, {-1, 0}, {0, -1}, {0, 1}, {1, 0}, {1, 1}, {1, -1}, {-1, 1}}
	results := make([][]int, 0)
	for i := range directions {
		x := directions[i][0] + coords[0]
		y := directions[i][1] + coords[1]

		// Make the grid wrap
		if x < 0 {
			x = x + GridSize
		}
		if y < 0 {
			y = y + GridSize
		}
		if x >= GridSize {
			x = 0
		}
		if y >= GridSize {
			y = 0
		}

		results = append(results, []int{x, y})

	}
	return results
}

func MakeGrid() [][]int {
	grid := make([][]int, GridSize)
	for i := range grid {
		grid[i] = make([]int, GridSize)
	}
	return grid
}

func InitalizeGrid() [][]int {
	grid := MakeGrid()
	// initialize the grid of zeros
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			grid[i][j] = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)
		}
	}

	// // Now we put some simple initial things in (a glider, and something which collapses to a stready diamond)
	// initialActive := [][]int{{34, 34}, {34, 35}, {35, 36}, {35, 35}, {9, 5}, {8, 6}, {10, 5}, {10, 6}, {10, 7}}
	// for _, element := range initialActive {
	// 	grid[element[0]][element[1]] = 1
	// }

	return grid
}

func PrintGrid(grid [][]int) {
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			icon := "_"
			if grid[i][j] == 1 {
				icon = "█"
			}
			fmt.Print(icon, " ")
		}
		fmt.Println()
	}
}

func PrintOutput(generation int, grid [][]int) {
	ClearTerminal()
	fmt.Print("Generation: ", generation)
	fmt.Println()
	PrintGrid(grid)
}

func ClearTerminal() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MakeGridChannels() [][]chan int {
	grid := make([][]chan int, GridSize)
	for i := range grid {
		grid[i] = make([]chan int, GridSize)
	}
	// initialize the grid of zeros
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			grid[i][j] = make(chan int, 8)
		}
	}
	return grid
}
