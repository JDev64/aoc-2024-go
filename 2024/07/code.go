package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("./2024/07/input-user.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)

	fmt.Println(part1(inputString))
	fmt.Println(part2(inputString))
}

func part1(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	var solutions []int

	for _, line := range lines {
		// convert line to solution and problem
		lineParts := strings.Split(line, ":")
		solution, _ := strconv.Atoi(lineParts[0])
		problem := strings.TrimSpace(lineParts[1])
		problemParts := strings.Split(problem, " ")
		numbers := make([]int, len(problemParts))
		for i, _ := range problemParts {
			numbers[i], _ = strconv.Atoi(problemParts[i])
		}
		// fill operator list with default operator
		operators := make([]string, len(numbers)-1)
		for i, _ := range operators {
			operators[i] = "+"
		}

		// find solution
		for {
			result := numbers[0]
			for i, _ := range operators {
				if operators[i] == "+" {
					result += numbers[i+1]
				} else {
					result *= numbers[i+1]
				}
			}
			if result == solution {
				solutions = append(solutions, solution)
				break
			} else {
				// find next combination
				if !slices.Contains(operators, "+") {
					break
				}
				takeNext := false
				if operators[len(operators)-1] == "+" {
					operators[len(operators)-1] = "*"
				} else {
					operators[len(operators)-1] = "+"
					takeNext = true
				}
				for i := len(operators) - 2; i >= 0; i-- {
					if takeNext {
						if operators[i] == "+" {
							operators[i] = "*"
							break
						} else {
							operators[i] = "+"
						}
					} else {
						break
					}
				}
			}
		}
	}

	// sum up solutions
	total := 0
	for _, solution := range solutions {
		total += solution
	}
	return total
}

func part2(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	var solutions []int

	for _, line := range lines {
		// convert line to solution and problem
		lineParts := strings.Split(line, ":")
		solution, _ := strconv.Atoi(lineParts[0])
		problem := strings.TrimSpace(lineParts[1])
		problemParts := strings.Split(problem, " ")
		numbers := make([]int, len(problemParts))
		for i, _ := range problemParts {
			numbers[i], _ = strconv.Atoi(problemParts[i])
		}
		// fill operator list with default operator
		operators := make([]string, len(numbers)-1)
		for i, _ := range operators {
			operators[i] = "+"
		}

		// find solution
		for {
			result := numbers[0]
			for i, _ := range operators {
				if operators[i] == "+" {
					result += numbers[i+1]
				} else if operators[i] == "*" {
					result *= numbers[i+1]
				} else {
					combindedNumber := fmt.Sprintf("%d%d", result, numbers[i+1])
					number, _ := strconv.Atoi(combindedNumber)
					result = number
				}
			}
			if result == solution {
				solutions = append(solutions, solution)
				break
			} else {
				// find next combination
				if !slices.Contains(operators, "*") && !slices.Contains(operators, "+") {
					break
				}
				takeNext := false
				if operators[len(operators)-1] == "+" {
					operators[len(operators)-1] = "*"
				} else if operators[len(operators)-1] == "*" {
					operators[len(operators)-1] = "||"
				} else {
					operators[len(operators)-1] = "+"
					takeNext = true
				}
				for i := len(operators) - 2; i >= 0; i-- {
					if takeNext {
						if operators[i] == "+" {
							operators[i] = "*"
							break
						} else if operators[i] == "*" {
							operators[i] = "||"
							break
						} else {
							operators[i] = "+"
						}
					} else {
						break
					}
				}
			}
		}
	}

	// sum up solutions
	total := 0
	for _, solution := range solutions {
		total += solution
	}
	return total
}
