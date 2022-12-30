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
	xdirection := unitVector(tail.x, head.x)
	ydirection := unitVector(tail.y, head.y)
	scale := 1
	if tail.x == head.x {
		tail.y += ydirection * scale
		return
	}
	if tail.y == head.y {
		tail.x += xdirection * scale
		return
	}
	if xdirection == 0 && ydirection == 0 {
		return
	}
	// diagonal > 1
	if xdirection == 0 {
		xdirection = sign(tail.x, head.x)
	}
	if ydirection == 0 {
		ydirection = sign(tail.y, head.y)
	}

	tail.x += xdirection * scale
	tail.y += ydirection * scale

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
