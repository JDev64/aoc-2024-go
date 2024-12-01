package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"math"
	"sort"
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
		lines := strings.Split(input, "\n")

		var leftList []int
		var rightList []int

		for _, line := range lines {
			lineParts := strings.Split(line, "   ")
			num, err := strconv.Atoi(strings.TrimSpace(lineParts[0]))
			if err != nil {
				panic(err)
			}

			leftList = append(leftList, num)

			num, err = strconv.Atoi(strings.TrimSpace(lineParts[1]))
			if err != nil {
				panic(err)
			}

			rightList = append(rightList, num)
		}

		occurrences := make(map[int]int)
		for _, num := range rightList {
			val, exists := occurrences[num]
			if !exists {
				occurrences[num] = 1
			} else {
				occurrences[num] = val + 1
			}
		}

		distanceSum := 0
		for _, num := range leftList {
			distanceSum += num * occurrences[num]
		}
		return distanceSum
	}
	// solve part 1 here
	lines := strings.Split(input, "\n")

	var leftList []int
	var rightList []int

	for _, line := range lines {
		lineParts := strings.Split(line, "   ")
		num, err := strconv.Atoi(strings.TrimSpace(lineParts[0]))
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, num)

		num, err = strconv.Atoi(strings.TrimSpace(lineParts[1]))
		if err != nil {
			panic(err)
		}

		rightList = append(rightList, num)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	distanceSum := 0
	for i, num := range leftList {
		distanceSum += int(math.Abs(float64(num - rightList[i])))
	}
	return distanceSum
}
