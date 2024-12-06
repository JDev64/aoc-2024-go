package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("./2024/04/input-user.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)

	fmt.Println(part1(inputString))
	fmt.Println(part2(inputString))
}

func part1(input string) int {
	input = strings.TrimRight(input, "\n")
	lines := strings.Split(input, "\n")
	width := len(lines[0])

	count := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < width-3; j++ {
			if lines[i][j] == 'X' && lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
				count++
			}

			if lines[i][j] == 'S' && lines[i][j+1] == 'A' && lines[i][j+2] == 'M' && lines[i][j+3] == 'X' {
				count++
			}
		}
	}

	for i := 0; i < width; i++ {
		for j := 0; j < len(lines)-3; j++ {
			if lines[j][i] == 'X' && lines[j+1][i] == 'M' && lines[j+2][i] == 'A' && lines[j+3][i] == 'S' {
				count++
			}

			if lines[j][i] == 'S' && lines[j+1][i] == 'A' && lines[j+2][i] == 'M' && lines[j+3][i] == 'X' {
				count++
			}
		}
	}

	// top left --> bottom right
	for i := 0; i < len(lines)-3; i++ {
		for j := 0; j < width-3; j++ {
			if lines[i][j] == 'X' && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
				count++
			}

			if lines[i][j] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' && lines[i+3][j+3] == 'X' {
				count++
			}
		}
	}

	// bottom left --> top right
	for i := 3; i < len(lines); i++ {
		for j := 0; j < width-3; j++ {
			if lines[i][j] == 'X' && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
				count++
			}

			if lines[i][j] == 'S' && lines[i-1][j+1] == 'A' && lines[i-2][j+2] == 'M' && lines[i-3][j+3] == 'X' {
				count++
			}
		}
	}

	return count
}

func part2(input string) int {
	input = strings.TrimRight(input, "\n")
	lines := strings.Split(input, "\n")
	width := len(lines[0])

	count := 0

	// top left --> bottom right
	for i := 0; i < len(lines)-2; i++ {
		for j := 0; j < width-2; j++ {
			if lines[i][j] == 'M' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'S' {
				if lines[i][j+2] == 'M' && lines[i+1][j+1] == 'A' && lines[i+2][j] == 'S' {
					count++
				}

				if lines[i][j+2] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j] == 'M' {
					count++
				}
			}
			if lines[i][j] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' {
				if lines[i][j+2] == 'M' && lines[i+1][j+1] == 'A' && lines[i+2][j] == 'S' {
					count++
				}

				if lines[i][j+2] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j] == 'M' {
					count++
				}
			}
		}
	}

	return count
}
