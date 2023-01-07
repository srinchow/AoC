package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/file"
	"os"
)

type point struct {
	x, y     int
	distance int
}

type predicate func(point, point) bool

func (p point) displace(q point) point {
	return point{x: p.x + q.x, y: p.y + q.y, distance: p.distance + 1}
}

func (p point) equal(q point) bool {
	return p.x == q.x && p.y == q.y
}

func (p point) toString() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

type queue []point

func (q *queue) push(p point) {
	*q = append(*q, p)
}

func (q *queue) pop() point {
	if len(*q) == 0 {
		panic("No more elements left in the queue to pop ")
	}
	ele := (*q)[0]
	*q = (*q)[1:]
	return ele
}

func main() {
	f, err := os.Open("./2022/Day12/input.txt")
	defer file.CloseFile(f)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
		return
	}
	parsedData := file.ParseFile(f)[0]
	var mat [][]int
	visited := make(map[string]int)
	var dest point
	var q queue

	startingPoint := 'S'
	endingPoint := 'E'
	for i, row := range parsedData {
		var r []int
		for j, val := range row {
			r = append(r, int(val-'a'))
			if val == startingPoint {
				q.push(point{x: i, y: j})
			}
			if val == endingPoint {
				dest.x = i
				dest.y = j
			}
		}
		mat = append(mat, r)
	}

	displacements := []point{
		{x: -1, y: 0}, // up
		{x: 1, y: 0},  // down
		{x: 0, y: 1},  // right
		{x: 0, y: -1}, // left
	}

	bfs(q, visited, displacements, mat, dest, point.equal)
	/*
		bfs(q, visited, displacements, mat, dest, func(p point, p2 point) bool {

			return mat[p.x][p.y] == 0
		})

	*/

}

func bfs(q queue, visited map[string]int, displacements []point, mat [][]int, dest point, endingCondition predicate) {
	for len(q) > 0 {
		front := q.pop()
		//fmt.Println(front)
		visited[front.toString()] = 1
		for _, d := range displacements {
			nextPoint := front.displace(d)
			if !isValid(nextPoint, front, mat) {
				continue
			}
			if _, ok := visited[nextPoint.toString()]; ok {
				continue
			}
			visited[nextPoint.toString()] = 1
			if endingCondition(nextPoint, dest) {
				fmt.Println(nextPoint.distance)
				return
			}

			q.push(nextPoint)
		}
	}
}

func isValid(nextPoint point, front point, mat [][]int) bool {
	if nextPoint.x < 0 || nextPoint.y < 0 {
		return false
	}

	if nextPoint.x >= len(mat) || nextPoint.y >= len(mat[0]) {
		return false
	}

	// first or last point

	if mat[front.x][front.y] < 0 || mat[nextPoint.x][nextPoint.y] < 0 {
		return true
	}

	if mat[nextPoint.x][nextPoint.y] > mat[front.x][front.y]+1 {
		return false
	}

	return true
}
