package day15

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

type Obstacle struct {
	x int
	y int
}

func (r *Robot) setDir(dir rune) {
	switch dir {
	case '^':
		r.vx = 0
		r.vy = -1
	case 'v':
		r.vx = 0
		r.vy = 1
	case '>':
		r.vx = 1
		r.vy = 0
	case '<':
		r.vx = -1
		r.vy = 0
	}
}

func (r *Robot) move(grid []string) {
	newX, newY := r.x+r.vx, r.y+r.vy
	if grid[newY][newX] != '#' {
		r.x, r.y = newX, newY
	}
}

func (o *Obstacle) push(r Robot, obs []*Obstacle, w, h int, grid []string) {
	newX, newY := o.x+r.vx, o.y+r.vy

	if newX >= 1 && newX < w-1 && newY >= 1 && newY < h-1 && grid[newY][newX] != '#' {
		for _, other := range obs {
			if other.x == newX && other.y == newY {
				other.push(r, obs, w, h, grid)
				return
			}
		}
		o.x, o.y = newX, newY
	}
}

func readData(path string) ([]string, []string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	grid := []string{}
	moves := []string{}

	currList := &grid

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currList = &moves
			continue
		}
		*currList = append(*currList, line)
	}
	return grid, moves
}

func calcGPS(obs []*Obstacle) int {
	gpsScore := 0
	for _, o := range obs {
		gpsScore += o.y*100 + o.x
	}
	return gpsScore
}

func partOne(path string) int {
	grid, moves := readData(path)

	w, h := len(grid[0]), len(grid)

	r := Robot{}
	obs := []*Obstacle{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '@' {
				r.x, r.y = x, y
			}
			if grid[y][x] == 'O' {
				obs = append(obs, &Obstacle{x: x, y: y})
			}
		}
	}

	for i := 0; i < len(moves); i++ {
		for j := 0; j < len(moves[i]); j++ {
			r.setDir(rune(moves[i][j]))

			nextX, nextY := r.x+r.vx, r.y+r.vy
			pushed := false

			for _, o := range obs {
				if o.x == nextX && o.y == nextY {
					o.push(r, obs, w, h, grid)
					pushed = true
					break
				}
			}

			if !pushed && grid[nextY][nextX] != '#' {
				r.move(grid)
			}
		}
	}
	return calcGPS(obs)
}

func DayFifteen() {
	start := time.Now()
	partOne := partOne("day15/day15_test.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 15 part one: %d\nexecution time: %v\n", partOne, duration)

	// start = time.Now()
	// partTwo := partTwo("day15/day15.txt")
	// duration = time.Since(start)
	// fmt.Printf("Solution for day 15 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
