package day12

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"time"
)

func readData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading the file:\n%v", err)
	}
	defer file.Close()

	plants := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		plants = append(plants, scanner.Text())
	}
	return plants
}

func isValid(x, y, w, h int) bool {
	if x >= 0 && x < w && y >= 0 && y < h {
		return true
	}
	return false
}

func bfs(garden []string, visited [][]bool, x, y int) (int, int) {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	h := len(garden)
	w := len(garden[0])
	area := 1
	perimeter := 0

	q := list.New()
	q.PushBack([2]int{x, y})

	visited[y][x] = true

	for q.Len() > 0 {
		element := q.Front()
		current := element.Value.([2]int)
		q.Remove(element)

		xx, yy := current[0], current[1]

		for _, dir := range directions {
			nx := xx + dir[0]
			ny := yy + dir[1]

			if !isValid(nx, ny, w, h) || garden[ny][nx] != garden[y][x] {
				perimeter++
			} else if !visited[ny][nx] {
				visited[ny][nx] = true
				area++
				q.PushBack([2]int{nx, ny})
			}
		}
	}
	return area, perimeter
}

func getSideCount(sides [][3]int) int {
	sideMap := make(map[[3]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0] == sides[j][0] {
			return sides[i][1] < sides[j][1]
		}
		return sides[i][0] < sides[j][0]
	})

	sideCount := 0

	for _, s := range sides {
		getCombinations := getNeighbors(s[0], s[1])
		combFound := false

		for _, c := range getCombinations {
			c[2] = s[2]
			if _, found := sideMap[c]; found {
				combFound = true
			}
		}
		if !combFound {
			sideCount++
		}

		sideMap[s] = true

	}

	return sideCount
}

func getNeighbors(i, j int) [][3]int {
	return [][3]int{
		{i - 1, j, 0},
		{i + 1, j, 1},
		{i, j - 1, 2},
		{i, j + 1, 3},
	}
}

func dfs(farm []string, visited [][]bool, i, j int, sides *[][3]int) (int, int) {

	visited[i][j] = true
	neighbors := getNeighbors(i, j)
	perimeter := 0
	area := 1
	for _, n := range neighbors {
		ni, nj := n[0], n[1]
		if ni < 0 || ni >= len(farm) || nj < 0 || nj >= len(farm[ni]) || farm[ni][nj] != farm[i][j] {
			if sides != nil {
				*sides = append(*sides, [3]int{ni, nj, n[2]})
			}
			perimeter++
		} else if !visited[ni][nj] {
			a, p := dfs(farm, visited, ni, nj, sides)
			area += a
			perimeter += p
		}
	}
	return area, perimeter
}

func partTwo(path string) int {
	farm := readData(path)
	visited := make([][]bool, len(farm))
	for i := range visited {
		visited[i] = make([]bool, len(farm[i]))
	}
	sum := 0
	for i := 0; i < len(farm); i++ {
		for j := 0; j < len(farm[i]); j++ {
			if !visited[i][j] {
				sides := make([][3]int, 0)
				area, _ := dfs(farm, visited, i, j, &sides)
				sideCount := getSideCount(sides)
				sum += area * sideCount
			}
		}
	}
	return sum
}

func partOne(path string) int {
	garden := readData(path)

	h := len(garden)
	w := len(garden[0])

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	totalPrice := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if !visited[y][x] {
				area, perimeter := bfs(garden, visited, x, y)
				totalPrice += area * perimeter
			}
		}
	}
	return totalPrice
}

func DayTwelve() {
	start := time.Now()
	partOne := partOne("day12/day12.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 12 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day12/day12.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 12 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
