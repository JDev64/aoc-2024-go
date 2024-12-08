package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("./2024/08/input-user.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)

	fmt.Println(part1(inputString))
	fmt.Println(part2(inputString))
}

func part1(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\r\n"), "\r\n")
	height := len(lines)
	width := len(lines[0])
	grid := make([][]byte, height)

	for i := range grid {
		grid[i] = make([]byte, width)
		for j := range lines[i] {
			grid[i][j] = lines[i][j]
		}
	}
	evaluatedAntennaTypes := make([]byte, 0)
	var antinodes []Point
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '.' || grid[i][j] == '#' {
				continue
			}
			antenna := grid[i][j]
			if !slices.Contains(evaluatedAntennaTypes, antenna) {
				similarAntennas := findSimilarAntennas(grid, Point{j, i}, antenna)
				if len(similarAntennas) > 1 {
					foundAntiNodes := calculateAntinodes(grid, similarAntennas)
					for _, antinode := range foundAntiNodes {
						if !slices.Contains(antinodes, antinode) {
							antinodes = append(antinodes, antinode)
						}
					}
				}

			}

		}
	}

	return len(antinodes)
}

func part2(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\r\n"), "\r\n")
	height := len(lines)
	width := len(lines[0])
	grid := make([][]byte, height)

	for i := range grid {
		grid[i] = make([]byte, width)
		for j := range lines[i] {
			grid[i][j] = lines[i][j]
		}
	}
	evaluatedAntennaTypes := make([]byte, 0)
	antennaSum := 0
	var antinodes []Point
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '.' || grid[i][j] == '#' {
				continue
			}
			antenna := grid[i][j]
			if !slices.Contains(evaluatedAntennaTypes, antenna) {
				evaluatedAntennaTypes = append(evaluatedAntennaTypes, antenna)
				similarAntennas := findSimilarAntennas(grid, Point{j, i}, antenna)
				if len(similarAntennas) >= 2 {
					antennaSum += len(similarAntennas)
					foundAntiNodes := calculateAntinodes(grid, similarAntennas)
					var filteredAntiNodes []Point
					for _, antinode := range foundAntiNodes {
						if grid[antinode.Y][antinode.X] == '.' {
							filteredAntiNodes = append(filteredAntiNodes, antinode)
						}
					}
					foundAntiNodes = filteredAntiNodes
					for _, antinode := range foundAntiNodes {
						if !slices.Contains(antinodes, antinode) {
							antinodes = append(antinodes, antinode)
						}
					}
				}

			}

		}
	}
	return len(antinodes) + antennaSum
}

func findSimilarAntennas(grid [][]byte, point Point, antenna byte) []Point {
	var similarAntennas []Point

	for i := point.Y; i < len(grid); i++ {
		for j := point.X; j < len(grid[i]); j++ {
			if grid[i][j] == antenna {
				similarAntennas = append(similarAntennas, Point{j, i})
			}
		}
		point.X = 0
	}
	return similarAntennas
}

func calculateAntinodes(grid [][]byte, antennas []Point) []Point {
	var antinodes []Point

	for i := 0; i < len(antennas); i++ {
		var temp []Point
		for _, antenna := range antennas {
			temp = append(temp, antenna)
		}
		temp = append(temp[:i], temp[i+1:]...)

		for _, tempAntenna := range temp {
			vector := antennas[i].CalculateVector(tempAntenna)
			minus := tempAntenna.Add(vector)
			plus := antennas[i].Add(vector.Invert())
			for {
				if !minus.IsInsideGrid(grid) {
					break
				}
				if !slices.Contains(antinodes, minus) {
					antinodes = append(antinodes, minus)
				}
				minus = minus.Add(vector)
			}
			for {
				if !plus.IsInsideGrid(grid) {
					break
				}
				if !slices.Contains(antinodes, plus) {
					antinodes = append(antinodes, plus)
				}
				plus = plus.Add(vector.Invert())
			}
		}
	}

	return antinodes
}

type Point struct {
	X int
	Y int
}

func (p Point) Add(p2 Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y}
}

func (p Point) Invert() Point {
	return Point{-p.X, -p.Y}
}

func (p Point) CalculateVector(p2 Point) Point {
	return p2.Add(p.Invert())
}

func (p Point) IsInsideGrid(grid [][]byte) bool {
	return p.X >= 0 && p.X < len(grid[0]) && p.Y >= 0 && p.Y < len(grid)
}

func createDeepGridCopy(grid [][]byte) [][]byte {
	copyGrid := make([][]byte, len(grid))
	for i := range grid {
		copyGrid[i] = make([]byte, len(grid[i]))
		copy(copyGrid[i], grid[i])
	}
	return copyGrid
}
