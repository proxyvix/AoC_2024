package day8

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Coords struct {
	x int
	y int
}

func readData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	antennas := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		antennas = append(antennas, scanner.Text())
	}
	return antennas
}

func findCoords(antennas []string) map[byte][]Coords {
	hashMap := make(map[byte][]Coords)
	for y := 0; y < len(antennas); y++ {
		for x := 0; x < len(antennas[y]); x++ {
			if antennas[y][x] != '.' {
				hashMap[antennas[y][x]] = append(hashMap[antennas[y][x]], Coords{x: x, y: y})
			}
		}
	}
	return hashMap
}

func isValidAntinode(antennas []string, antinode Coords) bool {
	return antinode.x >= 0 &&
		antinode.x < len(antennas[0]) &&
		antinode.y >= 0 &&
		antinode.y < len(antennas)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findAntinodes(coords map[byte][]Coords, antennas []string, includeHarmonics bool) int {
	uniqueAntiNodes := make(map[Coords]bool)

	for _, c := range coords {
		if len(c) < 2 {
			continue
		}

		for i := 0; i < len(c)-1; i++ {
			for j := i + 1; j < len(c); j++ {

				dx := c[j].x - c[i].x
				dy := c[j].y - c[i].y

				steps := 1
				if includeHarmonics {
					steps = max(len(antennas), len(antennas[0]))
				}

				for k := 1; k <= steps; k++ {
					antinode1 := Coords{x: c[i].x - k*dx, y: c[i].y - k*dy}
					antinode2 := Coords{x: c[j].x + k*dx, y: c[j].y + k*dy}

					if isValidAntinode(antennas, antinode1) {
						uniqueAntiNodes[antinode1] = true
					}
					if isValidAntinode(antennas, antinode2) {
						uniqueAntiNodes[antinode2] = true
					}
				}
			}
		}
	}
	return len(uniqueAntiNodes)
}

func partOne(path string) int {
	antennas := readData(path)
	coords := findCoords(antennas)
	counts := findAntinodes(coords, antennas, false)
	return counts
}

func partTwo(path string) int {
	antennas := readData(path)
	coords := findCoords(antennas)
	counts := findAntinodes(coords, antennas, true)
	return counts
}

func DayEight() {
	start := time.Now()
	partOne := partOne("day8/day8.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 8 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day8/day8.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 8 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
