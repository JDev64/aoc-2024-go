package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("./2024/06/input-user.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)

	fmt.Println(part1(inputString))
	fmt.Println(part2(inputString))
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

	// get start position
	var position Point

	// unique fields
	visitedFields := make(map[Point]bool)

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == '^' {
				position = Point{
					x: j,
					y: i,
				}
			}
		}
	}

	visitedFields[position] = true
	guardState := UP

	for {
		if position.isLeavingGrid(grid, guardState) {
			break
		}

		// Change Rotation if walking into object
		if walkingIntoObject(grid, guardState, position) {
			guardState = getNextState(guardState)
			continue
		}

		// walk
		position = position.getNextPosition(guardState)
		if position.isInsideGrid(grid) {
			_, ok := visitedFields[position]
			if !ok {
				visitedFields[position] = true
			}
		}
	}

	return len(visitedFields)
}

func part2(input string) int {
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

	// get start position
	var position Point

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == '^' {
				position = Point{
					x: j,
					y: i,
				}
			}
		}
	}

	guardState := UP
	loopsFound := 0

	positions := getWalkingPositions(grid, position, guardState)

	for _, possibleField := range positions {
		gridCopy := createDeepGridCopy(grid)
		gridCopy[possibleField.y][possibleField.x] = '#'
		if isGuardWalkingInALoop(gridCopy, position, guardState) {
			loopsFound++
		}
	}

	return loopsFound
}

func isGuardWalkingInALoop(grid [][]byte, position Point, state State) bool {
	visitedPositions := make(map[Key]bool)
	for {
		if _, ok := visitedPositions[Key{
			position: position,
			state:    state,
		}]; ok {
			return true
		}

		if position.isLeavingGrid(grid, state) {
			return false
		}

		// Change Rotation if walking into object
		if walkingIntoObject(grid, state, position) {
			state = getNextState(state)
			continue
		}

		// walk
		if position.isInsideGrid(grid) {
			visitedPositions[Key{
				position: position,
				state:    state,
			}] = true
		}
		position = position.getNextPosition(state)

	}
}

type State int

const (
	UP State = iota
	RIGHT
	DOWN
	LEFT
)

type Key struct {
	position Point
	state    State
}

func getNextState(state State) State {
	switch state {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		return UP
	}
}

func walkingIntoObject(grid [][]byte, state State, position Point) bool {
	switch state {
	case UP:
		return grid[position.y-1][position.x] == '#'
	case RIGHT:
		return grid[position.y][position.x+1] == '#'
	case DOWN:
		return grid[position.y+1][position.x] == '#'
	case LEFT:
		return grid[position.y][position.x-1] == '#'
	default:
		return false
	}
}

type Point struct {
	x, y int
}

func (p Point) isInsideGrid(grid [][]byte) bool {
	return p.y < len(grid) && p.y > -1 && p.x < len(grid[0]) && p.x > -1
}

func (p Point) getNextPosition(state State) Point {
	switch state {
	case UP:
		return Point{x: p.x, y: p.y - 1}
	case RIGHT:
		return Point{x: p.x + 1, y: p.y}
	case DOWN:
		return Point{x: p.x, y: p.y + 1}
	case LEFT:
		return Point{x: p.x - 1, y: p.y}
	default:
		return Point{x: p.x, y: p.y}
	}
}

func (p Point) isLeavingGrid(grid [][]byte, state State) bool {
	return p.isInsideGrid(grid) && !p.getNextPosition(state).isInsideGrid(grid)
}

func getWalkingPositions(grid [][]byte, position Point, state State) []Point {
	startPosition := position
	walkingPositions := make([]Point, 0)

	for {
		if position.isLeavingGrid(grid, state) {
			break
		}

		// Change Rotation if walking into object
		if walkingIntoObject(grid, state, position) {
			state = getNextState(state)
			continue
		}

		// walk
		position = position.getNextPosition(state)
		if position.isInsideGrid(grid) && position != startPosition && !slices.Contains(walkingPositions, position) {
			walkingPositions = append(walkingPositions, position)
		}
	}

	return walkingPositions
}

func createDeepGridCopy(grid [][]byte) [][]byte {
	copyGrid := make([][]byte, len(grid))
	for i := range grid {
		copyGrid[i] = make([]byte, len(grid[i]))
		copy(copyGrid[i], grid[i])
	}
	return copyGrid
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println("----------")
}
