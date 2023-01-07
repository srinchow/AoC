package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/collection"
	"github.com/srinchow/adventOfCode/utils/file"
	"os"
	"sort"
	"strings"
)

type predicate func(int) bool
type transformer func(int) int

func plus(a int) transformer {
	return func(b int) int {
		return a + b
	}
}

func mult(a int) transformer {
	return func(b int) int {
		return a * b
	}
}

func square() transformer {
	return func(b int) int {
		return b * b
	}
}

type monkey struct {
	startingItems    collection.Queue
	worryOperation   transformer
	test             predicate
	success, failure int
	id               int
	score            int
}

func (m *monkey) hasItem() (int, bool) {
	if m.startingItems.Length() == 0 {
		return -1, false
	}
	return m.startingItems.Pop(), true
}

type modRepresentation map[int]int

func newModRepresentation(val int, mods []int) modRepresentation {
	var m modRepresentation = make(map[int]int, len(mods))
	for _, mod := range mods {
		m[mod] = val % mod
	}
	return m
}

func (m modRepresentation) updateAll(transform transformer) {
	for key, value := range m {
		m[key] = transform(value) % key
	}
}

type monkey2 struct {
	startingItems               []modRepresentation
	test                        int
	transform                   transformer
	success, failure, id, score int
}

func (this *monkey2) nextMonkey(idx int) int {
	val := this.startingItems[idx]
	if (val[this.test])%this.test != 0 {
		return this.failure
	}
	return this.success
}

func main() {
	f, err := os.Open("./2022/Day11/input.txt")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}

	defer file.CloseFile(f)

	//parsedData := file.ParseFile(f)
	//part1(parsedData)

	// manual parsing for part 2
	monkeys := []monkey2{
		{test: 13, transform: mult(3), success: 1, failure: 7, id: 0},
		{test: 2, transform: plus(8), success: 7, failure: 5, id: 1},
		{test: 7, transform: square(), success: 3, failure: 4, id: 2},
		{test: 17, transform: plus(2), success: 4, failure: 6, id: 3},
		{test: 5, transform: plus(3), success: 6, failure: 0, id: 4},
		{test: 11, transform: mult(17), success: 2, failure: 3, id: 5},
		{test: 3, transform: plus(6), success: 1, failure: 0, id: 6},
		{test: 19, transform: plus(1), success: 2, failure: 5, id: 7},
	}
	var tests []int

	startingItems := [][]int{
		{84, 72, 58, 51},
		{88, 58, 58},
		{93, 82, 71, 77, 83, 53, 71, 89},
		{81, 68, 65, 81, 73, 77, 96},
		{75, 80, 50, 73, 88},
		{59, 72, 99, 87, 91, 81},
		{86, 69},
		{91},
	}

	for _, val := range monkeys {
		tests = append(tests, val.test)
	}

	for idx := range monkeys {
		for _, item := range startingItems[idx] {
			monkeys[idx].startingItems = append(monkeys[idx].startingItems, newModRepresentation(item, tests))
		}
	}

	for i := 0; i < 10000; i++ {
		simulateRound2(monkeys)
	}

	var scores []int

	for _, monkey := range monkeys {
		scores = append(scores, monkey.score)
	}
	sort.Ints(scores)
	fmt.Println(scores)
}

func simulateRound2(monkeys []monkey2) {
	for idx := range monkeys {
		monkeys[idx].score += len(monkeys[idx].startingItems)
		transform := monkeys[idx].transform
		next := monkeys[idx].nextMonkey
		for i, item := range monkeys[idx].startingItems {
			item.updateAll(transform)
			nextIndex := next(i)
			monkeys[nextIndex].startingItems = append(monkeys[nextIndex].startingItems, item)
		}
		monkeys[idx].startingItems = nil
	}
}

func part1(parsedData [][]string) {
	var monkeys []*monkey
	for _, data := range parsedData {
		monkeys = append(monkeys, parse(data))
	}

	for i := 0; i < 20; i++ {
		simulateRound(monkeys)
	}

	scores := make([]int, 8)

	for idx := range monkeys {
		scores[idx] = monkeys[idx].score
	}

	fmt.Println(scores)
}

func simulateRound(monkeys []*monkey) {
	for _, m := range monkeys {
		for true {
			item, ok := m.hasItem()
			if !ok {
				break
			}
			res := (m.worryOperation(item)) / 3
			nextMonkey := m.failure
			if m.test(res) {
				nextMonkey = m.success
			}
			for _, m1 := range monkeys {
				if m1.id == nextMonkey {
					m1.startingItems.Push(res)
					break
				}
			}
			m.score++
		}
	}
}

func parse(data []string) *monkey {
	var m monkey
	for _, line := range data {
		line := strings.Split(line, ":")
		info := strings.TrimSpace(line[0])

		if strings.HasPrefix(info, "Monkey") {
			m.id = collection.GetInt(strings.Split(line[0], " ")[1])
		} else if strings.HasPrefix(info, "Starting") {
			itemsStr := strings.Split(strings.TrimSpace(line[1]), ",")
			for _, itemStr := range itemsStr {
				m.startingItems.Push(collection.GetInt(strings.TrimSpace(itemStr)))
			}
		} else if strings.HasPrefix(info, "Operation") {
			op := strings.Split(strings.TrimSpace(line[1]), "=")
			exprTokens := strings.Fields(op[1])
			if t := collection.GetInt(exprTokens[2]); t != -1 {
				switch exprTokens[1] {
				case "+":
					{
						m.worryOperation = plus(t)
					}
				case "-":
					{
						m.worryOperation = plus(-t)
					}
				case "*":
					m.worryOperation = mult(t)

				}
			} else {
				m.worryOperation = square()
			}

		} else if strings.HasPrefix(info, "Test") {
			v := collection.GetInt(strings.Fields(line[1])[2])
			m.test = func(i int) bool {
				return i%v == 0
			}
		} else if strings.HasPrefix(info, "If true") {
			m.success = collection.GetInt(strings.Fields(strings.TrimSpace(line[1]))[3])
		} else if strings.HasPrefix(info, "If false") {
			m.failure = collection.GetInt(strings.Fields(strings.TrimSpace(line[1]))[3])
		}

	}

	return &m
}
