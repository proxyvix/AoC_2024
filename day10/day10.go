package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func readData(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	grid := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return grid
}

func isValid(x, y, w, h int) bool {
	return x >= 0 && y >= 0 && x < w && y < h
}

func dfs(grid [][]int, x, y, w, h int, visited [][]bool, reachable map[[2]int]bool) {
	if !isValid(x, y, w, h) || visited[y][x] {
		return
	}
	visited[y][x] = true

	if grid[y][x] == 9 {
		reachable[[2]int{x, y}] = true
		return
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if isValid(nx, ny, w, h) && grid[ny][nx] == grid[y][x]+1 {
			dfs(grid, nx, ny, w, h, visited, reachable)
		}
	}
}

func calculateTrailheadScores(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	totalScore := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 0 {
				visited := make([][]bool, h)
				for i := range visited {
					visited[i] = make([]bool, w)
				}

				reachable := make(map[[2]int]bool)
				dfs(grid, x, y, w, h, visited, reachable)
				totalScore += len(reachable) // Add count of unique reachable 9s
			}
		}
	}

	return totalScore
}

func dfsCountTrails(grid [][]int, x, y, w, h int, currentTrail []int) int {
	if !isValid(x, y, w, h) || (len(currentTrail) > 0 && grid[y][x] != currentTrail[len(currentTrail)-1]+1) {
		return 0
	}

	if grid[y][x] == 9 {
		return 1
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	trailCount := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		trailCount += dfsCountTrails(grid, nx, ny, w, h, append(currentTrail, grid[y][x]))
	}
	return trailCount
}

func calculateTrailheadRatings(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	totalRating := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 0 {
				totalRating += dfsCountTrails(grid, x, y, w, h, []int{})
			}
		}
	}

	return totalRating
}

func partOne(path string) int {
	grid := readData(path)
	return calculateTrailheadScores(grid)
}

func partTwo(path string) int {
	grid := readData(path)
	return calculateTrailheadRatings(grid)
}

func DayTen() {
	start := time.Now()
	result := partOne("day10/day10.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 10 part one: %d\nExecution time: %v\n", result, duration)

	start = time.Now()
	result = partTwo("day10/day10.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 10 part two: %d\nExecution time: %v\n", result, duration)
}
