package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/file"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("./2022/Day8/input.txt")

	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
	}

	defer file.CloseFile(f)
	parsedData := file.ParseFile(f)[0]
	//part1(parsedData)
	part2(parsedData)
}

func part1(parsedData []string) {
	var rowMajorHeights [][]int
	count := 0
	dedup := map[string]bool{}

	for _, row := range parsedData {
		rowMajorHeights = append(rowMajorHeights, getInts(row))
	}

	colsMajorHeights := transpose(rowMajorHeights)

	for i, row := range rowMajorHeights {
		count += getVisibleTreeCount(row, dedup, i, false)
	}

	for i, row := range colsMajorHeights {
		count += getVisibleTreeCount(row, dedup, i, true)
	}

	fmt.Println(count)

}

func part2(parsedData []string) {
	var rowMajorHeights [][]int
	senicScore := map[string]int64{}

	for _, row := range parsedData {
		rowMajorHeights = append(rowMajorHeights, getInts(row))
	}

	colsMajorHeight := transpose(rowMajorHeights)

	for i, row := range rowMajorHeights {
		getSenicScore(row, senicScore, i, false, false)
		getSenicScore(reverse(row), senicScore, i, false, true)
	}

	for i, row := range colsMajorHeight {
		getSenicScore(row, senicScore, i, true, false)
		getSenicScore(reverse(row), senicScore, i, true, true)
	}

	maxScore := int64(0)

	for _, val := range senicScore {
		if val > maxScore {
			maxScore = val
		}
	}

	fmt.Println(maxScore)
}

func reverse(row []int) []int {
	revRow := make([]int, 0)

	for i := len(row) - 1; i >= 0; i-- {
		revRow = append(revRow, row[i])
	}

	return revRow
}

func transpose(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	cols := make([][]int, m)

	for i := range cols {
		cols[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cols[i][j] = mat[j][i]
		}
	}

	return cols
}

func getVisibleTreeCount(trees []int, dedup map[string]bool, rowIndex int, transpose bool) int {
	tallest := -1
	cnt := 0
	for i := range trees {
		if trees[i] <= tallest {
			continue
		}

		tallest = trees[i]
		currPosition := fmt.Sprintf("%v,%v", rowIndex, i)
		if transpose {
			currPosition = fmt.Sprintf("%v,%v", i, rowIndex)
		}
		if _, ok := dedup[currPosition]; !ok {
			cnt++
			dedup[currPosition] = true
		}

	}

	//reverse Iterate
	tallest = -1
	for i := len(trees) - 1; i >= 0; i-- {
		if trees[i] <= tallest {
			continue
		}
		tallest = trees[i]
		currPosition := fmt.Sprintf("%v,%v", rowIndex, i)
		if transpose {
			currPosition = fmt.Sprintf("%v,%v", i, rowIndex)
		}
		if _, ok := dedup[currPosition]; !ok {
			cnt++
			dedup[currPosition] = true
		}

	}

	return cnt
}

func getSenicScore(trees []int, senicScoreForTree map[string]int64, rowIndex int, transpose bool, reverse bool) {
	if rowIndex == 0 || rowIndex == 98 {
		return
	}

	treePositions := make(map[int][]int)
	for idx, height := range trees {
		treePositions[height] = append(treePositions[height], idx)
	}

	for idx, val := range trees {
		if idx == 0 || idx == len(trees)-1 { //skip perimeter as value will be 0
			continue
		}

		leftClosest := 0
		for i := val; i <= 9; i++ {
			leftClosestI := lowerBound(treePositions[i], idx)
			if leftClosestI < idx && leftClosestI > leftClosest {
				leftClosest = leftClosestI
			}
		}
		position := idx
		if reverse {
			position = len(trees) - 1 - idx
		}
		currPosition := fmt.Sprintf("%v,%v", rowIndex, position)
		if transpose {
			currPosition = fmt.Sprintf("%v,%v", position, rowIndex)
		}

		if _, ok := senicScoreForTree[currPosition]; ok {
			senicScoreForTree[currPosition] *= int64(idx - leftClosest)
		} else {
			senicScoreForTree[currPosition] = int64(idx - leftClosest)
		}
	}
}

// first value less than the given value if no such value exists returns the smallest value in the array
func lowerBound(arr []int, val int) int {
	if arr == nil {
		return val
	}
	left, right := 0, len(arr)-1
	if arr[left] >= val {
		return arr[left]
	}
	if arr[right] < val {
		return arr[right]
	}
	for left < right {
		mid := left + ((right - left) >> 1) // overflow prevention

		if arr[mid] < val && arr[mid+1] >= val {
			return arr[mid]
		}

		if arr[mid] == val {
			if mid == 0 {
				return val
			} else {
				return arr[mid-1]
			}
		}

		if arr[mid] > val {
			right = mid - 1
		}

		if arr[mid] < val {
			left = mid + 1
		}

	}
	return arr[left]
}

func getInts(heights string) []int {
	res := make([]int, 0, 0)
	for _, height := range heights {
		res = append(res, getInt(string(height)))
	}
	return res
}

func getInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to integer")
		return -1
	}
	return res
}
