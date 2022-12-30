package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/file"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	f, err := os.Open("./2022/Day9/input.txt")

	if err != nil {
		fmt.Println(fmt.Sprintf("Error in opening file %v", err))
	}
	defer file.CloseFile(f)

	movements := file.ParseFile(f)[0]

	head, tail := point{x: 0, y: 0}, point{x: 0, y: 0}

	positions := map[string]bool{"0,0": true}

	for _, val := range movements {
		direction, distance := getVector(val)

		switch direction {
		case "U":
			{
				for i := 1; i <= distance; i++ {
					head.y += 1
					moveTail(&tail, head)
					positions[fmt.Sprintf("%v,%v", tail.x, tail.y)] = true
				}
			}
		case "D":
			{
				for i := 1; i <= distance; i++ {
					head.y -= 1
					moveTail(&tail, head)
					positions[fmt.Sprintf("%v,%v", tail.x, tail.y)] = true
				}
			}
		case "L":
			{
				for i := 1; i <= distance; i++ {
					head.x -= 1
					moveTail(&tail, head)
					positions[fmt.Sprintf("%v,%v", tail.x, tail.y)] = true
				}
			}
		case "R":
			{
				for i := 1; i <= distance; i++ {
					head.x += 1
					moveTail(&tail, head)
					positions[fmt.Sprintf("%v,%v", tail.x, tail.y)] = true
				}

			}
		}

	}

	fmt.Println(len(positions))
}

func moveTail(tail *point, head point) {
	xDirection := unitVector(tail.x, head.x)
	yDirection := unitVector(tail.y, head.y)
	scale := 1
	if tail.x == head.x {
		tail.y += yDirection * scale
		return
	}
	if tail.y == head.y {
		tail.x += xDirection * scale
		return
	}
	if xDirection == 0 && yDirection == 0 {
		return
	}
	// diagonal > 1
	if xDirection == 0 {
		xDirection = sign(tail.x, head.x)
	}
	if yDirection == 0 {
		yDirection = sign(tail.y, head.y)
	}

	tail.x += xDirection * scale
	tail.y += yDirection * scale

}

func sign(src, dest int) int {
	return int(float64(dest-src) / math.Abs(float64(dest-src)))
}

func unitVector(src int, dest int) int {
	if dest == src || math.Abs(float64(dest-src)) == 1 {
		return 0
	}
	if dest > src {
		return +1
	}
	return -1
}

func getVector(val string) (string, int) {
	res := strings.Fields(val)
	return res[0], getInt(res[1])
}

func getInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer")
		return -1
	}
	return res
}
