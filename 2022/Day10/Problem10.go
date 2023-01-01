package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/collection"
	"github.com/srinchow/adventOfCode/utils/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2022/Day10/input.txt")

	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
	}
	defer file.CloseFile(f)

	operations := file.ParseFile(f)[0]
	//part1(operations)
	part2(operations)
}

func part1(operations []string) {
	cycle := 0
	x := 1
	res := 0
	cycleToConsider := []int{20, 60, 100, 140, 180, 220}
	for _, op := range operations {
		if op == "noop" {
			cycle += 1
			if collection.Contains(cycleToConsider, cycle) {
				res += cycle * x
			}
		} else {
			cycle += 1
			if collection.Contains(cycleToConsider, cycle) {
				res += cycle * x
			}
			cycle += 1
			if collection.Contains(cycleToConsider, cycle) {
				res += cycle * x
			}
			x += getInt(strings.Fields(op)[1])
		}
	}

	fmt.Println(res)
}

func part2(operations []string) {
	spritePositions := []int{0, 1, 2}
	currentCRT := -1
	visible := map[int]bool{}
	for _, op := range operations {
		if op == "noop" {
			currentCRT += 1
			if collection.Contains(spritePositions, currentCRT%40) {
				visible[currentCRT] = true
			}
		} else {
			currentCRT += 1
			if collection.Contains(spritePositions, currentCRT%40) {
				visible[currentCRT] = true
			}
			currentCRT += 1
			if collection.Contains(spritePositions, currentCRT%40) {
				visible[currentCRT] = true
			}
			collection.IncrementElements(spritePositions, getInt(strings.Fields(op)[1]))
		}
	}

	prettyPrint := make([][]string, 6)

	for idx := range prettyPrint {
		prettyPrint[idx] = make([]string, 40)
		for i := 0; i < len(prettyPrint[idx]); i++ {
			prettyPrint[idx][i] = "."
		}
	}

	for key := range visible {
		row := key / 40
		col := key % 40

		prettyPrint[row][col] = "#"
	}

	for idx := range prettyPrint {
		fmt.Println(strings.Join(prettyPrint[idx], " "))
	}

}

func getInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer")
		return -1
	}
	return res
}
