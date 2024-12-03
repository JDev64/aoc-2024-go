package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("./2024/03/input-user.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)

	fmt.Println(part1(inputString))
	fmt.Println(part2(inputString))
}

func part1(input string) int {
	regex := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	matches := regex.FindAll([]byte(input), -1)

	sum := 0

	for _, match := range matches {
		finding := string(match)
		finding = strings.TrimLeft(finding, "mul(")
		finding = strings.TrimRight(finding, ")")
		nums := strings.Split(finding, ",")
		factor1, _ := strconv.Atoi(nums[0])
		factor2, _ := strconv.Atoi(nums[1])
		sum += factor1 * factor2
	}
	return sum
}

func part2(input string) int {
	regex := regexp.MustCompile("(mul\\([0-9]+,[0-9]+\\))|(don't)|(do)")
	matches := regex.FindAll([]byte(input), -1)
	sum := 0

	active := true

	for _, match := range matches {
		finding := string(match)

		if finding == "don't" {
			active = false
		} else if finding == "do" {
			active = true
		} else {
			if active {
				finding = strings.TrimLeft(finding, "mul(")
				finding = strings.TrimRight(finding, ")")
				nums := strings.Split(finding, ",")
				factor1, _ := strconv.Atoi(nums[0])
				factor2, _ := strconv.Atoi(nums[1])
				sum += factor1 * factor2
			}
		}
	}
	return sum
}
