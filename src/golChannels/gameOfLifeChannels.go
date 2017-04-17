// Use arrays to simulate the game of life

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// gridSize: N x N size of the grid
const gridSize = 30

// generations: How long to run the simulation for
const generations = 100

var done chan int

func main() {

	grid := initalizeGrid()
	inboxes := makeGridChannels()
	readyChannels := makeGridChannels()

	// Set up the go routines
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			neighbourChannels := getNeighboursChannels([]int{i, j}, inboxes)
			go worker(&grid[i][j], inboxes[i][j], readyChannels[i][j], neighbourChannels)
		}
	}

	numCells := gridSize * gridSize
	for generation := 0; generation < generations; generation++ {
		// Wait until we've got enough "done"s
		for i := 0; i < numCells; i++ {
			<-done
		}
		// Print the grid
		printOutput(generation, grid)
		// Tell all the workers to continue again
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				ready := readyChannels[i][j]
				ready <- 1
			}
		}
	}
}

func worker(cell *int, inbox chan int, ready chan int, neighbours []chan int) {
	for true {
		// Announce my value to my neighbours
		for _, neighbour := range neighbours {
			neighbour <- *cell
		}
		// Tell the main thread that I've done that
		done <- 1
		// Wait until everyone's filled up my inbox
		<-ready
		// Find out what's near me
		neighboursTotal := 0
		for i := 0; i < 8; i++ {
			val := <-inbox
			neighboursTotal += val
		}
		alive := *cell == 1
		// Update my value correctly
		if alive && neighboursTotal < 2 { // Under-population
			*cell = 0
		} else if alive && neighboursTotal > 3 { // Over-population
			*cell = 0
		} else if alive { // Lives on
			*cell = 1
		} else if !alive && neighboursTotal == 3 { // Reproduction
			*cell = 1
		}
	}
}

func getNeighboursChannels(coords []int, grid [][]chan int) []chan int {
	neighbours := getNeighbours(coords)
	neighbourChannels := make([]chan int, 0)
	for _, n := range neighbours {
		neighbourChannels = append(neighbourChannels, grid[n[0]][n[1]])
	}
	return neighbourChannels
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
			grid[i][j] = 0
		}
	}

	// Now we put some simple initial things in (a glider, and something which collapses to a stready diamond)
	initialActive := [][]int{{4, 4}, {4, 5}, {5, 6}, {5, 5}, {9, 5}, {8, 6}, {10, 5}, {10, 6}, {10, 7}}
	for _, element := range initialActive {
		grid[element[0]][element[1]] = 1
	}

	return grid
}

func makeGridChannels() [][]chan int {
	grid := make([][]chan int, gridSize)
	for i := range grid {
		grid[i] = make([]chan int, gridSize)
	}
	// initialize the grid of zeros
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			grid[i][j] = make(chan int, 8)
		}
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
