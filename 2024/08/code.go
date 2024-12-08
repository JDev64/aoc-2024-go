package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("./2024/08/input-example.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)

	fmt.Println(part1(inputString))
	part2(inputString)
}

func part1(input string) int {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
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
	return 0
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
			plus := antennas[i].Add(vector)
			minus := antennas[i].Add(vector.Invert())
			if plus.IsInsideGrid(grid) && !slices.Contains(antinodes, plus) {
				antinodes = append(antinodes, plus)
			}
			if minus.IsInsideGrid(grid) && !slices.Contains(antinodes, minus) {
				antinodes = append(antinodes, minus)

			}
		}

		/*for j := i + 1; j < len(antennas); j++ {
			vector1 := antennas[i].CalculateVector(antennas[j])
			antinode1 := antennas[i].Add(vector1)
			if antinode1.IsInsideGrid(grid) && !slices.Contains(antinodes, antinode1) {
				antinodes = append(antinodes, antinode1)
			}

			vector2 := antennas[j].CalculateVector(antennas[i])
			antinode2 := antennas[j].Add(vector2)
			if antinode2.IsInsideGrid(grid) && !slices.Contains(antinodes, antinode2) {
				antinodes = append(antinodes, antinode2)
			}
		}*/
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
