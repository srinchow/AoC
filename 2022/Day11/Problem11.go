package main

import (
	"fmt"
	"github.com/srinchow/adventOfCode/utils/collection"
	"github.com/srinchow/adventOfCode/utils/file"
	"os"
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

func main() {
	f, err := os.Open("./2022/Day11/input.txt")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}

	defer file.CloseFile(f)

	parsedData := file.ParseFile(f)
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
