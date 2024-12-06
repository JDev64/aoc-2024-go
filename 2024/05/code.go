package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("./2024/05/input-user.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)
	strings.TrimRight(inputString, "\n")

	//fmt.Println(part1(inputString))
	fmt.Println(part2(inputString))
}

func part1(input string) int {
	inputParts := strings.Split(input, "\n\n")

	rules := inputParts[0]
	updates := inputParts[1]

	rulesMap := make(map[int][]int)

	for _, rule := range strings.Split(rules, "\n") {
		nums := strings.Split(rule, "|")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])
		rulesMap[first] = append(rulesMap[first], second)
	}

	sum := 0
	for _, update := range strings.Split(updates, "\n") {
		validUpdate := true
		nums := strings.Split(update, ",")
		for i, num := range nums {
			after := nums[i+1:]
			for j, _ := range after {
				convertedNum, _ := strconv.Atoi(num)
				convertedAfter, _ := strconv.Atoi(after[j])
				if !slices.Contains(rulesMap[convertedNum], convertedAfter) {
					validUpdate = false
					break
				}
			}
			if !validUpdate {
				break
			}
		}
		if validUpdate {
			middle := nums[len(nums)/2]
			num, _ := strconv.Atoi(middle)
			sum += num
		}
	}

	return sum
}

func part2(input string) int {
	inputParts := strings.Split(input, "\n\n")

	rules := inputParts[0]
	updates := inputParts[1]

	rulesMap := make(map[int][]int)

	for _, rule := range strings.Split(rules, "\n") {
		nums := strings.Split(rule, "|")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])
		rulesMap[first] = append(rulesMap[first], second)
	}

	var invalidUpdates []string
	sum := 0

	for _, update := range strings.Split(updates, "\n") {
		validUpdate := true
		nums := strings.Split(update, ",")
		for i, num := range nums {
			after := nums[i+1:]
			for j, _ := range after {
				convertedNum, _ := strconv.Atoi(num)
				convertedAfter, _ := strconv.Atoi(after[j])
				if !slices.Contains(rulesMap[convertedNum], convertedAfter) {
					validUpdate = false
					invalidUpdates = append(invalidUpdates, update)
					break
				}
			}
			if !validUpdate {
				break
			}
		}
	}

	for _, update := range invalidUpdates {

		numbersStr := strings.Split(update, ",")

		numbers := make([]int, len(numbersStr))
		for j, _ := range numbersStr {
			numbers[j], _ = strconv.Atoi(numbersStr[j])
		}

		swapped := true

		for swapped {
			swapped = false
			for j, num := range numbers {
				currentRules := rulesMap[num]
				for _, rule := range currentRules {
					index1 := j
					index2 := indexOf(numbers, rule)
					if index2 > -1 && index1 > index2 {
						numbers[index1], numbers[index2] = numbers[index2], numbers[index1]
						swapped = true
					}
				}
			}
		}

		mid := len(numbers) / 2
		sum += numbers[mid]
	}

	return sum
}

func indexOf(slice []int, val int) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}
