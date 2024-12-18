package day16

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
)

type Deer struct {
	x int
	y int
}

func readData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	maze := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	return maze
}

func DSA(maze []string, d Deer, w, h int) {
	distances := make([][]int, h)
	for i := range distances {
		distances[i] = make([]int, w)
		for j := range distances[i] {
			distances[i][j] = math.MaxInt
		}
	}
	distances[d.y][d.x] = 0

	q := list.New()
	q.PushBack([3]int{d.x, d.y, 0})

	directions := [][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}

	for q.Len() > 0 {
		e := q.Front()
		curr := e.Value.([3]int)
		q.Remove(e)

		xx, yy, dist := curr[0], curr[1], curr[2]

		for _, dir := range directions {
			nextX, nextY := xx+dir[0], yy+dir[1]

			if maze[nextY][nextX] == 'E' {
				fmt.Println("Found:", xx, yy, dist)
			}

			if maze[nextY][nextX] == '#' {
				continue
			}

			newDist := dist + 1
			if newDist < distances[nextY][nextX] {
				distances[nextY][nextX] = newDist
				q.PushBack([3]int{nextX, nextY, newDist})
			}
			fmt.Println(xx, yy, dist)
		}
	}
}

func partOne(path string) {
	maze := readData(path)

	w, h := len(maze[0]), len(maze)

	d := Deer{}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if maze[y][x] == 'S' {
				d = Deer{x: x, y: y}
			}
		}
	}

	DSA(maze, d, w, h)
}

func DaySixteen() {
	partOne("day16/day16_test.txt")
}