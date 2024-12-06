package day6

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Coordinates struct {
	x  int
	y  int
	dx int
	dy int
}

func (c *Coordinates) turn() {
	newDX := c.dx
	newDY := c.dy
	c.dx = -newDY
	c.dy = newDX
}

func (c *Coordinates) simulatePath(grid []string, obstructionX, obstructionY int) (map[string]int, [][]int, bool) {
	visited := make(map[string]int)
	visitedCoords := [][]int{}

	inBounds := true
	for inBounds {
		nextX := c.x + c.dx
		nextY := c.y + c.dy

		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			inBounds = false
			break
		}

		if nextX == obstructionX && nextY == obstructionY {
			c.turn()
			continue
		}

		if grid[nextY][nextX] == '#' {
			c.turn()
			continue
		}

		if visited[fmt.Sprintf("(%d, %d)", c.x, c.y)] > 10 {
			return visited, visitedCoords, true
		}

		c.x, c.y = nextX, nextY

		visited[fmt.Sprintf("(%d, %d)", c.x, c.y)]++
		visitedCoords = append(visitedCoords, []int{c.x, c.y})
	}
	return visited, visitedCoords, false
}

func readData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n %v", err)
	}
	defer file.Close()

	var grid []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func partOne(path string) int {
	grid := readData(path)
	c := Coordinates{
		x:  0,
		y:  0,
		dx: 0,
		dy: -1,
	}

	visited := make(map[string]int)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '^' {
				c.x, c.y = x, y
				visited, _, _ = c.simulatePath(grid, -1, -1)
			}
		}
	}
	return len(visited)
}

func partTwo(path string) int {
	grid := readData(path)
	c := Coordinates{
		x:  0,
		y:  0,
		dx: 0,
		dy: -1,
	}

	var origX, origY int

	visitedCoords := [][]int{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '^' {
				c.x, c.y = x, y
				origX, origY = x, y
				_, visitedCoords, _ = c.simulatePath(grid, -1, -1)
			}
		}
	}

	uniqueObstacles := make(map[string]bool)
	for i := 0; i < len(visitedCoords); i++ {
		ox, oy := visitedCoords[i][0], visitedCoords[i][1]
		if grid[oy][ox] == '.' {
			tempC := Coordinates{
				x:  origX,
				y:  origY,
				dx: 0,
				dy: -1,
			}
			_, _, loopFound := tempC.simulatePath(grid, ox, oy)
			if loopFound {
				obstacleKey := fmt.Sprintf("%d,%d", ox, oy)
				uniqueObstacles[obstacleKey] = true
			}
		}
	}
	return len(uniqueObstacles)
}

func DaySix() {
	start := time.Now()
	partOne := partOne("day6/day6.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 6 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day6/day6.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 6 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
