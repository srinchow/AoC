package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils"
	"os"
)

func main() {
	file, err := os.Open("./Day3/input.txt")

	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening file %v", err))
	}

	defer utils.CloseFile(file)

	data := utils.ParseFile(file)
	rucksacks := data[0]

	part1(rucksacks)
	part2(rucksacks)

}

func part2(rucksacks []string) {
	score := 0
	for i := 0; i < len(rucksacks); i += 3 {
		score += getScores(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
	}
	fmt.Println(score)
}

func getScores(s string, s2 string, s3 string) int {
	p1 := make(map[rune]int)
	p2 := make(map[rune]int)
	p3 := make(map[rune]int)

	for _, val := range s {
		p1[val]++
	}

	for _, val := range s2 {
		p2[val]++
	}

	for _, val := range s3 {
		p3[val]++
	}
	sum := 0
	for key := range p3 {
		_, ok := p1[key]
		_, ok2 := p2[key]
		if ok && ok2 {
			sum += getScore(byte(key))
		}
	}
	return sum
}

func part1(rucksacks []string) {
	totalPriority := 0
	for _, rucksack := range rucksacks {
		totalPriority += itemScore(rucksack)
	}
	fmt.Println(totalPriority)
}

func itemScore(rucksack string) int { // len of rucksack can only be even
	firstHalfElement := make(map[byte]int)
	secondHalfElement := make(map[byte]int)
	sum := 0
	for i := 0; i < len(rucksack); i++ {
		if i < (len(rucksack) / 2) {
			firstHalfElement[rucksack[i]]++
			continue
		}
		secondHalfElement[rucksack[i]]++
	}

	for key := range secondHalfElement {
		_, ok := firstHalfElement[key]
		if !ok {
			continue
		}
		sum += getScore(key)
	}
	return sum
}

func getScore(u uint8) int {
	if u <= 'z' && u >= 'a' {
		return int(u - 'a' + 1)
	}
	return int(u - 'A' + 27)
}
