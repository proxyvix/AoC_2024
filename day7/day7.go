package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Operations struct {
	op []rune
}

func readData(path string) map[int][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	equations := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		testVal, _ := strconv.Atoi(parts[0])
		numStrings := strings.Fields(parts[1])

		eq := []int{}
		for _, s := range numStrings {
			num, _ := strconv.Atoi(s)
			eq = append(eq, num)
		}
		equations[testVal] = eq
	}
	return equations
}

func digitCount(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func intPow(base, exp int) int {
	res := 1
	for exp > 0 {
		res *= base
		exp--
	}
	return res
}

func concat(a, b int) int {
	return a*intPow(10, digitCount(b)) + b
}

func operate(eq []int, ops []rune) int {
	res := eq[0]
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case '+':
			res += eq[i+1]
		case '*':
			res *= eq[i+1]
		case '|':
			res = concat(res, eq[i+1])
		}
	}
	return res
}

func generateCombinations(n int, curr []rune, res *[][]rune, ops Operations) {
	if len(curr) == n {
		*res = append(*res, append([]rune{}, curr...))
		return
	}
	for _, op := range ops.op {
		curr = append(curr, op)
		generateCombinations(n, curr, res, ops)
		curr = curr[:len(curr)-1]
	}
}

func solve(testVal int, eq []int, ops Operations) bool {
	numOperators := len(eq) - 1
	allCombinations := [][]rune{}
	generateCombinations(numOperators, []rune{}, &allCombinations, ops)

	for _, combination := range allCombinations {
		if operate(eq, combination) == testVal {
			return true
		}
	}
	return false
}

func partOne(path string) int {
	equations := readData(path)
	res := 0
	ops := Operations{
		op: []rune{'+', '*'},
	}

	for testVal, eq := range equations {
		if solve(testVal, eq, ops) {
			res += testVal
		}
	}
	return res
}

func partTwo(path string) int {
	equations := readData(path)
	total := 0
	ops := Operations{
		op: []rune{'+', '*', '|'},
	}

	for testVal, eq := range equations {
		if solve(testVal, eq, ops) {
			total += testVal
		}
	}
	return total
}

func DaySeven() {
	start := time.Now()
	partOne := partOne("day7/day7.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 7 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day7/day7.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 7 part two: %d\nexecution time: %v\n\n", partTwo, duration)

}
