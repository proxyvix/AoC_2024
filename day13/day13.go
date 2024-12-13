package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/mat"
)

type vec2 struct {
	x float64
	y float64
}

func readData(path string) [][3]vec2 {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	vec := [3]vec2{}

	matricies := [][3]vec2{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, ":")

		if cols[0] == "" {
			continue
		}
		if cols[0] == "Button A" {
			nums := strings.Split(cols[1], ",")

			x, _ := strconv.ParseFloat(strings.Split(nums[0], "+")[1], 64)
			y, _ := strconv.ParseFloat(strings.Split(nums[1], "+")[1], 64)

			vec[0] = vec2{x, y}
		} else if cols[0] == "Button B" {
			nums := strings.Split(cols[1], ",")

			x, _ := strconv.ParseFloat(strings.Split(nums[0], "+")[1], 64)
			y, _ := strconv.ParseFloat(strings.Split(nums[1], "+")[1], 64)

			vec[1] = vec2{x, y}
		} else {
			nums := strings.Split(cols[1], ",")

			x, _ := strconv.ParseFloat(strings.Split(nums[0], "=")[1], 64)
			y, _ := strconv.ParseFloat(strings.Split(nums[1], "=")[1], 64)

			vec[2] = vec2{x: x, y: y}

			matricies = append(matricies, vec)
		}
	}
	return matricies
}

func isInteger(n float64) bool {
	return math.Abs(n-math.Round(n)) < 1e-3
}

func isNonNegativeInteger(n float64) bool {
	return isInteger(n) && n >= 0
}

func partOne(path string) int {
	matricies := readData(path)

	score := 0

	for _, m := range matricies {
		A := mat.NewDense(2, 2, []float64{
			m[0].x, m[1].x,
			m[0].y, m[1].y,
		})
		B := mat.NewVecDense(2, []float64{
			m[2].x,
			m[2].y,
		})

		var x mat.VecDense

		err := x.SolveVec(A, B)
		if err != nil {
			fmt.Printf("Error solving the matrix:\n%v", err)
			return 0
		}

		nA, nB := x.At(0, 0), x.At(1, 0)

		if isNonNegativeInteger(nA) && isNonNegativeInteger(nB) {
			score += int(math.Round(nA))*3 + int(math.Round(nB))*1
		}
	}
	return score
}

func partTwo(path string) int {
	matricies := readData(path)

	score := 0

	for _, m := range matricies {
		A := mat.NewDense(2, 2, []float64{
			m[0].x, m[1].x,
			m[0].y, m[1].y,
		})
		B := mat.NewVecDense(2, []float64{
			m[2].x + 1e13,
			m[2].y + 1e13,
		})

		var x mat.VecDense

		err := x.SolveVec(A, B)
		if err != nil {
			fmt.Printf("Error solving the matrix:\n%v", err)
			return 0
		}

		nA, nB := x.At(0, 0), x.At(1, 0)

		if isNonNegativeInteger(nA) && isNonNegativeInteger(nB) {
			score += int(math.Round(nA))*3 + int(math.Round(nB))*1
		}
	}
	return score
}

func DayThirteen() {
	start := time.Now()
	partOne := partOne("day13/day13.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 13 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day13/day13.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 13 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
