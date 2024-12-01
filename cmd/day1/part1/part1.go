package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	sort.Ints(leftList)
	sort.Ints(rightList)

	distanceSum := 0
	for i, num := range leftList {
		distanceSum += int(math.Abs(float64(num - rightList[i])))
	}
	fmt.Println(distanceSum)
}
