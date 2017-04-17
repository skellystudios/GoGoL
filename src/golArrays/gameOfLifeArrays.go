// Use arrays to simulate the game of life

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// gridSize: N x N size of the grid
const gridSize = 45

// generations: How long to run the simulation for
const generations = 1000

func main() {
	grid := initalizeGrid()

	// Every 'tick'
	// Any live cell with fewer than two live neighbours dies, as if caused by under-population.
	// Any live cell with two or three live neighbours lives on to the next generation.
	// Any live cell with more than three live neighbours dies, as if by over-population.
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	for generation := 0; generation < generations; generation++ {
		printOutput(generation, grid)
		newGrid := makeGrid()
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				alive := grid[i][j] == 1
				neighboursTotal := sumNeighbours([]int{i, j}, grid)
				if alive && neighboursTotal < 2 { // Under-population
					newGrid[i][j] = 0
				} else if alive && neighboursTotal > 3 { // Over-population
					newGrid[i][j] = 0
				} else if alive { // Lives on
					newGrid[i][j] = 1
				} else if !alive && neighboursTotal == 3 { // Reproduction
					newGrid[i][j] = 1
				}
			}
		}
		grid = newGrid

	}

}

func sumNeighbours(coords []int, grid [][]int) int {
	neighbours := getNeighbours(coords)
	var sum int
	for _, n := range neighbours {
		sum += grid[n[0]][n[1]]
	}
	return sum
}

func getNeighbours(coords []int) [][]int {
	directions := [][]int{{-1, -1}, {-1, 0}, {0, -1}, {0, 1}, {1, 0}, {1, 1}, {1, -1}, {-1, 1}}
	results := make([][]int, 0)
	for i := range directions {
		x := directions[i][0] + coords[0]
		y := directions[i][1] + coords[1]

		// Make the grid wrap
		if x < 0 {
			x = x + gridSize
		}
		if y < 0 {
			y = y + gridSize
		}
		if x >= gridSize {
			x = 0
		}
		if y >= gridSize {
			y = 0
		}

		results = append(results, []int{x, y})

	}
	return results
}

func makeGrid() [][]int {
	grid := make([][]int, gridSize)
	for i := range grid {
		grid[i] = make([]int, gridSize)
	}
	return grid
}

func initalizeGrid() [][]int {
	grid := makeGrid()
	// initialize the grid of zeros
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			grid[i][j] = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)
		}
	}

	// Now we put some simple initial things in (a glider, and something which collapses to a stready diamond)
	initialActive := [][]int{{34, 34}, {34, 35}, {35, 36}, {35, 35}, {9, 5}, {8, 6}, {10, 5}, {10, 6}, {10, 7}}
	for _, element := range initialActive {
		grid[element[0]][element[1]] = 1
	}

	return grid
}

func printGrid(grid [][]int) {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			icon := "_"
			if grid[i][j] == 1 {
				icon = "█"
			}
			fmt.Print(icon, " ")
		}
		fmt.Println()
	}
}

func printOutput(generation int, grid [][]int) {
	clearTerminal()
	fmt.Print("Generation: ", generation)
	fmt.Println()
	printGrid(grid)
}

func clearTerminal() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
