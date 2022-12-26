package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils"
	"math"
	"os"
	"sort"
	"strconv"
)

func getInt(calorie string) int {
	value, err := strconv.Atoi(calorie)
	if err != nil {
		return 0
	}
	return value
}

func findSum(calorie []string) int {
	sum := 0
	for _, value := range calorie {
		sum += getInt(value)
	}
	return sum
}

func findSumInts(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func part1(calories [][]string) int {
	maxCalories := 0
	for _, val := range calories {
		totalCalories := findSum(val)
		maxCalories = int(math.Max(float64(maxCalories), float64(totalCalories)))
	}
	return maxCalories
}

func part2(calories [][]string) int {
	var caloriePerUnit []int
	for _, val := range calories {
		caloriePerUnit = append(caloriePerUnit, findSum(val))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(caloriePerUnit)))
	return findSumInts(caloriePerUnit[:3])

}

func main() {
	f, err := os.Open("./Day1/input.txt")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
		return
	}
	defer utils.CloseFile(f)
	calories := utils.ParseFile(f)
	r1 := part1(calories)
	r2 := part2(calories)

	fmt.Println(r1)
	fmt.Println(r2)
}
