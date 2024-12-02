package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"math"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2Solution(input)
	}
	// solve part 1 here
	return part1(input)
}

func part1(input string) any {
	var matrix [][]int

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		numStrings := strings.Split(line, " ")
		matrix = append(matrix, make([]int, len(numStrings)))
		for j, str := range numStrings {
			matrix[i][j], _ = strconv.Atoi(str)
		}
	}

	unsaveLevels := 0
	for i := 0; i < len(matrix); i++ {
		if !isLineValid(matrix[i]) {
			unsaveLevels++
		}
	}
	return len(matrix) - unsaveLevels
}

func part2Solution(input string) any {
	var matrix [][]int

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		numStrings := strings.Split(line, " ")
		matrix = append(matrix, make([]int, len(numStrings)))
		for j, str := range numStrings {
			matrix[i][j], _ = strconv.Atoi(str)
		}
	}

	unsaveLevels := 0
	for i := 0; i < len(matrix); i++ {
		if !isLineValid(matrix[i]) {
			line := matrix[i]
			var lineValid []bool
			for j, _ := range line {
				trimmedLine := make([]int, len(line))
				copy(trimmedLine, line)
				trimmedLine = append(trimmedLine[:j], trimmedLine[j+1:]...)
				if isLineValid(trimmedLine) {
					lineValid = append(lineValid, true)
					break
				} else {
					lineValid = append(lineValid, false)
				}
			}
			if !lineValid[len(lineValid)-1] {
				unsaveLevels += 1
			}
			fmt.Println(unsaveLevels)
		}
	}
	return len(matrix) - unsaveLevels
}

func isLineValid(line []int) bool {
	var asc bool
	var desc bool
	for j := 0; j < len(line); j++ {
		if j == len(line)-1 {
			break
		}

		difference := line[j] - line[j+1]
		if difference == 0 {
			return false
		} else if math.Abs(float64(difference)) > 3 {
			return false
		}

		if !asc && !desc {
			if difference < 0 {
				asc = true
			} else {
				desc = true
			}
			continue
		}

		if difference > 0 && asc {
			return false
		}

		if difference < 0 && desc {
			return false
		}
	}
	return true
}
