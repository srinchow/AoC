package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils"
	"os"
	"regexp"
	"strconv"
)

func getInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return value
}

func main() {
	file, err := os.Open("./Day5/input.txt")

	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
	}

	defer utils.CloseFile(file)

	response := utils.ParseFile(file)

	movements := response[1]
	stackArrangement := response[0]

	var stacks []*utils.Stack

	for range stackArrangement[len(stackArrangement)-1] {
		stacks = append(stacks, &utils.Stack{})
	}

	for i := len(stackArrangement) - 2; i >= 0; i-- {
		for index, val := range stackArrangement[i] {
			if val == 32 {
				continue
			}
			stacks[index].Push(val)
		}
	}

	//move 2 from 9 to 6
	r, err := regexp.Compile(`move (?P<count>\d+) from (?P<source>\d+) to (?P<dest>\d+)`)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error compiling regex %v", err))
	}

	//part1(movements, stacks, r)
	part2(movements, stacks, r)

}

func part1(movements []string, stacks []*utils.Stack, r *regexp.Regexp) {
	for _, val := range movements {
		if val == "" {
			continue
		}
		matches := r.FindStringSubmatch(val)
		count, src, dest := matches[1], matches[2], matches[3]
		for i := getInt(count); i > 0; i-- {
			data := stacks[getInt(src)-1].Pop()
			if data == 32 {
				continue
			}
			stacks[getInt(dest)-1].Push(data)
		}
	}

	for _, s := range stacks {
		fmt.Print(string(s.Top()))
	}
	fmt.Println("---------")

}

func part2(movements []string, stacks []*utils.Stack, r *regexp.Regexp) {
	for _, val := range movements {
		if val == "" {
			continue
		}
		matches := r.FindStringSubmatch(val)
		count, src, dest := matches[1], matches[2], matches[3]
		var arr []rune
		for i := getInt(count); i > 0; i-- {
			data := stacks[getInt(src)-1].Pop()
			if data == 32 {
				continue
			}
			arr = append(arr, data)
		}

		for i := len(arr) - 1; i >= 0; i-- {
			stacks[getInt(dest)-1].Push(arr[i])
		}
	}

	for _, s := range stacks {
		fmt.Print(string(s.Top()))
	}
	fmt.Println("---------")
}
