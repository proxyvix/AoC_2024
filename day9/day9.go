package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type DiskMap struct {
	files []int
}

func readData(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading the file:\n%v", err)
	}
	defer file.Close()

	blocks := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range line {
			nums, _ := strconv.Atoi(string(ch))
			blocks = append(blocks, nums)
		}
	}
	return blocks
}

func createDiskMap(blocks []int) DiskMap {
	diskMap := DiskMap{files: []int{}}
	idx := 0

	for i := 0; i < len(blocks); i++ {
		if i%2 != 0 {
			for j := 0; j < blocks[i]; j++ {
				diskMap.files = append(diskMap.files, 0)
			}
		} else {
			for j := 0; j < blocks[i]; j++ {
				diskMap.files = append(diskMap.files, idx)
			}
			idx++
		}
	}
	return diskMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (d *DiskMap) sort(idxShift int) {
	left := idxShift
	right := len(d.files) - 1

	for left < right {
		for left < right && d.files[left] != 0 {
			left++
		}
		for left < right && d.files[right] == 0 {
			right--
		}
		if left < right {
			d.files[left], d.files[right] = d.files[right], d.files[left]
		}
	}
}

func (d *DiskMap) sortFiles(idxShift int) {
	maxID := 0
	for _, file := range d.files {
		if file > maxID {
			maxID = file
		}
	}

	for id := maxID; id > 0; id-- {
		fileStart, fileEnd := -1, -1

		for i := idxShift; i < len(d.files); i++ {
			if d.files[i] == id {
				if fileStart == -1 {
					fileStart = i
				}
				fileEnd = i
			}
		}

		if fileStart == -1 {
			continue
		}

		fileLength := fileEnd - fileStart + 1

		spaceStart, spaceLength := -1, 0
		for i := idxShift; i < fileStart; i++ {
			if d.files[i] == 0 {
				if spaceStart == -1 {
					spaceStart = i
				}
				spaceLength++

				if spaceLength >= fileLength {
					break
				}
			} else {
				spaceStart, spaceLength = -1, 0
			}
		}

		if spaceStart != -1 && spaceLength >= fileLength {
			for i := 0; i < fileLength; i++ {
				d.files[spaceStart+i] = id
			}

			for i := fileStart; i <= fileEnd; i++ {
				d.files[i] = 0
			}
		}
	}
}

func partOne(path string) int {
	blocks := readData(path)
	diskMap := createDiskMap(blocks)
	idxShift := blocks[0]

	diskMap.sort(idxShift)

	checkSum := 0

	for i, f := range diskMap.files {
		checkSum += i * f
	}
	return checkSum
}

func partTwo(path string) int {
	blocks := readData(path)
	diskMap := createDiskMap(blocks)
	idxShift := blocks[0]

	diskMap.sortFiles(idxShift)

	checkSum := 0

	for i, f := range diskMap.files {
		checkSum += i * f
	}
	return checkSum
}

func DayNine() {
	start := time.Now()
	partOne := partOne("day9/day9.txt")
	duration := time.Since(start)
	fmt.Printf("Solution for day 9 part one: %d\nexecution time: %v\n", partOne, duration)

	start = time.Now()
	partTwo := partTwo("day9/day9.txt")
	duration = time.Since(start)
	fmt.Printf("Solution for day 9 part two: %d\nexecution time: %v\n\n", partTwo, duration)
}
