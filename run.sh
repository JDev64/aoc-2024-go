#!/bin/bash
set -euf -o pipefail

# functions
function echogrey() {
	echo -e "\033[0;90m$1\033[0m"
}

function template() {
	cat <<EOF
package main

import (
	"os"
)

func main() {
	fileContent, err := os.ReadFile("input-user.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(fileContent)

	part1(inputString)
	part2(inputString)
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
EOF
}

# two args YEAR and DAY
YEAR="${1:-}"
DAY="${2:-}"
if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
	echo "Usage: $0 <YEAR> <DAY>"
	exit 1
fi
# pad DAY to 2 digits
DAY=$(printf "%02d" $DAY)
DIR="./$YEAR/$DAY"
# create missing files as needed
if [ ! -d "$DIR" ]; then
	mkdir -p "$DIR"
	echogrey "Created directory $DIR"
fi
if [ ! -f "$DIR/code.go" ]; then
	template >"$DIR/code.go"
	echogrey "Created file code.go"
fi
# go run
cd "$DIR"
touch input-example.txt
touch input-user.txt

