package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/file"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p *Point) add(q Point) {
	p.x += q.x
	p.y += q.y
}

func (p *Point) String() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

type rope []Point

func main() {
	f, err := os.Open("./2022/Day9/input.txt")

	if err != nil {
		fmt.Println(fmt.Sprintf("Error in opening file %v", err))
	}
	defer file.CloseFile(f)

	movements := file.ParseFile(f)[0]

	r := make(rope, 0)
	r = append(r, Point{x: 0, y: 0}) // adding head node

	tailLength := 9
	for i := 0; i < tailLength; i++ {
		r = append(r, Point{x: 0, y: 0})
	}

	lastTailPositions := map[string]bool{"0,0": true}

	for _, val := range movements {
		direction, distance := getVector(val)

		switch direction {
		case "U":
			{
				moveTails(Point{0, 1}, distance, lastTailPositions, r)
			}
		case "D":
			{
				moveTails(Point{0, -1}, distance, lastTailPositions, r)
			}
		case "L":
			{
				moveTails(Point{-1, 0}, distance, lastTailPositions, r)
			}
		case "R":
			{
				moveTails(Point{1, 0}, distance, lastTailPositions, r)
			}
		}

	}

	fmt.Println(len(lastTailPositions))
}

func moveTails(displacement Point, scale int, positions map[string]bool, r rope) {
	if len(r) < 2 {
		panic("INCORRECT ROPE LENGTH")
	}
	for i := 0; i < scale; i++ {
		r[0].add(displacement)
		for j := 1; j < len(r); j++ {
			moveTail(&r[j], r[j-1])
		}
		t := fmt.Sprint(&r[len(r)-1])
		positions[t] = true
	}
}

func moveTail(tail *Point, head Point) {
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
