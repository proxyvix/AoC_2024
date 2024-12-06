package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func readData(path string) (map[int][]int, [][]int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
	}
	defer file.Close()

	var (
		list1       []string
		list2       []string
		currentList *[]string
	)

	currentList = &list1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currentList = &list2
			continue
		}
		*currentList = append(*currentList, line)
	}

	ruleSet := make(map[int][]int)
	for _, l := range list1 {
		nums := strings.Split(l, "|")

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		ruleSet[num1] = append(ruleSet[num1], num2)
	}

	numSet := [][]int{}
	for _, l := range list2 {
		nums := strings.Split(l, ",")

		set := []int{}

		for _, n := range nums {
			num, _ := strconv.Atoi(n)
			set = append(set, num)
		}
		numSet = append(numSet, set)
	}
	return ruleSet, numSet
}

func isCorrect(ruleSet map[int][]int, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		items := ruleSet[update[i+1]]
		notValid := slices.Contains(items, update[i])
		if notValid {
			return false
		}
	}
	return true
}

func partOne(path string) int {
	ruleSet, updates := readData(path)

	var pageNumber int
	for i := 0; i < len(updates); i++ {
		if isCorrect(ruleSet, updates[i]) {
			pageNumber += updates[i][len(updates[i])/2]
		}
	}
	return pageNumber
}

func sortUpdate(ruleSet map[int][]int, update []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for j := 0; j < len(update)-1; j++ {
			if slices.Contains(ruleSet[update[j+1]], update[j]) {
				temp := update[j]
				update[j] = update[j+1]
				update[j+1] = temp
				swapped = true
			}
		}
	}
	return update
}

func partTwo(path string) int {
	ruleSet, updates := readData(path)

	var pageNumber int
	for i := 0; i < len(updates); i++ {
		if !isCorrect(ruleSet, updates[i]) {
			updates[i] = sortUpdate(ruleSet, updates[i])
			pageNumber += updates[i][len(updates[i])/2]
		}
	}
	return pageNumber
}

func DayFive() {
	start := time.Now()
	partOne := partOne("day5/day5.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 5 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day5/day5.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 5 part two: %d\nexecution time: %v\n\n", partTwo, duration)

}
