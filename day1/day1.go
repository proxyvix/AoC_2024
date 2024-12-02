package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getLists() ([]int, []int) {
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var listOne []int
	var listTwo []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if len(columns) >= 2 {
			firstVal, _ := strconv.Atoi(columns[0])
			secondVal, _ := strconv.Atoi(columns[1])
			listOne = append(listOne, firstVal)
			listTwo = append(listTwo, secondVal)
		}
	}

	return listOne, listTwo
}

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func totalDistance(arr1, arr2 []int) int {
	diff := 0.0
	for i := 0; i < len(arr1); i++ {
		diff += math.Abs(float64(arr1[i]) - float64(arr2[i]))
	}

	return int(diff)
}

func similarityScore(arr1, arr2 []int) int {
	// var scores []int

	n := len(arr1)
	sumOfScores := 0

	for i := 0; i < n; i++ {
		score := 0
		for j := 0; j < n; j++ {
			if arr1[i] == arr2[j] {
				score += 1
			}
		}
		sumOfScores += arr1[i] * score
	}

	return sumOfScores
}

func dayOne() {
	listOne, listTwo := getLists()

	bubbleSort(listOne[:])
	bubbleSort(listTwo[:])

	partOne := totalDistance(listOne, listTwo)
	fmt.Printf("Solution for day 1 part one: %d\n", partOne)

	partTwo := similarityScore(listOne, listTwo)
	fmt.Printf("Solution for day 1 part two: %d\n", partTwo)
}
