package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readData(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	stones := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := strings.Fields(scanner.Text())
		for _, char := range chars {
			stone, _ := strconv.Atoi(char)
			stones = append(stones, stone)
		}
	}
	return stones
}

func blink(stones []int) []int {
	newStones := []int{}

	for j := 0; j < len(stones); j++ {
		if stones[j] == 0 {
			newStones = append(newStones, 1)
		} else {
			charStone := strconv.Itoa(stones[j])

			if len(charStone)%2 != 0 {
				newStones = append(newStones, stones[j]*2024)
			} else {
				num1, _ := strconv.Atoi(charStone[:len(charStone)/2])
				num2, _ := strconv.Atoi(charStone[len(charStone)/2:])
				newStones = append(newStones, num1, num2)
			}
		}
	}
	return newStones
}

func blinkPartTwo(stones []int, iterations int) map[int]int {
	stoneCounts := make(map[int]int)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for i := 0; i < iterations; i++ {
		nextStoneCounts := make(map[int]int)
		for stone, count := range stoneCounts {
			if stone == 0 {
				nextStoneCounts[1] += count
			} else {
				charStone := strconv.Itoa(stone)
				if len(charStone)%2 != 0 {
					nextStoneCounts[stone*2024] += count
				} else {
					num1, _ := strconv.Atoi(charStone[:len(charStone)/2])
					num2, _ := strconv.Atoi(charStone[len(charStone)/2:])
					nextStoneCounts[num1] += count
					nextStoneCounts[num2] += count
				}
			}
		}
		stoneCounts = nextStoneCounts
	}
	return stoneCounts
}

func partOne(path string) int {
	stones := readData(path)
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}
	return len(stones)
}

func partTwo(path string) int {
	stones := readData(path)
	stoneCounts := blinkPartTwo(stones, 75)
	counts := 0
	for _, count := range stoneCounts {
		counts += count
	}
	return counts
}

func DayEleven() {
	start := time.Now()
	result := partOne("day11/day11.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 11 part one: %d\nExecution time: %v\n", result, duration)

	start = time.Now()
	result = partTwo("day11/day11.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 11 part two: %d\nExecution time: %v\n\n", result, duration)
}
