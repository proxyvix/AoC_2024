package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	px int
	py int
}

type Velocity struct {
	vx int
	vy int
}

type Robots struct {
	p Position
	v Velocity
}

const (
	W = 101
	H = 103
)

var triangle = [][]int{
	{0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 1, 0, 0, 0},
	{0, 0, 1, 1, 1, 1, 1, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 0},
	{1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func readData(path string) []Robots {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening the file:\n%v", err)
	}
	defer file.Close()

	robots := []Robots{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		p := strings.Split(line[0], "=")
		v := strings.Split(line[1], "=")

		px, _ := strconv.Atoi(strings.Split(p[1], ",")[0])
		py, _ := strconv.Atoi(strings.Split(p[1], ",")[1])

		vx, _ := strconv.Atoi(strings.Split(v[1], ",")[0])
		vy, _ := strconv.Atoi(strings.Split(v[1], ",")[1])

		robots = append(robots, Robots{
			p: Position{px: px, py: py},
			v: Velocity{vx: vx, vy: vy},
		})
	}
	return robots
}

func (r *Robots) move() {
	r.p.px = (r.p.px + r.v.vx + W) % W
	r.p.py = (r.p.py + r.v.vy + H) % H
}

func checkTree(cnt [][]int) bool {
	N := 10
	sm := 0

	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			if j >= N && j <= H-N {
				continue
			}
			sm += cnt[i][j]
		}
	}
	return sm < 10
}

func partOne(path string) int {
	robots := readData(path)

	midW := W / 2
	midH := H / 2

	quadrants := [4]int{}

	safetyFactor := 1

	for i := range robots {
		for j := 0; j < 100; j++ {
			robots[i].move()
		}

		if robots[i].p.px < midW && robots[i].p.py < midH {
			quadrants[0]++
		}
		if robots[i].p.px > midW && robots[i].p.py < midH {
			quadrants[1]++
		}
		if robots[i].p.px < midW && robots[i].p.py > midH {
			quadrants[2]++
		}
		if robots[i].p.px > midW && robots[i].p.py > midH {
			quadrants[3]++
		}
	}

	for i := range quadrants {
		safetyFactor *= quadrants[i]
	}
	return safetyFactor
}

func partTwo(path string) int {
	robots := readData(path)
	floorPlan := make([][]int, H)
	for i := range floorPlan {
		floorPlan[i] = make([]int, W)
	}

	for _, r := range robots {
		floorPlan[r.p.py][r.p.px]++
	}

	for steps := 1; ; steps++ {
		for i, r := range robots {
			floorPlan[r.p.py][r.p.px]-- // Remove robot's old position
			r.move()
			floorPlan[r.p.py][r.p.px]++ // Add robot's new position
			robots[i] = r
		}

		// Check for matching triangle pattern
		for x := 0; x <= H-len(triangle); x++ {
			for y := 0; y <= W-len(triangle[0]); y++ {
				if isMatch(floorPlan, x, y) {
					return steps
				}
			}
		}
	}
}

// Checks if the triangle pattern matches a given region in the floor plan
func isMatch(floorPlan [][]int, x, y int) bool {
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[0]); j++ {
			if triangle[i][j] == 1 && floorPlan[x+i][y+j] == 0 {
				return false
			}
		}
	}
	return true
}

func DayFourteen() {
	start := time.Now()
	partOne := partOne("day14/day14.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 14 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day14/day14.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 14 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
