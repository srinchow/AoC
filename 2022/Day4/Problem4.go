package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/file"
	"os"
	"strconv"
	"strings"
)

type rt struct {
	left  int
	right int
}

func newRt(r string) *rt {
	left, right := getRange(r, "-")
	leftInt, err := strconv.Atoi(left)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error converting left value to int %v", left))
		return nil
	}
	rightInt, err := strconv.Atoi(right)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error converting right value to int %v", right))
		return nil
	}
	val := rt{left: leftInt, right: rightInt}
	return &val
}

func getRange(pair, sep string) (string, string) {
	ranges := strings.Split(pair, sep)
	return ranges[0], ranges[1]
}

func (ranges rt) contains(otherRange *rt) bool {
	if otherRange.left >= ranges.left && otherRange.right <= ranges.right {
		return true
	}
	return false
}

func (ranges rt) overlaps(otherRange *rt) bool {
	if otherRange.left >= ranges.left && otherRange.right <= ranges.right {
		return true
	}
	if otherRange.right <= ranges.right && otherRange.right >= ranges.left {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("./2022/Day4/input.txt")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while opening file %v", f))
		return
	}

	defer file.CloseFile(f)
	temp := file.ParseFile(f)
	pairs := temp[0]
	part1(pairs)
	part2(pairs)
}

func part1(pairs []string) {
	cnt := 0
	for _, value := range pairs {
		p1, p2 := getRange(value, ",")
		leftRange := newRt(p1)
		rightRange := newRt(p2)

		if leftRange.contains(rightRange) {
			cnt += 1
		} else if rightRange.contains(leftRange) {
			cnt += 1
		}
	}

	fmt.Println(cnt)
}

func part2(pairs []string) {
	cnt := 0
	for _, value := range pairs {
		p1, p2 := getRange(value, ",")
		leftRange := newRt(p1)
		rightRange := newRt(p2)

		if leftRange.overlaps(rightRange) {
			cnt += 1
		} else if rightRange.overlaps(leftRange) {
			cnt += 1
		}
	}

	fmt.Println(cnt)
}
