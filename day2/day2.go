package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readData(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		levels := make([]int, len(columns))
		for i, col := range columns {
			levels[i], _ = strconv.Atoi(col)
		}
		reports = append(reports, levels)
	}
	return reports
}

func isIncreasing(row []int) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i] > row[i+1] {
			return false
		}
	}
	return true
}

func isSafe(row []int) bool {
	var localScore int

	for i := 0; i < len(row)-1; i++ {
		diff := row[i] - row[i+1]
		if isIncreasing(row) {
			if diff <= -1 && diff >= -3 {
				localScore += 1
			}
		} else {
			if diff >= 1 && diff <= 3 {
				localScore += 1
			}
		}
	}
	if localScore == len(row)-1 {
		return true
	}
	return false
}

func partOne(path string) int {
	reports := readData(path)

	var safeReports int
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func isSafeDampener(row []int) bool {

	if isSafe(row) {
		return true
	}

	for i := range row {
		var temp []int
		temp = append(temp, row[:i]...)
		temp = append(temp, row[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}
	return false
}

func partTwo(path string) int {
	reports := readData(path)

	var safeReports int
	for _, report := range reports {
		if isSafeDampener(report) {
			safeReports++
		}
	}
	return safeReports
}

func DayTwo() {
	start := time.Now()
	partOne := partOne("day2/day2.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 2 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day2/day2.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 2 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
