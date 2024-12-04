package day4

import (
	"bufio"
	"fmt"
	"os"
)

func readData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}
	defer file.Close()

	var text []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	return text
}

type Coordinates struct {
	dx int
	dy int
}

func isMatch(x, y, dx, dy int, word string, text []string) bool {
	rows := len(text)
	cols := len(text[0])
	for i := 0; i < len(word); i++ {
		ndx := x + dx*i
		ndy := y + dy*i
		if ndx < 0 || ndx >= rows || ndy < 0 || ndy >= cols || text[ndx][ndy] != word[i] {
			return false
		}
	}
	return true
}

func partOne(path string) int {
	dirs := []Coordinates{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Down right
		{1, -1},  // Down left
		{-1, 1},  // Up right
		{-1, -1}, // Up left
	}

	text := readData(path)
	word := "XMAS"

	var score int
	for i, t := range text {
		for j := range t {
			if t[j] == 'X' {
				for _, dir := range dirs {
					if isMatch(i, j, dir.dx, dir.dy, word, text) {
						score++
					}
				}
			}
		}
	}
	return score
}

func partTwo(path string) int {
	dirs := []Coordinates{
		{1, 1},   // Down right
		{1, -1},  // Down left
		{-1, 1},  // Up right
		{-1, -1}, // Up left
	}

	text := readData(path)
	word := "MAS"

	var score int
	for i, t := range text {
		for j := range t {
			if j+2 < len(text[0]) {
				if text[i][j] == 'M' && text[i][j+2] == 'M' {
					if isMatch(i, j, dirs[0].dx, dirs[0].dy, word, text) && isMatch(i, j+2, dirs[1].dx, dirs[1].dy, word, text) {
						score++
					}
					if isMatch(i, j, dirs[2].dx, dirs[2].dy, word, text) && isMatch(i, j+2, dirs[3].dx, dirs[3].dy, word, text) {
						score++
					}
				}
			}
			if i+2 < len(text) {
				if text[i][j] == 'M' && text[i+2][j] == 'M' {
					if isMatch(i, j, dirs[0].dx, dirs[0].dy, word, text) && isMatch(i+2, j, dirs[2].dx, dirs[2].dy, word, text) {
						score++
					}
					if isMatch(i, j, dirs[1].dx, dirs[1].dy, word, text) && isMatch(i+2, j, dirs[3].dx, dirs[3].dy, word, text) {
						score++
					}
				}
			}
		}
	}
	return score
}

func DayFour() {
	partOne := partOne("day4/day4.txt")
	fmt.Printf("Solution for day 4 part one: %d\n", partOne)

	partTwo := partTwo("day4/day4.txt")
	fmt.Printf("Solution for day 4 part two: %d\n", partTwo)
}
