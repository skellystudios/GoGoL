// Use arrays to simulate the game of life

package main

import "GoGOL/src/golutils"

func main() {
	grid := golutils.InitalizeGrid()
	gridSize := golutils.GridSize
	// Every 'tick'
	// Any live cell with fewer than two live neighbours dies, as if caused by under-population.
	// Any live cell with two or three live neighbours lives on to the next generation.
	// Any live cell with more than three live neighbours dies, as if by over-population.
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	for generation := 0; generation < golutils.Generations; generation++ {
		golutils.PrintOutput(generation, grid)
		newGrid := golutils.MakeGrid()
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
	neighbours := golutils.GetNeighbours(coords)
	var sum int
	for _, n := range neighbours {
		sum += grid[n[0]][n[1]]
	}
	return sum
}
