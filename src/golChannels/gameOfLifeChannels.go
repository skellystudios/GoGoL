// Use arrays to simulate the game of life

package main

import (
	"GoGOL/src/golutils"
)

const gridSize = golutils.GridSize
const numCells = gridSize * gridSize

var done = make(chan int, numCells)

func main() {

	grid := golutils.InitalizeGrid()
	inboxes := golutils.MakeGridChannels()
	readyChannels := golutils.MakeGridChannels()

	// Set up the go routines
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			neighbourChannels := getNeighboursChannels([]int{i, j}, inboxes)
			go worker(&grid[i][j], inboxes[i][j], readyChannels[i][j], neighbourChannels)
		}
	}

	for generation := 0; generation < golutils.Generations; generation++ {
		// Wait until we've got enough "done"s
		for i := 0; i < numCells; i++ {
			<-done
		}
		// Print the grid
		golutils.PrintOutput(generation, grid)
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
	neighbours := golutils.GetNeighbours(coords)
	neighbourChannels := make([]chan int, 0)
	for _, n := range neighbours {
		neighbourChannels = append(neighbourChannels, grid[n[0]][n[1]])
	}
	return neighbourChannels
}
