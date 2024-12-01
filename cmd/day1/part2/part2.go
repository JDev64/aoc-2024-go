package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./cmd/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var leftList []int
	var rightList []int

	for scanner.Scan() {
		scannedText := scanner.Text()
		lineParts := strings.Split(scannedText, "   ")
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
	fmt.Println(distanceSum)
}
